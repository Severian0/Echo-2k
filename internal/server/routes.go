package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080","http://localhost:5173"}, // Add your frontend URL
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
    r.Use(func(c *gin.Context) {
        // only intercept GET/HEAD
        // if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
        //     c.Next()
        //     return
        // }
        // let anything under /api through
        if strings.HasPrefix(c.Request.URL.Path, "/api/") {
            c.Next()
            return
        }
        // try to serve the file if it exists
        f := filepath.Join(staticDir, c.Request.URL.Path)
        if info, err := os.Stat(f); err == nil && !info.IsDir() {
            c.File(f)
            c.Abort()
            return
        }
        // otherwise SPA fallback
        c.File(filepath.Join(staticDir, "index.html"))
        c.Abort()
    })
 
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
