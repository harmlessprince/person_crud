# application repository and binary file name
NAME=person_crud

# application repository path
REPOSITORY=github.com/projects/${NAME}

install:
	go mod download

run-dev:
	echo "Starting Application In Development Mode"
	CompileDaemon -command="./person-crud"

migrate:
	go run ./migrations/migrations.go

migrate_down:
	go run ./migrations/drop_migrations.go

test:
	 gotest ./tests/... -v