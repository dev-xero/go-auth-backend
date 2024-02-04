package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dev-xero/authentication-backend/model"
	"github.com/dev-xero/authentication-backend/util"
	_ "github.com/lib/pq"
)

type PostGreSQL struct {
	database *sql.DB
}

func (repo *PostGreSQL) InsertUser(ctx context.Context, user model.User) error {
	tx, err := repo.database.BeginTx(ctx, nil)
	if err != nil {
		_, err = util.Fail(err, "[FAIL]: could not begin database transaction")
		return err
	}

	// Rollback transaction incase of failure
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var insertQuery = `
		INSERT INTO users (id, username, email, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err = tx.ExecContext(ctx, insertQuery, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("[FAIL]: could not execute insert query")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[FAIL]: could not commit transaction")
	}

	return nil
}
