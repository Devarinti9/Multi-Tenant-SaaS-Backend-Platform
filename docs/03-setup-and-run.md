# Setup and Run Guide

## Prerequisites
- Go
- PostgreSQL
- Optional: Docker / Docker Compose

## Environment variables
Create `.env` from `.env.example`.

### Important variables
- `APP_PORT`
- `DATABASE_URL`
- `JWT_SECRET`
- `AWS_REGION`
- `S3_BUCKET`
- `S3_ENDPOINT`
- `S3_ACCESS_KEY`
- `S3_SECRET_KEY`
- `S3_USE_PATH_STYLE`

## Docker start
```bash
docker compose up --build
```

## Direct Go run
```bash
go mod download
go run ./cmd/server
```

## Startup behavior
On boot the service:
1. loads configuration
2. connects to PostgreSQL
3. auto-migrates the three core tables
4. starts the Echo HTTP server

## Recommended validation sequence
1. `GET /health`
2. register a user
3. log in and capture the JWT
4. call protected organization endpoint with `Authorization: Bearer <token>`

## Seeding demo data
```bash
python scripts/seed_demo_data.py
```
