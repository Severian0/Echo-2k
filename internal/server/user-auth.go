package server

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
    s := os.Getenv("JWT_SECRET")
    if s == "" {
        log.Fatal("JWT_SECRET must be set")
    }
    jwtSecret = []byte(s)
}


// USER REGISTER AND LOGIN HANDLER
type credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// /api/register
func (s *Server) RegisterUserHandler(c *gin.Context) {
 	var creds credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := s.db.CreateUser(creds.Username, creds.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// /api/login
func (s *Server) loginHandler(c *gin.Context) {
	var creds credentials
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user_id, err := s.db.AuthenticateUser(creds.Username, creds.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if user_id == -1 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    token, err := GenerateToken(creds.Username, strconv.Itoa(user_id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}


// Middleware to check JWT token
type UserClaims struct {
	Username string `json:"username"`
    User_id       string `json:"user_id"`
	jwt.RegisteredClaims
}
func GenerateToken(username string, id string) (string, error) {
	claims := UserClaims{
		Username: username,
        User_id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "bPPQ-API",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
func ParseToken(signedToken string) (*UserClaims, error) {
    token, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, jwt.ErrTokenInvalidClaims
}

func (s *Server) AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var tokenStr string

		auth := c.GetHeader("Authorization")
		parts := strings.Fields(auth)
		if len(parts) == 2 && parts[0] == "Bearer" {
			tokenStr = parts[1]
		}
        if tokenStr == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            return
        }
        claims, err := ParseToken(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": string(err.Error())})
            return
        }
        c.Set("username", claims.Username)
        c.Set("user_id", claims.User_id)
        c.Next()
    }
}

func (s *Server) privHandler(c *gin.Context) {
	username, exists := c.Get("username")
    user_id, exists2 := c.Get("user_id")
    if !exists || !exists2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"username": username.(string), "user_id": user_id.(string)})
}