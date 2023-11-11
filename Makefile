# One file to rule to them all

ifndef PROJECT_NAME
PROJECT_NAME := petme
endif

ifndef PRODUCTION_ENVIRONMENT:
PRODUCTION_ENVIRONMENT := prod
endif

ifndef DOCKER_BIN:
DOCKER_BIN := docker
endif

ifndef DOCKER_COMPOSE_BIN:
DOCKER_COMPOSE_BIN := docker-compose
endif

# Initialize the config for your local copy of the repo
init:
	echo "machine https://github.com/vunquitk11\nlogin username@spdigital.sg\npassword personalaccesstoken" > build/.netrc

build-local-go-image:
	${DOCKER_BIN} build -f build/local.go.Dockerfile -t ${PROJECT_NAME}-go-local:latest .
	-${DOCKER_BIN} images -q -f "dangling=true" | xargs ${DOCKER_BIN} rmi -f

# ----------------------------
# Project level Methods
# ----------------------------
teardown:
	${COMPOSE} down -v
	${COMPOSE} rm --force --stop -v

setup: api-setup
migrate: api-pg-migrate

# ----------------------------
# api Methods
# ----------------------------
API_COMPOSE = ${COMPOSE} run --name ${PROJECT_NAME}-api-$${CONTAINER_SUFFIX:-local} --rm --service-ports -w /api api
ifdef CONTAINER_SUFFIX
api-test: api-setup
endif
api-test:
	${API_COMPOSE} sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."
api-run:
	${API_COMPOSE} sh -c "go run -mod=vendor cmd/serverd/*.go"
api-pg-migrate:
	${COMPOSE} run --rm pg-migrate sh -c './migrate -path /api-migrations -database $$PG_URL up'
api-pg-drop:
	${COMPOSE} run --rm pg-migrate sh -c './migrate -path /api-migrations -database $$PG_URL drop'
api-pg-redo: api-pg-drop api-pg-migrate

ifdef CONTAINER_SUFFIX
api-setup: volumes pg sleep api-pg-migrate
else
api-setup: pg sleep api-pg-migrate
api-setup:
	${DOCKER_BIN} image inspect ${PROJECT_NAME}-go-local:latest >/dev/null 2>&1 || make build-local-go-image
endif


# ----------------------------
# Base Methods
# ----------------------------
volumes:
	${COMPOSE} up -d alpine
	${DOCKER_BIN} cp ${shell pwd}/api/. ${PROJECT_NAME}-alpine-$${CONTAINER_SUFFIX:-local}:/api
	${DOCKER_BIN} cp ${shell pwd}/api/data/migrations/. ${PROJECT_NAME}-alpine-$${CONTAINER_SUFFIX:-local}:/api-migrations

COMPOSE := PROJECT_NAME=${PROJECT_NAME} ${DOCKER_COMPOSE_BIN} -f build/docker-compose.base.yaml
ifdef CONTAINER_SUFFIX
COMPOSE := ${COMPOSE} -f build/docker-compose.ci.yaml -p ${CONTAINER_SUFFIX}
else
COMPOSE := ${COMPOSE} -f build/docker-compose.local.yaml
endif

pg:
	${COMPOSE} up -d pg

sleep:
	sleep 5