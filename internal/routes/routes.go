package routes

import (
	"multi-tenant-saas-backend-platform/internal/config"
	"multi-tenant-saas-backend-platform/internal/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterPublicRoutes(e *echo.Echo, db *gorm.DB, cfg config.Config) {
	e.POST("/api/v1/auth/register", handlers.Register(db))
	e.POST("/api/v1/auth/login", handlers.Login(db, cfg))
}

func RegisterProtectedRoutes(g *echo.Group, db *gorm.DB, cfg config.Config) {
	g.POST("/organizations", handlers.CreateOrganization(db, cfg))
	g.GET("/organizations", handlers.ListOrganizations(db))
	g.POST("/users", handlers.CreateUser(db))
}
