DB_URL=postgresql://root:qwerty@localhost:5432/club_hub?sslmode=disable

# run test coverage
test:
	go test -v -cover ./...
# run linter run it to avoid waste actions minuts
lint:
	golangci-lint run
# start the application
start:
	docker-compose up

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

.PHONY: migrateup