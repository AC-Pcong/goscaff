package user

import (
	"log/slog"
	"net/http"

	"github.com/AC-Pcong/goscaff/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db     *gorm.DB
	logger *slog.Logger
}

// NewUserHandler creates a new UserHandler.
// It is intended to be used as a wire provider.
func NewUserHandler(db *gorm.DB, logger *slog.Logger) *UserHandler {
	// Auto-migrate the User model
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		logger.Error("Failed to auto-migrate User model", "error", err)
		// Depending on your application's error handling strategy,
		// you might want to panic here or return an error.
		// For now, we'll just log and continue, but this might lead to issues.
	}
	return &UserHandler{db: db, logger: logger}
}

// RegisterRoutes registers user-specific routes to the Gin router group.
func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/:id", h.GetUser)
	rg.POST("/", h.CreateUser)
	rg.PUT("/:id", h.UpdateUser)
	rg.DELETE("/:id", h.DeleteUser)
	rg.GET("/", h.ListUsers)
}

// GetUser handles fetching a single user by ID.
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user model.User
	if result := h.db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser handles creating a new user.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles updating an existing user.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user model.User
	if result := h.db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&model.User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ListUsers handles fetching all users.
func (h *UserHandler) ListUsers(c *gin.Context) {
	var users []model.User
	if result := h.db.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
