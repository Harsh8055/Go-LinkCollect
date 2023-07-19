postgres: 
	docker run --name postgresDB -p 5433:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=root -d postgres:16beta1-alpine

createdb:
	docker exec -it postgresDB createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresDB dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable" -verbose down
	
sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: createdb migrateup migratedown dropdb postgres sqlc
