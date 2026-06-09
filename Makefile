run:
	go run cmd/main.go

migrate-up:
	migrate -path db/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DATABASE_URL)" down

migrate-fix:
	migrate -path db/migrations -database "$(DATABASE_URL)" force $(version)

setup:
	cp .env.example .env
	go mod download
	make migrate-up