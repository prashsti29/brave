# Brave

Local development and secure setup

- Do NOT commit `.env` to version control. This repository already includes `.gitignore` containing `.env`.

Quick setup (local, interactive):

```bash
cp .env.example .env
# Edit .env and fill required values, or run ./setup.sh which will prompt interactively
./setup.sh
```

Secure non-interactive / CI setup

- Do not store secrets in the repository. Use your CI provider's secrets management to inject environment variables.
- Provide a `DATABASE_URL` or `MIGRATE_DATABASE_URL` env var containing the full Postgres URL (e.g., `postgres://user:password@host:port/dbname?sslmode=disable`).

Example CI usage

```bash
# CI should set these as secure variables
export DATABASE_URL="$DATABASE_URL"
# Run migrations without interactive prompts
./setup.sh
```

Notes

- `setup.sh` will refuse to prompt for secrets when running non-interactively; instead, it will exit with instructions to provide secrets via `.env` or env vars.
- `config.ConnectDB()` prefers `DATABASE_URL` if present and falls back to `.env` values.
- Avoid pasting or logging `.env` contents or secrets in public channels.
