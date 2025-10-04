# Copilot Instructions for Federation: 1999

## Repository Overview

This is a recreation of the Federation space fantasy game as it was in 1999. It's a **mixed-language project** combining Go (primary language), C, and C++ code. The game is a multi-user online text-based space game deployed to Fly.io.

**Project Statistics:**
- Primary languages: Go 1.25, C (gcc-13), C++ (g++-13)
- Project size: ~7 main binaries (Go and C/C++)
- Runtime: Linux (amd64), though non-Linux systems use Docker for building
- Web frontend: Hugo 0.150.1 (extended) with Node.js/npm for dependencies

## Critical Build Information

### Build Order Dependencies

**ALWAYS** follow this sequence when building from clean:

1. **Generate** (required before any build or test):
   ```bash
   make generate
   ```
   This creates:
   - `internal/server/parser/parser.tab.go` (using goyacc)
   - `internal/text/messages.go` (using shell/awk/perl scripts)
   
   **Note**: Without `make generate`, builds and tests will fail with missing file errors.

2. **Build** (platform-specific):
   ```bash
   make build
   ```
   - On Linux: Builds directly using system gcc/g++/go
   - On non-Linux (macOS): Uses Docker with `Dockerfile.build` for cross-compilation
   - Builds C/C++ libraries first (internal/ibgames, internal/fed), then binaries
   - Output: Binaries in `bin/linux-amd64/`
   
   **Build time**: ~6 seconds on Linux for full build

3. **Website** (optional, requires Hugo):
   ```bash
   make website
   ```
   - Requires Hugo 0.150.1 extended version installed
   - First runs `npm ci` in web/, then `hugo build --minify`
   - Hugo is NOT available in standard environments - check CI workflow for installation steps

### Individual Build Targets

**Go binaries** (built with `go build`):
- `fedtpd` - Game engine (main server)
- `httpd` - HTTP server
- `login` - Authentication service
- `perivale-go` - Go version of driver (incomplete)
- `workbench-go` - Go version of planet editor (incomplete)

**C/C++ binaries** (built with Makefiles in cmd/):
- `modemd` - Modem speed limiter (C)
- `perivale` - Player I/O driver (C++)
- `workbench` (workbench-c) - Planet editor/checker (C++)

**C/C++ libraries** (built with Makefiles in internal/):
- `internal/ibgames/lib/` - Common utilities (C)
- `internal/fed/lib/` - Workbench library (C++)

### Testing

```bash
make test
```

- On macOS: Runs with coverage (`-coverprofile=coverage.out`)
- On Linux: Runs without coverage
- **Important**: Tests require `make generate` to have been run first
- **Time**: Tests take 5+ minutes to complete; be patient
- No test failures are expected on a clean build

### Linting

```bash
make lint
```

- Uses `golangci-lint` (must be installed separately)
- Configuration: `.golangci.yaml`
- **Important**: Run `make generate` first, or linting will fail
- Linter config excludes generated files like `internal/text/messages.go`
- Some warnings in C/C++ code are expected (conversion warnings)

### Cleaning

```bash
make clean
```

Removes all generated files, builds, and artifacts. Safe to run anytime.

## CI/CD Pipeline

### GitHub Actions Workflows

**CI Workflow** (`.github/workflows/ci.yml`):
1. Installs Go (from go.mod)
2. Installs Hugo 0.150.1 extended
3. Runs `make generate`
4. Runs golangci-lint-action
5. Runs `make` (builds everything including website)

**CodeQL Workflow** (`.github/workflows/codeql.yml`):
- Analyzes Go, C/C++, JavaScript/TypeScript, and GitHub Actions
- Runs `make` for manual build mode
- Uses security-extended queries

**To replicate CI locally**:
```bash
# Install Hugo 0.150.1 extended if needed
make generate
golangci-lint run
make
```

## Architecture & Layout

### Directory Structure

**Root files:**
- `Makefile` - Main build orchestration
- `go.mod`, `go.sum` - Go dependencies (module: github.com/nosborn/federation-1999)
- `fly.yaml` - Fly.io deployment configuration
- `start` - Legacy startup script
- `Dockerfile.build` - Build environment for non-Linux systems
- `.golangci.yaml` - Linter configuration
- `.editorconfig` - Code style (tabs for Go/Makefile, spaces for JSON/YAML)

**cmd/** - Main binaries:
- Each subdirectory has its own Makefile for C/C++ components
- Go binaries have main.go in the subdirectory
- `cmd/telnetd/` - Documentation only (not built)

**internal/** - Internal packages:
- `internal/server/` - Main game engine code
  - `internal/server/parser/` - Command parser (uses yacc grammar)
  - `internal/server/sol/` - Sol system locations
  - `internal/server/database/` - Player database
- `internal/model/` - Game data structures
- `internal/text/` - Generated message text (from data/messages.txt)
- `internal/fed/` - C++ workbench library
- `internal/ibgames/` - C utility libraries
- `internal/monitoring/` - Prometheus metrics and health checks
- `internal/config/`, `internal/login/`, etc. - Supporting packages

**pkg/** - Public packages:
- `pkg/ibgames/` - IBGames integration (auth, billing, database)

**data/** - Game data:
- `data/messages.txt` - Message templates (processed by tools/gen-text.sh)
- Most data files in .gitignore (runtime data)

**web/** - Hugo website:
- `web/hugo.yaml` - Hugo configuration
- `web/package.json` - NPM dependencies (@xterm/xterm 5.5.0)
- `web/content/` - Content files
- `web/layouts/` - Templates

**scripts/** - Utility scripts:
- `scripts/docker-build.sh` - Docker build for deployment
- `scripts/init-ibgames-db.sh` - Database initialization
- `scripts/publish-to-github.sh` - Public repository sync

**tools/** - Code generation:
- `tools/gen-text.sh`, `tools/gen-text.awk`, `tools/gen-text.pl` - Generate messages.go

**db/migrations/** - Database migrations:
- `db/migrations/001_create_tables.sql` - Initial schema

**deployments/** - Deployment configurations:
- `deployments/docker/` - Docker deployment files
- `deployments/grafana/` - Grafana dashboards

### Key Configuration Files

- `.golangci.yaml` - Enables many linters (exhaustive, gocritic, gosec, etc.), disables errcheck
- `.pre-commit-config.yaml` - Pre-commit hooks (not required for manual changes)
- `.editorconfig` - Tab size 8 for Go/Makefiles, 2-space for JSON/YAML
- `.gitignore` - Excludes bin/, data-*.tar.bz2, coverage files, logs

## Common Pitfalls & Workarounds

### Build Issues

1. **"parser.tab.go: no such file"**: Run `make generate` first
2. **"hugo: not found"**: Hugo is not installed; only needed for `make website`
3. **Tests timing out**: Tests can take 5+ minutes; increase timeout or use `go test ./...` directly
4. **C/C++ compilation warnings**: Expected; many conversion warnings in legacy C++ code
5. **Non-Linux builds failing**: Ensure Docker and buildx are available; Makefile uses Docker automatically

### Go Version

- Go 1.25 is required (specified in go.mod)
- Uses `golang.org/x/tools/cmd/goyacc` as a tool dependency

### Platform-Specific Notes

- **Linux**: Direct builds using system compilers
- **macOS/Windows**: Uses Docker container for cross-compilation to Linux
- Test coverage is only collected on macOS (see Makefile:99-105)

## Development Workflow

### Making Code Changes

1. **Before changing code**: Run `make generate && make build` to establish baseline
2. **After changing Go code**: Run `make build` (or `go build ./cmd/...`)
3. **After changing C/C++ code**: Run `make build` (rebuilds affected components)
4. **After changing parser grammar** (`internal/server/parser/parser.tab.y`): Run `make generate`
5. **After changing messages** (`data/messages.txt`): Run `make generate`
6. **Before committing**: Run `make lint` (if golangci-lint available) and `make test`

### Running the Application

```bash
make dev
```
This builds and runs locally (see `start` script for startup sequence).

### Deployment

```bash
make deploy  # Requires fly CLI
make docker  # Builds Docker image and runs locally
```

## Testing Strategy

- Unit tests in `*_test.go` files alongside source
- Uses `github.com/stretchr/testify` for assertions
- Parser tests in `internal/server/parser/parser_test.go` validate ~500+ commands
- Database tests use `internal/testutil/` helpers
- Integration tests in `pkg/ibgames/billing/billing_integration_test.go`

## Important Notes

- **C/C++ components**: perivale and workbench are from original 1999 codebase, modernized
- **Go rewrites**: Some components have incomplete Go versions (perivale-go, workbench-go)
- **Single-threaded design**: Game engine largely eschews Go concurrency (preserves 1999 behavior)
- **Database**: Uses sqlite3 for IBGames data, custom format for persona/planet data
- **Deployment**: Fly.io with persistent volume mount for game data

## Quick Reference

**Most common commands:**
```bash
make generate      # Generate parser and messages (required first)
make build         # Build all binaries
make test          # Run tests (takes 5+ minutes)
make lint          # Run linter (requires golangci-lint)
make clean         # Remove all generated files and builds
make              # Same as: make build website (needs Hugo)
```

**If something doesn't work:**
1. Check if `make generate` was run
2. Check if you're on Linux (or have Docker for cross-compilation)
3. For website: Check if Hugo 0.150.1 extended is installed
4. For linting: Check if golangci-lint is installed

**Trust these instructions first** - only search for additional information if these instructions are incomplete or found to be incorrect.
