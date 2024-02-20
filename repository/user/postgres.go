package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/dev-xero/authentication-backend/model"
	"github.com/dev-xero/authentication-backend/util"
	_ "github.com/lib/pq"
)

type PostGreSQL struct {
	Database *sql.DB
}

/*
Handles inserting users into the database

Objectives:
  - Create the users table if absent
  - Hash the user password before storing
  - Construct a query to store the user in the database
  - Commit the transaction

Params:
  - ctx:  Method context
  - user: User model to construct the query from

Returns:
  - An error if any stage fails
*/
func (repo *PostGreSQL) InsertUser(ctx context.Context, user model.User) error {
	// Begin a new database transaction
	tx, err := repo.Database.BeginTx(ctx, nil)
	if err != nil {
		_, err = util.Fail(err, "[FAIL]: could not begin database transaction")
		return err
	}

	// Rollback transaction incase of failure (deferred)
	defer func() {
		if rErr := tx.Rollback(); rErr != nil && err == nil {
			err = fmt.Errorf("[FAIL]: rollback failed: %w", rErr)
		}
	}()

	// Create the table if it doesn't yet exist
	if err := repo.createTableIfNonExistent(ctx, tx, "users"); err != nil {
		return err
	}

	// Hash the user password
	user.Password, err = util.GenerateHash(user.Password, util.DefaultHashCost)
	if err != nil {
		return err
	}

	log.Println("Hashed user password:", user.Password)

	// Construct a query to insert the user from the model data
	var insertQuery = `
		INSERT INTO users (id, username, email, password)
		VALUES ($1, $2, $3, $4)
	`

	// Execute the insertion query
	_, err = tx.ExecContext(ctx, insertQuery, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("[FAIL]: could not execute insert query")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[FAIL]: could not commit transaction")
	}

	return nil
}

/*
Checks whether the user with the provided email exists

Objectives:
  - Check if the user exists in the table

Params:
  - ctx:   Method context
  - email: Provided user email

Returns:
  - A boolean indicating whether the user exists
  - An error, in case the query failed
*/
func (repo *PostGreSQL) UserExists(ctx context.Context, email string, username string) (bool, error) {
	// Construct a query to check if a column with the email exists
	var checkUserExistsQuery = `
		SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 OR username = $2)
	`

	// Exists is set to false by default
	var exists = false

	// Check if the user is already stored in the database
	err := repo.Database.QueryRowContext(ctx, checkUserExistsQuery, email, username).Scan(&exists)
	if err != nil {
		log.Println(err)

		// Return false if the table doesn't exist
		if strings.Contains(err.Error(), "does not exist") {
			return false, nil
		}

		return false, fmt.Errorf("[FAIL]: could not check if user already exists")
	}

	return exists, nil
}

/*
Returns a user model with the corresponding id

Objectives:
  - Create the user table if it doesn't already exist
  - Construct a query to return the user details from the id
  - Execute the query and return the user data model

Params:
  - ctx: Method context
  - id:  The user id

Returns:
  - A user data model
  - An error if any
*/
func (repo *PostGreSQL) GetUserByID(ctx context.Context, id string) (model.User, error) {
	// Begin the database transaction
	tx, err := repo.Database.BeginTx(ctx, nil)
	if err != nil {
		_, err = util.Fail(err, "[FAIL]: could not begin database transaction")
		return model.User{}, err
	}

	// Create the users table if it doesn't already exist
	repo.createTableIfNonExistent(ctx, tx, "users")

	// Construct a query to return the user details from the provided id
	var getUserByIDQuery = `
		SELECT id, username, email, password FROM users WHERE id = $1
	`

	// Allocate memory for the user model data
	var user model.User

	// Execute the query, returns the row with the details
	err = tx.QueryRowContext(ctx, getUserByIDQuery, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("[FAIL]: user with ID %s not found", id)
		}
		return model.User{}, fmt.Errorf("[FAIL]: could not execute query: %w", err)
	}

	return user, nil
}

/*
Returns a user with the corresponding email

Objectives:
  - Create the user table if it doesn't already exist
  - Construct a query to return the user details from the email
  - Execute the query and return the user data model

Params:
  - ctx: Method context
  - id:  The user id

Returns:
  - A user data model
  - An error if any
*/
func (repo *PostGreSQL) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	// Begin the transaction
	tx, err := repo.Database.BeginTx(ctx, nil)
	if err != nil {
		_, err = util.Fail(err, "[FAIL]: could not begin database transaction")
		return model.User{}, err
	}

	// Create the users table if it doesn't already exist
	repo.createTableIfNonExistent(ctx, tx, "users")

	// construct a query to return the data model using the email provided
	var getUserByIDQuery = `
		SELECT id, username, email, password FROM users WHERE email = $1
	`

	// Allocate memory for the user data
	var user model.User

	// Execute the query
	err = tx.QueryRowContext(ctx, getUserByIDQuery, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("[FAIL]: user with email %s not found", email)
		}
		return model.User{}, fmt.Errorf("[FAIL]: could not execute query: %w", err)
	}

	return user, nil
}

/*
Creates a table if it doesn't already exist

Objectives:
  - Create the specified table if it doesn't exist

Params:
  - ctx:   Method context
  - tx:    A pointer to the transaction object
  - table: A string indicating the name of the table to create

Returns:
  - An error if any step fails
*/
func (repo *PostGreSQL) createTableIfNonExistent(ctx context.Context, tx *sql.Tx, table string) error {
	// Construct a query to create the table if it doesn't yet exist
	var createTableQuery = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id UUID PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL
		);
	`, table)

	// Execute the query with the context
	_, err := tx.ExecContext(ctx, createTableQuery)
	if err != nil {
		return fmt.Errorf("[FAIL]: could not create users table: %w", err)
	}

	return nil
}
