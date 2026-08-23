run:
	go run cmd/main.go

migrate-up:
	migrate -path db/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DATABASE_URL)" down

migrate-fix:
	migrate -path db/migrations -database "$(DATABASE_URL)" force $(version)

seed:
	go run cmd/seed/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

setup:
	cp .env.example .env
	go mod download
	make migrate-up