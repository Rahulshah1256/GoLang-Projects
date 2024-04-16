migrateup:
	migrate -path pck/db/migration -database "postgresql://postgres:Neebal@123@localhost:5432/golang_admission?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/db/migration -database "postgresql://postgres:Neebal@123@localhost:5432/golang_admission?sslmode=disable" -verbose down

sqlc:
	sqlc generate
protoc:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto
evens:
	.\evans --host localhost --port 8089 -r repl

.PHONY migrateup migratedown sqlc protoc