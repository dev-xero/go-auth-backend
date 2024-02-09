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

func (repo *PostGreSQL) InsertUser(ctx context.Context, user model.User) error {
	tx, err := repo.Database.BeginTx(ctx, nil)
	if err != nil {
		_, err = util.Fail(err, "[FAIL]: could not begin database transaction")
		return err
	}

	// Rollback transaction incase of failure
	defer func() {
		if rErr := tx.Rollback(); rErr != nil && err == nil {
			err = fmt.Errorf("[FAIL]: rollback failed: %w", rErr)
		}
	}()

	// Create the table if it doesn't exist
	if err := repo.createTableIfNonExistent(ctx, tx, "users"); err != nil {
		return err
	}

	// Hash user password
	user.Password, err = util.GenerateHash(user.Password, util.DefaultHashCost)
	if err != nil {
		return err
	}

	log.Println("Hashed user password:", user.Password)

	var insertQuery = `
		INSERT INTO users (id, username, email, password)
		VALUES ($1, $2, $3, $4)
	`
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

func (repo *PostGreSQL) UserExists(ctx context.Context, email string) (bool, error) {
	var checkUserExistsQuery = `
		SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)
	`

	var exists = false

	// Check if the user is already stored in the database
	err := repo.Database.QueryRowContext(ctx, checkUserExistsQuery, email).Scan(&exists)
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

func (repo *PostGreSQL) GetUserByID(ctx context.Context, id string) (model.User, error) {
	tx, err := repo.Database.BeginTx(ctx, nil)
	if err != nil {
		_, err = util.Fail(err, "[FAIL]: could not begin database transaction")
		return model.User{}, err
	}

	repo.createTableIfNonExistent(ctx, tx, "users")

	var getUserByIDQuery = `
		SELECT id, username, email, password FROM users WHERE id = $1
	`
	var user model.User

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

func (repo *PostGreSQL) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	tx, err := repo.Database.BeginTx(ctx, nil)
	if err != nil {
		_, err = util.Fail(err, "[FAIL]: could not begin database transaction")
		return model.User{}, err
	}

	repo.createTableIfNonExistent(ctx, tx, "users")

	var getUserByIDQuery = `
		SELECT id, username, email, password FROM users WHERE email = $1
	`
	var user model.User

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

func (repo *PostGreSQL) createTableIfNonExistent(ctx context.Context, tx *sql.Tx, table string) error {
	var createTableQuery = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id UUID PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`, table)

	_, err := tx.ExecContext(ctx, createTableQuery)
	if err != nil {
		return fmt.Errorf("[FAIL]: could not create users table: %w", err)
	}

	return nil
}
