package server

import (
	"betterPPQ/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getProfileHandler(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if user_id == "" || !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	profile, err := s.db.GetProfileByUserID(user_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve profile: " + err.Error()})
		return
	}

	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (s *Server) updateProfileHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}
	var profile_data database.Profile
	if err := c.ShouldBindJSON(&profile_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile data"})
		return
	}

	if err := s.db.UpdateProfile(username, profile_data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}