include .env

mock:
	mockery --all --keeptree

migrate:
	migrate -source file://postgres/migrations \
			-database $(DATABASE_URL) up

rollback:
	migrate -source file://postgres/migrations \
			-database $(DATABASE_URL) down

drop:
	migrate -source file://postgres/migrations \
			-database $(DATABASE_URL) drop

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir postgres/migrations $$name

run:
	go run cmd/graphqlserver/*.go

generate:
	go generate ./..

test:
	go test ./... --tags="integration"

test-remove-cache:
	go test ./... --tags="integration" -count=1
