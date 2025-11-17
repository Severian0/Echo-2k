package database

import (
	"database/sql"
	"fmt"
)

// each entry is run in order
var migrations = []string{
  `CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     username TEXT UNIQUE NOT NULL,
     hashed_password TEXT NOT NULL,
     created_at TIMESTAMP DEFAULT NOW()
   );`,
   `CREATE TABLE IF NOT EXISTS profiles (
     id SERIAL PRIMARY KEY,
     user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
     display_name TEXT NOT NULL,
     bio TEXT,
     favourite_colour CHAR(7) CHECK (favourite_colour ~ '^#[0-9A-Fa-f]{6}$'),
     image_url TEXT,
     updated_at TIMESTAMP DEFAULT NOW()
    );`,
}

// Migrate applies all pending migrations
func Migrate(db *sql.DB) error {
  for i, stmt := range migrations {
    if _, err := db.Exec(stmt); err != nil {
      return fmt.Errorf("migration[%d] failed: %w", i, err)
    }
  }
  return nil
}