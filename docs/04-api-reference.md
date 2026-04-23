# API Reference

## `GET /health`
Returns a simple service health response.

## `GET /`
Returns service metadata and a short description.

## `POST /api/v1/auth/register`
### Purpose
Create a user record.

### Required fields
- `email`
- `password`
- `organization_id`

### Optional fields
- `name`
- `role`

## `POST /api/v1/auth/login`
### Purpose
Authenticate a user and receive a JWT token.

### Required fields
- `email`
- `password`

### Response
```json
{
  "token": "<jwt>"
}
```

## `POST /api/v1/organizations`
### Purpose
Create an organization record.

### Headers
`Authorization: Bearer <jwt>`

### Request body
```json
{
  "name": "Acme Corp",
  "plan": "pro"
}
```

### Behavior
- creates the organization row
- derives a slug from the name
- optionally uploads a manifest to object storage when S3 configuration is present

## `GET /api/v1/organizations`
Returns all organizations, preloading their users.

## `POST /api/v1/memberships`
Creates a project membership under an organization.

## `GET /api/v1/memberships`
Returns all membership rows.

## `POST /api/v1/users`
Creates a user via the protected route group using the same registration logic.
