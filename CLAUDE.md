# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Architecture

This is a full-stack web application called "jax" with a Go backend API and Vue.js frontend.

### Backend (Go)
- **Entry Point**: `cmd/api/main.go` - HTTP server with graceful shutdown
- **Routing**: Two-tier routing system using `httprouter`
  - Public routes: `/healthcheck`, `/users`, `/tokens`, `/login`, `/logout` 
  - Protected routes: `/v1/*` (requires Bearer token authentication)
- **Data Layer**: `internal/data/` - Database models for Users, Tokens, Stores, Tags, Products, Blacklist
- **Authentication**: Bearer token-based with scope validation, user activation status checking
- **Database**: MySQL with connection pooling, migrations via `golang-migrate`
- **Middleware Stack**: Panic recovery → CORS → Authentication (for protected routes)
- **Utilities**: 
  - `internal/validator/` - Input validation
  - `internal/mailer/` - SMTP email functionality
  - `internal/constant/` - Application constants

### Frontend (Vue.js)
- **Framework**: Vue 3 with TypeScript, Vite build system
- **UI Library**: PrimeVue with custom theming
- **State Management**: Pinia stores in `src/stores/` (account, app, home, store)
- **Routing**: Vue Router with authentication guards via `requireAuth` middleware
- **Layout**: DefaultLayout component wrapping protected routes
- **API Communication**: Axios for HTTP requests to backend

## Development Commands

### Backend (Go)
```bash
# Development
make run/api                    # Run the API server
make audit                      # Tidy deps, format, and vet code
make build/api                  # Build production binaries (local + linux)

# Database
make db/migrations/up           # Apply all migrations
make db/migrations/down         # Rollback most recent migration
make db/migrations/force version=N  # Force migration to specific version
```

### Frontend (Vue.js)
```bash
cd web
npm run dev                     # Start dev server
npm run build                   # Production build
npm run type-check              # TypeScript checking
npm run lint                    # ESLint with auto-fix
npm run format                  # Prettier formatting
```

### Infrastructure
```bash
# Start database
docker compose up

# Environment setup
cp .env.template .env           # Backend environment
cd web && cp .env.template .env # Frontend environment
```

## Key Configuration

### Environment Variables (Backend)
- Database: `MYSQL_*` variables for connection
- Server: `PORT`, `ENVIRONMENT`
- Email: `SMTP_*` variables for mail functionality
- Connection pooling: `DB_MAX_OPEN_CONNS`, `DB_MAX_IDLE_CONNS`, `DB_MAX_IDLE_TIME`

### Environment Variables (Frontend)
- `VITE_PORT` - Development server port
- `VITE_API_ENDPOINT` - Backend API URL

## Database Structure

Core entities managed through `internal/data/models.go`:
- **Users** - Authentication, activation status, user management
- **Tokens** - Authentication tokens with scopes and expiry
- **Stores** - Business/store entities owned by users
- **Tags** - Categorization system
- **Products** - Product catalog linked to stores
- **Blacklist** - Access control

## Authentication Flow

1. User registration creates inactive user + activation token
2. Email activation required before login
3. Login returns Bearer token for API access
4. Protected routes (`/v1/*`) require valid Bearer token
5. User status validation on protected routes

## Code Patterns

### Backend Error Handling
- Custom error responses via `cmd/api/errors.go`
- Standardized JSON error format
- Panic recovery middleware

### Frontend State Management
- Pinia stores for different domains (account, app, home, store)
- Authentication state managed in account store
- Loading states via app store

### Database Access
- Repository pattern via model structs with embedded `*sql.DB`
- Prepared statements for queries
- Context-aware database operations