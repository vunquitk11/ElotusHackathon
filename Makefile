ifndef PROJECT_NAME
PROJECT_NAME := petme
endif

ifndef DOCKER_COMPOSE_BIN:
DOCKER_COMPOSE_BIN := docker-compose
endif

postgres:
	docker run --name petme -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	docker exec -it petme createdb --username=postgres --owner=postgres petmedb

dropdb:
	docker exec -it petme dropdb --username=postgres petmedb

migrateup:
	migrate -path data/migration --database "postgres://postgres:postgres@localhost:5432/petmedb?sslmode=disable" --verbose up

migratedown:
	migrate -path data/migration --database "postgres://postgres:postgres@localhost:5432/petmedb?sslmode=disable" --verbose down

gen-orm-models:
	sqlboiler --wipe psql && GOFLAGS="-mod=vendor" goimports -w repository/orm/*.go

.PHONY: postgres createdb dropdb migrateup migratedown gen-orm-models
