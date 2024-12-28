package services

import (
	"CheckVerifier/db"
	"context"
	"database/sql"
	"log"
)

type Database interface {
	GetLanguage(ctx context.Context, userID int64) (string, error)
}

var queries *db.Queries

func InitDB(dbConn *sql.DB) {
	queries = db.New(dbConn)
}

func SaveUser(user db.User) error {
	params := db.CreateUserParams{
		UserID:   user.UserID,
		Username: user.Username,
	}

	ctx := context.Background()
	err := queries.CreateUser(ctx, params)
	if err != nil {
		log.Printf("Error saving user to database: %v", err)
		return err
	}
	return nil
}
