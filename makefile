seeder:
	GO_ENV=development go run ./db/seeder.go

run-http:
	GO_ENV=development \
	GIN_MODE=debug \
	ELASTIC_APM_SERVICE_NAME=project \
	ELASTIC_APM_SERVICE_VERSION=0.0.1 \
	ELASTIC_APM_SECRET_TOKEN= \
	ELASTIC_APM_SERVER_URL=http://localhost:8200 \
	ELASTIC_APM_ENVIRONMENT=staging \
	go run main.go http

run-gql:
	GO_ENV=development \
	GIN_MODE=debug \
	ELASTIC_APM_SERVICE_NAME=project \
	ELASTIC_APM_SERVICE_VERSION=0.0.1 \
	ELASTIC_APM_SECRET_TOKEN= \
	ELASTIC_APM_SERVER_URL=http://localhost:8200 \
	ELASTIC_APM_ENVIRONMENT=staging \
	go run main.go gql

run-seeders:
	go run main.go seeders

run:
	GO_ENV=development \
	GIN_MODE=debug \
	ELASTIC_APM_SERVICE_NAME=project \
	ELASTIC_APM_SERVICE_VERSION=0.0.1 \
	ELASTIC_APM_SECRET_TOKEN= \
	ELASTIC_APM_SERVER_URL=http://localhost:8200 \
	ELASTIC_APM_ENVIRONMENT=staging \
	go run main.go http && go run main.go gql
	
test:
	GO_ENV=development GIN_MODE=debug go test ./pkg/... -v

test-cov:
	if ! [ -d "coverage" ]; then \
		echo "Creating covorage folder" ; \
		mkdir coverage; \
	fi
	GO_ENV=development go test ./... -cover -v -coverprofile ./coverage/c.out && go tool cover -html=./coverage/c.out -o ./coverage/coverage.html

build:
	go build -o bin/exec -ldflags "-s -w" main.go

docker-compose:
	docker-compose -f docker-compose.yml up --build -d