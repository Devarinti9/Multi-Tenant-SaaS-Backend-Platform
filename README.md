# Multi-Tenant SaaS Backend Platform

Scalable multi-tenant backend for user, organization, and access management with secure APIs, PostgreSQL, Docker, AWS S3, and Cloudflare R2 compatibility.

## Tech Stack
- Go
- Echo Framework
- PostgreSQL
- GORM
- JWT authentication
- Docker / Docker Compose
- AWS S3 / Cloudflare R2 compatible object storage
- Python helper script for demo data seeding

## Features
- User registration and login with JWT authentication
- Organization creation and listing
- Project membership creation for shared tenant environments
- PostgreSQL persistence with GORM auto-migrations
- Optional S3/R2 manifest upload during organization creation
- Dockerized local or Codespaces-ready setup

## API Endpoints
### Public
- `GET /health`
- `GET /`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

### Protected
- `POST /api/v1/organizations`
- `GET /api/v1/organizations`
- `POST /api/v1/memberships`
- `GET /api/v1/memberships`
- `POST /api/v1/users`

## Example Workflow
1. Register a user tied to an organization.
2. Log in to receive a JWT token.
3. Create a tenant organization.
4. Create project memberships for environment-specific access.
5. Optionally push a JSON manifest to S3 or Cloudflare R2.

## Run Locally
```bash
cp .env.example .env

docker compose up --build
```

Or with Go directly:
```bash
go mod download
go run ./cmd/server
```

## Seed Demo Data
```bash
python scripts/seed_demo_data.py
```

## Future Improvements
- role-based authorization rules
- refresh tokens and session management
- tenant-level audit logging
- admin dashboard integration
- CI/CD pipeline and deployment manifests
