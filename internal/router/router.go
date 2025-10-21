package router

import (
	"log/slog"
	"net/http"

	"github.com/AC-Pcong/goscaff/internal/handler/user"
	"github.com/gin-gonic/gin"
)

// NewRouter creates a new Gin engine and registers all routes.
// It is intended to be used as a wire provider.
func NewRouter(userHandler *user.UserHandler, logger *slog.Logger) *gin.Engine {
	// Set Gin to production mode in production environments
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API V1 group
	v1 := r.Group("/api/v1")
	{
		// User routes
		userGroup := v1.Group("/users")
		userHandler.RegisterRoutes(userGroup)
	}

	logger.Info("All routes registered")

	return r
}
