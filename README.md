# `GO Authentication Backend` 

A web server backend with complete JWT user authentication, written in GO.

## Stack

1. GO: programming language
2. JWT: authentication strategy
3. PostgreSQL: primary database
4. Render: cloud hosting platform

## Application Architecture

<img src="./diagrams/app-architecture.svg" alt="Application Architecture">
   
## Running locally

Before you can run the server locally, you need to create a .env file which stores most of the server's private configurations. An example env file shows all the parameters required, then the server can be run by using the command:

```bash
go run server.go
```

Alternatively, you can batch execute some pre-commands and run the server at once using make.  

Installation (Unix):

```bash
sudo apt update
sudo apt install make
```

Then run the server using:

```bash
make server
```

## Endpoints

1. domain`/`
2. domain`/auth/sign-up`
3. domain`/auth/sign-in`
4. domain`/auth/sign-out`
5. domain`/user/id`

## 1. Sign Up
    
  The URL and port number can be different depending on your configurations
  
  ### Request
  
  ```url
  [POST] http://localhost:3000/auth/sign-up
  ```
  ### Body (JSON)
  
  ```json
  {
    "email":    "string",
    "username": "string",
    "password": "string"
  }
  ```
  Upon successful sign-up, a response like the one below will be sent along with a token stored in the client's cookie store
  
  ### Response

  ```json
  {
      "message": "Successfully inserted user into database",
      "success": true,
      "payload": {
          "id":       "generated-uuid",
          "username": "username",
          "email":    "email"
      }
  }
  ```



<!--- Eraser file: https://app.eraser.io/workspace/OC7cihBr6RUkY14rwfby --->
