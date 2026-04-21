package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"multi-tenant-saas-backend-platform/internal/config"
	"multi-tenant-saas-backend-platform/internal/models"
	"multi-tenant-saas-backend-platform/internal/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type organizationRequest struct {
	Name string `json:"name"`
}

func CreateOrganization(db *gorm.DB, cfg config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req organizationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		org := models.Organization{Name: req.Name}
		if err := db.Create(&org).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		if cfg.S3Bucket != "" {
			if client, err := storage.NewS3Client(cfg); err == nil {
				payload, _ := json.Marshal(map[string]any{
					"id":   org.ID,
					"name": org.Name,
				})
				_, _ = storage.UploadOrganizationManifest(context.Background(), client, cfg.S3Bucket, org.Name, payload)
			}
		}

		return c.JSON(http.StatusCreated, org)
	}
}

func ListOrganizations(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var orgs []models.Organization
		if err := db.Preload("Users").Find(&orgs).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, orgs)
	}
}

func CreateUser(db *gorm.DB) echo.HandlerFunc {
	return Register(db)
}
