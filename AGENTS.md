# PG Back Web - AI Coding Agent Instructions

## Architecture Overview

PG Back Web is a Go-based PostgreSQL backup management web application with these core layers:

### Service Layer Pattern

- **Service aggregation**: `internal/service/service.go` creates a single `Service` struct containing all domain services
- **Domain services**: Each service in `internal/service/{domain}/` follows the pattern:
  - `{domain}.go` - service struct and constructor
  - SQL files - queries are embedded alongside Go methods, every Go file that uses SQL has it's own SQL file with the same name.
  - One Go file per major operation (create, update, delete, etc.)

### Database & Code Generation

- **SQLC-based**: Database queries are written in SQL files and auto-generated to Go code
- **Migration-first**: Schema defined in `internal/database/migrations/` using Goose
- **Generated code**: `internal/database/dbgen/` contains SQLC-generated code (never edit directly)
- **Regeneration**: Use `task gen:db` to regenerate database code after SQL changes
- **Prefix queries**: Always prefix queries, for example for the auth service, the queries should be prefixed with `AuthServiceXXXX`, `AuthServiceYYYY`, etc.
- **Don't reuse queries**: Try to avoid reusing queries across services or files, even if they are identical. This keeps the code easier to read and maintain because if a query is changed, it won't affect other services unexpectedly.

### Key Components

- **Integration layer**: `internal/integration/` - abstracts PostgreSQL client and storage operations
- **Echo router**: Web and API routes mounted in `internal/view/router.go`
- **Cron scheduling**: `internal/cron/` wraps gocron for backup scheduling
- **Configuration**: Environment-based config in `internal/config/env.go` with validation

## Development Workflows

### Core Commands (via Taskfile)

Read the `Taskfile.yml` file, but here is a summary of key commands:

```bash
task build        # Build Go binary and frontend assets
task gen:db       # Regenerate SQLC code after SQL changes
task goose -- up  # Run database migrations
task test         # Run all tests
task lint         # Run linters
```

### Frontend Build Process

- **Two-stage build**: TypeScript build script combines app.js + \*.inc.js files
- **Alpine.js integration**: Page-specific JavaScript in `*.inc.js` files within web views
- **TailwindCSS + DaisyUI**: Styling framework with build integration
- **Static embedding**: Frontend assets embedded in Go binary via `embed`

### Database Development

1. Create migration: Add SQL file to `internal/database/migrations/`
2. Write queries: Add `.sql` files to service directories
3. Regenerate: Run `task gen:db` to update Go code
4. Use generated methods in service Go files

## Project-Specific Patterns

### Service Dependencies

Services follow dependency injection with specific patterns:

- **WebhooksService**: Injected into most services for event notifications
- **ExecutionsService**: Central for backup/restore operations
- **Integration layer**: Shared by services needing PostgreSQL/storage operations

### Error Handling & Logging

- **Structured logging**: Use `logger.KV{"key": value}` for context
- **Fatal errors**: Use `logger.FatalError()` for startup failures
- **Request context**: Available via `internal/view/reqctx` for web handlers

### Configuration

- **Required env vars**: `PBW_ENCRYPTION_KEY`, `PBW_POSTGRES_CONN_STRING`
- **Optional env vars**: `PBW_LISTEN_HOST`, `PBW_LISTEN_PORT`, `TZ`
- **Validation**: Environment validation happens in `config/env_validate.go`

### Testing Conventions

- **Table-driven tests**: Use struct slices with `t.Run()` for test cases
- **Helper functions**: Mark test helpers with `t.Helper()`
- **Utility tests**: Focus on `internal/util/` packages with comprehensive coverage

## Critical Integration Points

### Backup Execution Flow

1. **Cron trigger** → `backups.Service` → `executions.Service`
2. **PostgreSQL operations** via `integration.PGClient`
3. **Storage operations** via `integration.StorageClient`
4. **Webhook notifications** via `webhooks.Service`

### Web Request Flow

1. **Router mounting**: Web/API groups with middleware injection
2. **Request context**: Injected via `mids.InjectReqctx` middleware
3. **Service access**: All services available through main `Service` struct

When modifying core components, always regenerate database code and test the full backup execution flow.
