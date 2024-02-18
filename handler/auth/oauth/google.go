package handler

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dev-xero/authentication-backend/service"
	"github.com/dev-xero/authentication-backend/util"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  "",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

const oauthGoogleUserURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func initializeGoogleConfig() error {
	// Load environment variables from .env file in development
	if env := os.Getenv("ENVIRONMENT"); env != "production" {
		err := godotenv.Load()
		if err != nil {
			return fmt.Errorf("[FAIL]: could not load environment variables: %w", err)
		}
	}

	// Configure sensitive information
	googleOauthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	googleOauthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	googleOauthConfig.RedirectURL = os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")

	return nil
}

/*
Handles Google account sign-in with OAuth 2.0

Objectives:
  - Setup auth config struct
  - Request authentication from Google
  - Handle auth callback

Params:
  - auth: The auth repo service
  - w:    A http response writer
  - r:    A pointer to a http request object

Returns:
  - No return value
*/
func GoogleSignIn(auth *service.AuthService, w http.ResponseWriter, r *http.Request) {
	// Initialize the config
	err := initializeGoogleConfig()
	if err != nil {
		log.Println("[FAIL]: Could not configure google oauth")
		util.JsonResponse(w, "Failed to configure Google OAuth", http.StatusInternalServerError, nil)
	}

	// Generate the auth state and redirect
	oauthState := util.GenerateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

/*
Handles callbacks to Google sign-in

Objectives:
  - Handle auth callback
  - Save user to the database

Params:
  - auth: The auth repo service
  - w:    A http response writer
  - r:    A pointer to a http request object

Returns:
  - No return value
*/
func GoogleSignInCallback(auth *service.AuthService, w http.ResponseWriter, r *http.Request) {
	// Read state from cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("[AUTH]: Oauth states do not match")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Get user info from Google
	data, err := getGoogleUserData(r.FormValue("code"))
	if err != nil {
		log.Println("[FAIL]:", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Respond with the data
	fmt.Fprintf(w, "%s\n", data)
}

/*
Handles getting user data from Google provided an auth code

Params:
  - code: The auth exchange code

Returns:
  - The user data which is a byte slice
  - An error if any step fails
*/
func getGoogleUserData(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code")
	}

	url := fmt.Sprintf("%s%s", oauthGoogleUserURL, token.AccessToken)
	log.Println("[LOG]:", url)

	// Make a response using the token
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info")
	}

	// Read user data
	defer res.Body.Close()
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response")
	}

	return contents, nil
}
