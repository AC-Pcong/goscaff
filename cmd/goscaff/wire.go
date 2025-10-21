//go:build wireinject
// +build wireinject

package main

import (
	"github.com/AC-Pcong/goscaff/internal/handler/user"
	"github.com/AC-Pcong/goscaff/internal/router"
	"github.com/AC-Pcong/goscaff/pkg/config"
	"github.com/AC-Pcong/goscaff/pkg/database"
	"github.com/AC-Pcong/goscaff/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InitializeApp initializes the Gin application with all its dependencies.
func InitializeApp() (*gin.Engine, func(), error) {
	wire.Build(
		config.LoadConfig,
		logger.NewLogger,
		database.NewDB,
		user.NewUserHandler,
		router.NewRouter,
	)
	return nil, nil, nil // This line is just for compilation, wire will replace it.
}
