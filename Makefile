postgres:   
    docker run --name postgres1 -p 8888\:5432 -e POSTGRES_USER=mtt16 -e POSTGRES_PASSWORD=Vanthinh11 -e POSTGRES_DB=BarberShop -d postgres:latest

createdb:
    docker exec -it postgres1 createdb --username=mtt16 --owner=mtt16 barbershop

# dropdb:
#     docker exec -it postgres1 dropdb barbershop

migrateup:
    migrate -path db/migration -database "postgresql\://mtt16\:Vanthinh11@localhost:8888/barbershop?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migration -database "postgresql\://mtt16\:Vanthinh11@localhost:8888/barbershop?sslmode=disable" -verbose down

# test:
#     go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test