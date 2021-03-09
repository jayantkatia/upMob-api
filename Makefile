postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine 
createdb:
	docker exec -it postgres13 createdb --username=root --owner=root upcoming_mobiles
dropdb:
	docker exec -it postgres13 dropdb upcoming_mobiles
migrateup:
	~/Downloads/migrate.linux-amd64 -path db/migration -database "postgresql://root:secret@localhost:5432/upcoming_mobiles?sslmode=disable" -verbose up
migratedown:
	~/Downloads/migrate.linux-amd64 -path db/migration -database "postgresql://root:secret@localhost:5432/upcoming_mobiles?sslmode=disable" -verbose down
dockersql:
	docker exec -it postgres13 psql -U root -d upcoming_mobiles
sqlcgenerate:
	~/Downloads/sqlc-v1.7.0-linux-amd64/sqlc generate
sqlccompile:
	~/Downloads/sqlc-v1.7.0-linux-amd64/sqlc compile
sqlcinit:
	~/Downloads/sqlc-v1.7.0-linux-amd64/sqlc init -f sqlc.yaml
test:
	go test -v -cover ./...
.PHONY: createdb dropdb postgres migratedown migrateup sqlccompile sqlcgenerate sqlcinit