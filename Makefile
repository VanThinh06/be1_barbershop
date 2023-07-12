postgres:   
    docker run --name postgres1 --network barbershop_network  -p 5432:5432 -e POSTGRES_USER=mtt16 -e POSTGRES_PASSWORD=Vanthinh11 -e POSTGRES_DB=BarberShop -d postgres:latest

createdb:
    docker exec -it postgres1 createdb --username=mtt16 --owner=mtt16 barbershop

dropdb:
    docker exec -it postgres1 dropdb barbershop

migratecreate:
    migrate create -ext sql -dir db/migration -seq add_filed_users
    
migrateup:
    migrate -path db/migration -database "postgresql://mtt16:Vanthinh11@localhost:5432/barbershop?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migration -database "postgresql://mtt16:Vanthinh11@localhost:5432/barbershop?sslmode=disable" -verbose down

test:
    go test -v -cover ./...

mock: 
    mockgen -package mockdb  -destination db/mock/store.go  barbershop/db/sqlc StoreMain

docker: 
    docker build

dropcontainer: 
    docker exec -it postgres1 dropdb barbershop



.PHONY: postgres createdb dropdb migrateup migratedown test