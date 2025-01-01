docker:
	docker run --name check -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

sqlc:
	sqlc generate

mgup:
	migrate -database "postgres://root:secret@db:5432/postgres?sslmode=disable" -path ./migrations up
	#migrate -database "postgres://root:secret@localhost:5433/postgres?sslmode=disable" -path ./migrations up

mgdown:
	migrate -database "postgres://root:secret@localhost:5433/postgres?sslmode=disable" -path ./migrations down -all


phony : sqlc, docker