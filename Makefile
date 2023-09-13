hello:
	echo "Hello World"

run-dev:
	echo "Starting Application In Development Mode"
	CompileDaemon -command="./person-crud"

migrate:
	go run ./migrations/migrations.go

migrate_down:
	go run ./migrations/drop_migrations.go

test:
	 gotest ./tests/... -v