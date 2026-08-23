# Brave - Village Builder Strategy Game

A multiplayer strategy game inspired by Clash of Clans, built with Go and PostgreSQL. Players build villages, construct defenses, train troops, and battle other players.

## Quick Start

```bash
# Setup
cp .env.example .env
go mod download
make migrate-up

# Run server
make run
```

## Architecture

- **Framework**: Gorilla Mux (HTTP routing)
- **Database**: PostgreSQL with golang-migrate
- **Structure**: MVC pattern
  - `controllers/` - HTTP handlers
  - `service/` - Business logic
  - `repository/` - Data access layer
  - `models/` - Entity definitions
  - `db/migrations/` - Schema migrations

## Features

- Player management (registration, profiles)
- Village layout customization
- Building construction and upgrades
- Battle system