package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"multi-tenant-saas-backend-platform/internal/config"
	"multi-tenant-saas-backend-platform/internal/models"
	"multi-tenant-saas-backend-platform/internal/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrganizationRequest struct {
	Name string `json:"name"`
	Plan string `json:"plan"`
}

type MembershipRequest struct {
	OrganizationID uint   `json:"organization_id"`
	ProjectName    string `json:"project_name"`
	Environment    string `json:"environment"`
}

func slugify(name string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(name)), " ", "-")
}

func CreateOrganization(db *gorm.DB, cfg config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req OrganizationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}
		if req.Name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "organization name is required"})
		}

		org := models.Organization{Name: req.Name, Slug: slugify(req.Name), Plan: req.Plan}
		if err := db.Create(&org).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		manifest := map[string]any{"id": org.ID, "name": org.Name, "slug": org.Slug, "plan": org.Plan}
		storagePath := ""
		if cfg.S3Bucket != "" {
			if client, err := storage.NewS3Client(cfg); err == nil {
				payload, _ := json.Marshal(manifest)
				storagePath, _ = storage.UploadOrganizationManifest(context.Background(), client, cfg.S3Bucket, org.Slug, payload)
			}
		}

		return c.JSON(http.StatusCreated, map[string]any{"organization": org, "manifest_path": storagePath})
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

func CreateMembership(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req MembershipRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}
		if req.OrganizationID == 0 || req.ProjectName == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "organization_id and project_name are required"})
		}

		membership := models.ProjectMembership{
			OrganizationID: req.OrganizationID,
			ProjectName:    req.ProjectName,
			Environment:    req.Environment,
			StoragePath:    "/organizations/demo/manifest.json",
		}
		if err := db.Create(&membership).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, membership)
	}
}

func ListMemberships(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var memberships []models.ProjectMembership
		if err := db.Find(&memberships).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, memberships)
	}
}
