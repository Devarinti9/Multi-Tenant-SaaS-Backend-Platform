package routes

import (
	"net/http"

	"multi-tenant-saas-backend-platform/internal/config"
	"multi-tenant-saas-backend-platform/internal/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"service": "multi-tenant-saas-backend-platform",
		"message": "Scalable multi-tenant backend for organizations, users, and project memberships",
	})
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok", "service": "multi-tenant-saas-backend-platform"})
}

func RegisterPublicRoutes(e *echo.Echo, db *gorm.DB, cfg config.Config) {
	e.POST("/api/v1/auth/register", handlers.Register(db))
	e.POST("/api/v1/auth/login", handlers.Login(db, cfg))
}

func RegisterProtectedRoutes(g *echo.Group, db *gorm.DB, cfg config.Config) {
	g.POST("/organizations", handlers.CreateOrganization(db, cfg))
	g.GET("/organizations", handlers.ListOrganizations(db))
	g.POST("/memberships", handlers.CreateMembership(db))
	g.GET("/memberships", handlers.ListMemberships(db))
	g.POST("/users", handlers.Register(db))
}
