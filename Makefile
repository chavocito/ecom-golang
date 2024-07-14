run:
	go run cmd/main.go

build:
	go build cmd/main.go

compose:
	docker-compose up -d

migrateup:
	migrate -path db/migration -database "jdbc:mysql://localhost:3306/ecoms?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "jdbc:mysql://localhost:3306/ecoms?sslmode=disable" -verbose down

test:
	go test -v -cover -short ./...

PHONY: run build