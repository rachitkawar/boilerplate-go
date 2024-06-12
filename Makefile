
migrations-up:
	go run src/cmd/scripts/migration-up/main.go

migrations-down:
	go run src/cmd/scripts/migration-down/main.go

# make migrations-create table_name=profile
migrations-create:
	migrate create -ext sql -dir src/internal/database/migrations ${table_name}

build:
	go build -o ./bin/.main ./src/cmd/app/main.go

run:
	./bin/.main