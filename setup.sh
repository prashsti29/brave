#!/bin/bash
echo "Setting up Brave..."
if [ ! -f .env ]; then
    cp .env.example .env
    echo "Fill in your .env file then press enter to continue..."
    read
fi
# If DB_PASSWORD is empty in .env, prompt the user and save it (only in interactive mode)
DB_PASSWORD_VALUE=$(grep -m1 '^DB_PASSWORD=' .env | cut -d '=' -f2-)
if [ -z "$DB_PASSWORD_VALUE" ]; then
    # If not running interactively (CI or external automation), refuse to prompt and instruct the user
    if [ ! -t 0 ]; then
        echo "ERROR: DB_PASSWORD is not set in .env and script is running non-interactively."
        echo "Set DB_PASSWORD in .env or export DB_PASSWORD in the environment and re-run the script."
        echo "For CI, provide DB credentials via secure CI secrets, not by committing .env."
        exit 1
    fi

    DB_USER_VALUE=$(grep -m1 '^DB_USER=' .env | cut -d '=' -f2-)
    echo "Warning: do not share your .env file or paste secrets publicly."
    echo -n "Enter DB password for user ${DB_USER_VALUE}: "
    read -s INPUT_DB_PASSWORD
    echo
    # Replace or add DB_PASSWORD in .env
    if grep -q '^DB_PASSWORD=' .env; then
        sed -i "s@^DB_PASSWORD=.*@DB_PASSWORD=${INPUT_DB_PASSWORD}@" .env
    else
        echo "DB_PASSWORD=${INPUT_DB_PASSWORD}" >> .env
    fi
fi
go mod download

# Prefer an explicit migrate URL provided via env (useful for CI):
# MIGRATE_DATABASE_URL or DATABASE_URL take precedence over composing from .env
MIGRATE_DB_URL="${MIGRATE_DATABASE_URL:-${DATABASE_URL:-}}"
if [ -z "$MIGRATE_DB_URL" ]; then
    # load values from .env (fall back to env vars if set)
    DB_USER_VALUE="${DB_USER:-$(grep -m1 '^DB_USER=' .env | cut -d '=' -f2-)}"
    DB_PASSWORD_VALUE="${DB_PASSWORD:-$(grep -m1 '^DB_PASSWORD=' .env | cut -d '=' -f2-)}"
    DB_HOST_VALUE="${DB_HOST:-$(grep -m1 '^DB_HOST=' .env | cut -d '=' -f2-)}"
    DB_PORT_VALUE="${DB_PORT:-$(grep -m1 '^DB_PORT=' .env | cut -d '=' -f2-)}"
    DB_NAME_VALUE="${DB_NAME:-$(grep -m1 '^DB_NAME=' .env | cut -d '=' -f2-)}"

    if [ -z "$DB_USER_VALUE" ] || [ -z "$DB_PASSWORD_VALUE" ] || [ -z "$DB_HOST_VALUE" ] || [ -z "$DB_PORT_VALUE" ] || [ -z "$DB_NAME_VALUE" ]; then
        echo "ERROR: Missing DB connection info. Set DATABASE_URL or fill .env with DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME."
        exit 1
    fi

    MIGRATE_DB_URL="postgres://${DB_USER_VALUE}:${DB_PASSWORD_VALUE}@${DB_HOST_VALUE}:${DB_PORT_VALUE}/${DB_NAME_VALUE}?sslmode=disable"
fi

# Run migrations (do not print the URL as it may contain secrets)
migrate -path db/migrations -database "$MIGRATE_DB_URL" up

# Start the server
go run cmd/main.go