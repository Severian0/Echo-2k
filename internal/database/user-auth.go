package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user in the database.
func (s *service) CreateUser(username string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	_, err = s.db.Exec("INSERT INTO USERS (username, hashed_password) VALUES ($1,$2)", username, string(hash))
	if err != nil {
		return fmt.Errorf("failed to perform user entry: %w", err)
	}
	_, err = s.db.Exec("INSERT INTO PROFILES (user_id, display_name) VALUES (currval('USERS_id_seq'), $1)", username)
	if err != nil {
		return fmt.Errorf("failed to perform profile entry: %w", err)
	}


	return nil
}

func (s *service) AuthenticateUser(username string, password string) (int, error) {
	var storedHash string
	var id int
	err := s.db.QueryRow("SELECT hashed_password, id FROM USERS WHERE username = $1", username).Scan(&storedHash, &id)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, nil // User not found
		}
		return -1, fmt.Errorf("failed to query user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return -1, nil // Password does not match
		}
		return -1, fmt.Errorf("failed to compare password: %w", err)
	}

	return id, nil // Authentication successful
}