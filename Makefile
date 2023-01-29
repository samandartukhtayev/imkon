POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_DATABASE=demo_project

-include .env
  
DB_URL=postgresql://postgres:12345@localhost:5432/demo_project?sslmode=disable


print:
	echo "$(DB_URL)"
	
swag:
	swag init -g api/api.go -o api/docs

run:
	go run "./cmd/main.go"


migrate_file:
	migrate create -ext sql -dir migrations/ -seq alter_some_table



migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

.PHONY: run migrateup migratedown