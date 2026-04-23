# Architecture

## High-level flow
Client -> Echo API -> Handler -> GORM -> PostgreSQL
                               -> Optional S3/R2 manifest upload

## Application structure
- `cmd/server/main.go` — application entry point and HTTP server bootstrap
- `internal/config/config.go` — configuration loading from environment variables
- `internal/db/db.go` — PostgreSQL connection and migrations
- `internal/models/models.go` — GORM entity definitions
- `internal/handlers/auth.go` — registration and login handlers
- `internal/handlers/organization.go` — organization and membership handlers
- `internal/middleware/jwt.go` — JWT middleware for protected routes
- `internal/routes/routes.go` — route registration
- `internal/storage/s3.go` — S3 / R2-compatible upload helper
- `scripts/seed_demo_data.py` — helper script for sample data seeding

## Public routes
- `GET /health`
- `GET /`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

## Protected routes
Available under `/api/v1` with JWT middleware enabled:
- `POST /organizations`
- `GET /organizations`
- `POST /memberships`
- `GET /memberships`
- `POST /users`

## Current tenant model
### Organization
Represents a tenant boundary.

### User
Belongs to one organization and carries a role.

### ProjectMembership
Represents project access context inside an organization.

## Storage behavior
When S3 settings are provided, organization creation attempts to upload a small JSON manifest to object storage.
