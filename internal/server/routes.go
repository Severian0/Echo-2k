package server

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:5173", "https://echo-2k.hydroxide.systems", "https://echo360.org.uk"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	// ========== 1) API group ==========
	api := r.Group("/api")
	{
		api.GET("/hello", s.HelloWorldHandler)
		api.GET("/DBhealth", s.healthHandler)
		api.POST("/register", s.RegisterUserHandler)
		api.POST("/login", s.loginHandler)
		api.GET("/private", s.AuthMiddleware(), s.privHandler)
		api.GET("/profile/:username", s.AuthMiddleware(), s.getProfileHandler)
		api.PUT("/profile/:username", s.AuthMiddleware(), s.updateProfileHandler)
		// mount other api handlers here…
	}

	staticDir := filepath.Join("/", "dist")
	redirectToApp := func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	r.GET("/", redirectToApp)
	r.HEAD("/", redirectToApp)

	spaHandler := func(c *gin.Context) {
		reqPath := c.Param("path")
		cleaned := path.Clean("/" + reqPath)
		relPath := strings.TrimPrefix(cleaned, "/")

		filePath := staticDir
		if relPath != "" && relPath != "." {
			filePath = filepath.Join(staticDir, relPath)
		}

		if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
			c.File(filePath)
			return
		}

		c.File(filepath.Join(staticDir, "index.html"))
	}

	app := r.Group("/app")
	app.GET("/*path", spaHandler)
	app.HEAD("/*path", spaHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
