ifndef PROJECT_NAME
PROJECT_NAME := elotus_hackathon
endif

ifndef DOCKER_COMPOSE_BIN:
DOCKER_COMPOSE_BIN := docker-compose
endif

postgres:
	docker run --name elotus_hackathon -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	docker exec -it elotus_hackathon createdb --username=postgres --owner=postgres elotus

dropdb:
	docker exec -it elotus_hackathon dropdb --username=postgres elotus

migrateup:
	migrate -path data/migration --database "postgres://postgres:postgres@localhost:5432/elotus?sslmode=disable" --verbose up

migratedown:
	migrate -path data/migration --database "postgres://postgres:postgres@localhost:5432/elotus?sslmode=disable" --verbose down

gen-orm-models:
	sqlboiler --wipe psql && GOFLAGS="-mod=vendor" goimports -w repository/orm/*.go

.PHONY: postgres createdb dropdb migrateup migratedown gen-orm-models
