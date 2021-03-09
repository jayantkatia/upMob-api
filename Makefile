postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine 
createdb:
	docker exec -it postgres13 createdb --username=root --owner=root upcoming_mobiles
dropdb:
	docker exec -it postgres13 dropdb upcoming_mobiles
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/upcoming_mobiles?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/upcoming_mobiles?sslmode=disable" -verbose down
dockersql:
	docker exec -it postgres13 psql -U root -d upcoming_mobiles
sqlcgenerate:
	sqlc generate
sqlccompile:
	sqlc compile
sqlcinit:
	sqlc init -f sqlc.yaml
test:
	go test -v -cover ./...
.PHONY: createdb dropdb postgres migratedown migrateup sqlccompile sqlcgenerate sqlcinit