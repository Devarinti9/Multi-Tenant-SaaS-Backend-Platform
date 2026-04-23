# Data Model and Security

## Data model
### Organization
- `id`
- `name`
- `slug`
- `plan`
- `created_at`

### User
- `id`
- `organization_id`
- `name`
- `email`
- `password_hash`
- `role`
- `created_at`

### ProjectMembership
- `id`
- `organization_id`
- `project_name`
- `environment`
- `storage_path`
- `created_at`

## Security behavior
### Password handling
Passwords are hashed with bcrypt before storage.

### Authentication
Login returns a signed JWT containing:
- `user_id`
- `email`
- `org_id`
- `exp`

### Authorization scope
Protected routes require a valid JWT. The current starter implementation does not yet enforce fine-grained role-based access control.

## Storage integration notes
The repository supports optional object upload to an S3-compatible endpoint. This is a convenience feature, not a required path for the API to function.

## Important limitation
This starter implementation is suitable for demos and learning, but it still needs:
- stronger authorization rules
- refresh token/session handling
- audit logging
- production-grade secret management
