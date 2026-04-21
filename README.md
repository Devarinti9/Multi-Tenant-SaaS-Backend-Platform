# Multi-Tenant SaaS Backend Platform

A Go-based backend platform for user, organization, and access management across shared application environments.

## Tech Stack
- Golang
- Echo Framework
- PostgreSQL
- Docker
- AWS S3 / Cloudflare R2-compatible object storage
- JWT authentication

## Features
- User registration and login
- JWT-protected APIs
- Multi-tenant organization management
- PostgreSQL persistence
- Dockerized local setup
- S3/R2-compatible storage integration scaffold

## API Endpoints
### Public
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

### Protected
- `POST /api/v1/organizations`
- `GET /api/v1/organizations`
- `POST /api/v1/users`

## Run Locally
```bash
docker compose up --build
```

API will be available at `http://localhost:8082`

## Resume Description
Built a scalable multi-tenant backend platform to support user, organization, and access management across shared application environments, enabling secure and structured handling of multi-user workflows.
