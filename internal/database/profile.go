package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Profile struct {
	User_id          int    `json:"id"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	FavouriteColour string `json:"favourite_colour"`
	ImageURL    string `json:"image_url"`
	
}

func (s *service) GetProfileByUserID(user_id string) (*Profile, error) {
	var profile Profile
	err := s.db.QueryRow("SELECT display_name, bio, favourite_colour, image_url FROM PROFILES WHERE user_id = $1", user_id).Scan(
		&profile.DisplayName,
		&profile.Bio,
		&profile.FavouriteColour,
		&profile.ImageURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Profile not found
		}
		return nil, fmt.Errorf("failed to query profile: %w", err)
	}
	return &profile, nil
}

func (s *service) UpdateProfile(user_id string, profile_data Profile) error {
	_, err := s.db.Exec(`
	UPDATE PROFILES
	SET display_name = $1,
		bio = $2,
		favourite_colour = $3,
		image_url = $4
	WHERE user_id = $5
	`, profile_data.DisplayName, profile_data.Bio, profile_data.FavouriteColour, profile_data.ImageURL, user_id)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}
	return nil
}