DB_URL=postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable

npm:
	sudo apt-get install npm
	sudo npm -g install create-react-app

config:
	curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
	sudo apt-get update
	sudo apt-get install -y migrate
	sudo rm -rf /usr/local/go
	sudo tar -C /usr/local/ -xzf go1.22.0.linux-amd64.tar.gz
	make network
	make docker
	make createdb
	make migrateup

docker:
	docker stop $$(docker ps -q | tail -n 1); \
	docker rm $$(docker ps -a --filter "status=exited" --format "{{.ID}}"); \
	make postgres
	docker start $$(docker ps -a --format "{{.ID}}" | head -n 1)

network:
	docker network create reddit-network

postgres:
	docker run --name postgres --network reddit-network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:16.2

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=postgres -d mysql:8

createdb:
	docker exec -it postgres createdb --username=postgres reddit || \
	(sleep 10 && docker exec -it postgres createdb --username=postgres reddit)

dropdb:
	docker exec -it postgres dropdb -U postgres reddit

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=reddit \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis
