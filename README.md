<p><a target="_blank" href="https://app.eraser.io/workspace/OC7cihBr6RUkY14rwfby" id="edit-in-eraser-github-link"><img alt="Edit in Eraser" src="https://firebasestorage.googleapis.com/v0/b/second-petal-295822.appspot.com/o/images%2Fgithub%2FOpen%20in%20Eraser.svg?alt=media&amp;token=968381c8-a7e7-472a-8ed6-4a6626da5501"></a></p>

# `GO Authentication Backend` 
A web server backend with complete JWT user authentication, written in GO.

## Stack
1. GO: programming language
2. JWT: authentication strategy
3. PostgreSQL: primary database
4. Render: cloud hosting platform
## Running locally
Before you can run the server locally, you need to create a .env file which stores most of the server's private configurations. An example env file shows all the parameters required, then the server can be run by using the command:

```bash
go run server.go
```
## Endpoints
Available endpoints for testing:

### Authentication
1. Sign-up
The URL and port number can be different depending on your configurations

**Request**

```
[POST] http://localhost:3000/sign-up
```
**Body** (JSON)

```
{
  "email":    "string",
  "username": "string",
  "password": "string"
}
```
Upon successful sign-up, a response like the one below will be sent along with a token stored in the client's cookie store

**Response**

```
{
    "message": "Successfully inserted user into database",
    "success": true,
    "payload": {
        "ID": "some-uuid",
        "Username": "username",
        "Email": "email",
        "Password": "password"
    }
}
```



<!--- Eraser file: https://app.eraser.io/workspace/OC7cihBr6RUkY14rwfby --->