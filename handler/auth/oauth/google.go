package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dev-xero/authentication-backend/authentication"
	"github.com/dev-xero/authentication-backend/model"
	"github.com/dev-xero/authentication-backend/service"
	"github.com/dev-xero/authentication-backend/util"
	"github.com/google/uuid"
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
  - dbService: The database service provider
  - w:         A http response writer
  - r:         A pointer to a http request object

Returns:
  - No return value
*/
func GoogleSignIn(w http.ResponseWriter, r *http.Request) {
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
  - dbService: The database service provider
  - w:         A http response writer
  - r:         A pointer to a http request object

Returns:
  - No return value
*/
func GoogleSignInCallback(dbService *service.DatabaseProvider, w http.ResponseWriter, r *http.Request) {
	// Read state from cookie
	oauthState, _ := r.Cookie("oauthstate")
	failureRedirectURL := "failure"

	if r.FormValue("state") != oauthState.Value {
		log.Println("[AUTH]: Oauth states do not match")
		http.Redirect(w, r, failureRedirectURL, http.StatusTemporaryRedirect)
		return
	}

	// Get user info from Google
	userData, err := getGoogleUserData(r.FormValue("code"))
	if err != nil {
		log.Println("[FAIL]:", err.Error())
		http.Redirect(w, r, failureRedirectURL, http.StatusTemporaryRedirect)
		return
	}

	// Make sure the user doesn't already exist
	userExists, _ := dbService.Repo.UserExists(r.Context(), userData.Email, userData.Username)
	if userExists {
		util.JsonResponse(w, "User already exists", http.StatusConflict, nil)
		return
	}

	// Save user to database
	err = dbService.Repo.InsertUser(r.Context(), *userData)
	if err != nil {
		log.Println("[FAIL]: could not insert user")
		util.JsonResponse(w, "Failed to create new user", http.StatusInternalServerError, nil)
		return
	}

	// Generate a new token and send response
	token, err := authentication.CreateJWToken(userData.ID)
	if err != nil {
		log.Println(err)
		msg := "Failed to create token"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// Create the user payload
	var userPayload = util.UserPayload{
		ID:       userData.ID,
		Username: userData.Username,
		Email:    userData.Email,
	}

	// Respond with the payload
	cookie := util.CreateTokenCookie(token)
	http.SetCookie(w, &cookie)
	util.JsonResponse(w, "Successfully signed-in with Google", http.StatusOK, userPayload)
}

/*
Handles getting user data from Google provided an auth code

Params:
  - code: The auth exchange code

Returns:
  - The user data which is a byte slice
  - An error if any step fails
*/
func getGoogleUserData(code string) (*model.User, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return &model.User{}, fmt.Errorf("failed to exchange code")
	}

	url := fmt.Sprintf("%s%s", oauthGoogleUserURL, token.AccessToken)
	log.Println("[LOG]:", url)

	// Make a response using the token
	res, err := http.Get(url)
	if err != nil {
		return &model.User{}, fmt.Errorf("failed to get user info")
	}

	// Read user data
	defer res.Body.Close()

	var responseData struct {
		Username string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"id"`
	}

	// Read response into struct
	if err := json.NewDecoder(res.Body).Decode(&responseData); err != nil {
		return &model.User{}, fmt.Errorf("failed to decode response: %w", err)
	}

	// Create user data model
	var userData = &model.User{
		ID:       uuid.New(),
		Username: responseData.Username,
		Email:    responseData.Email,
		Password: responseData.Password,
	}

	return userData, nil
}
