# Troubleshooting

## Database connection failure
### Cause
`DATABASE_URL` is incorrect or PostgreSQL is not reachable.

### Fix
Verify the database URL and confirm the database container or service is running.

## JWT login works but protected route fails
### Cause
Missing or malformed `Authorization` header.

### Fix
Use:
```http
Authorization: Bearer <token>
```

## S3 upload path is empty
### Clarification
The API still works even if manifest upload is skipped. Upload only occurs when S3 settings are provided and the client is created successfully.

## Go dependency issue
### Cause
Dependencies may not be fully resolved on a fresh clone.

### Fix
Run:
```bash
go mod tidy
go mod download
```

## Register endpoint returns validation error
### Cause
Required fields are missing.

### Fix
Provide:
- `email`
- `password`
- `organization_id`
