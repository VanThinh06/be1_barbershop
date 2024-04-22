postgres:   
				docker run --name postgres1 -p 5432:5432 -e POSTGRES_USER=mtt16 -e POSTGRES_PASSWORD=Vanthinh11 -e POSTGRES_DB=BarberShop -d postgres:latest
networdcreate: 
				docker network create my-network

postgrespostgis:
				docker run --name my-postgres -p 5433:5432 --network my-network -e POSTGRES_USER=mtt16 -e POSTGRES_PASSWORD=Vanthinh11 -e POSTGRES_DB=BarberShop -d postgis/postgis
postgis:
				docker run --name postgis1 -e POSTGRES_PASSWORD=Vanthinh11 -d -p 5433:5432 postgis/postgis

createdb:
				docker exec -it postgres1 createdb --username=mtt16 --owner=mtt16 barbershop

dropdb:
				docker exec -it postgres1 dropdb barbershop

migratecreate:
				migrate create -ext sql -dir src/db/migration -seq add_column_barbershopservice
				
migrateup:
				migrate -path src/db/migration -database "postgres://mtt16:Vanthinh11@localhost:5433/BarberShop?sslmode=disable" up

migratedown:
				migrate -path src/db/migration -database "postgres://mtt16:Vanthinh11@localhost:5433/BarberShop?sslmode=disable" -verbose down

mock: 
				mockgen -package mockdb  -destination barbershop/src/db/mock/store.go  barbershop/src/db/sqlc StoreMain

test:   
				go test -v -cover ./...

docker: 
				docker build

dropcontainer: 
				docker exec -it postgres1 dropdb barbershop

sqlc:
				docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

sqlc_linux:
				sudo docker run --rm -v "$(pwd):/src" -w /src sqlc/sqlc generate

proto:
				protoc --proto_path=src/proto --go_out=src/pb --go_opt=paths=source_relative \
				--go-grpc_out=src/pb --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=src/pb --grpc-gateway_opt=paths=source_relative,allow_delete_body=true \
				--openapiv2_out=src/shared/doc/swagger --openapiv2_opt=allow_delete_body=true,allow_merge=true,merge_file_name=barbershop \
				src/proto/barber/*.proto
				statik -f -src=src/shared/doc/swagger -dest=src/shared/doc/ 

proto_customer:
				protoc --proto_path=src/proto --go_out=src/pb --go_opt=paths=source_relative \
				--go-grpc_out=src/pb --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=src/pb --grpc-gateway_opt=paths=source_relative \
				--openapiv2_out=src/shared/doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=barbershop \
				src/proto/customer/*.proto
				statik -f -src=src/shared/doc/swagger -dest=src/shared/doc/ 

evans:
				evans --host 192.168.1.15 --port 9999 -r repl 

.PHONY: postgres createdb dropdb migrateup migratedown test sqlc sqlc_linux proto proto_customer evans