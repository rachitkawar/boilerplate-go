
migrations-up:
	go run src/cmd/scripts/migration-up/migration-up.go

migrations-down:
	go run src/cmd/scripts/migration-down/migration-down.go

# make migrations-create table_name=profile
migrations-create:
	migrate create -ext sql -dir src/internal/database/migrations ${table_name}

build:
	go build -o ./bin/.main ./src/cmd/app/main.go

run:
	./bin/.main