docker:
	docker run --name check -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

sqlc:
	sqlc generate

mgup:
	migrate -database "postgres://root:4lnv041DhUf1BYeFmnX09lRPEVN2SaFV@dpg-ctqk213tq21c73a2te1g-a.oregon-postgres.render.com:5432/checkdatabase?sslmode=require" -path ./migrations up

mgdown:
	migrate -database "postgres://root:4lnv041DhUf1BYeFmnX09lRPEVN2SaFV@dpg-ctqk213tq21c73a2te1g-a.oregon-postgres.render.com:5432/checkdatabase?sslmode=require" -path ./migrations down -all



phony : sqlc, docker