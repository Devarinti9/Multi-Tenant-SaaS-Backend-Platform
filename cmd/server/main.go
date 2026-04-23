package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"multi-tenant-saas-backend-platform/internal/config"
	"multi-tenant-saas-backend-platform/internal/db"
	appmiddleware "multi-tenant-saas-backend-platform/internal/middleware"
	"multi-tenant-saas-backend-platform/internal/routes"
)

func main() {
	cfg := config.Load()

	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", routes.Health)
	e.GET("/", routes.Index)

	routes.RegisterPublicRoutes(e, database, cfg)
	protected := e.Group("/api/v1")
	protected.Use(appmiddleware.JWTMiddleware(cfg.JWTSecret))
	routes.RegisterProtectedRoutes(protected, database, cfg)

	log.Printf("server running on %s", cfg.AppPort)
	e.Logger.Fatal(e.Start(cfg.AppPort))
}
