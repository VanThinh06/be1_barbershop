postgres:   
				docker run --name postgres1 -p 5432:5432 -e POSTGRES_USER=mtt16 -e POSTGRES_PASSWORD=Vanthinh11 -e POSTGRES_DB=BarberShop -d postgres:latest

createdb:
				docker exec -it postgres1 createdb --username=mtt16 --owner=mtt16 barbershop

dropdb:
				docker exec -it postgres1 dropdb barbershop

migratecreate:
				migrate create -ext sql -dir db/migration -seq barbershop
				
migrateup:
				migrate -path db/migration -database "postgresql://mtt16:Vanthinh11@localhost:5432/BarberShop?sslmode=disable" -verbose up

migratedown:
				migrate -path db/migration -database "postgresql://mtt16:Vanthinh11@localhost:5432/BarberShop?sslmode=disable" -verbose down

mock: 
				mockgen -package mockdb  -destination db/mock/store.go  barbershop/db/sqlc StoreMain

test:   
				go test -v -cover ./...

docker: 
				docker build

dropcontainer: 
				docker exec -it postgres1 dropdb barbershop

sqlc:
				docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

proto:
				protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
				--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
				--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=barber_shop \
				proto/*.proto
				@REM statik -src=./docs/swagger -dest=./docs

evans:
				evans --host 192.168.1.15 --port 9999 -r repl 

.PHONY: postgres createdb dropdb migrateup migratedown test sqlc proto evans