# Values used in the build steps
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date "+%Y-%m-%d-%T:%N_%Z")

# build: creates a Docker image ready for deployment
build:
	@echo "===> Creating Production-ready image, tag $(COMMIT_HASH)\n"
	@docker build \
		--no-cache \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--build-arg COMMIT_HASH="$(COMMIT_HASH)" \
		-t da-api:$(COMMIT_HASH) -f build/Dockerfile.prd .

# run-prd: starts the project in the deployment environment
run-prd:
	@echo "===> Starting deployment environment!\n"
	@docker run -ti --rm -p 8080:8080 da-api:$(COMMIT_HASH)

# run: build and start the deployment environment
run: build run-prd

# run-dev: starts the project in the local environment
run-dev: create-env-file create-db
	@echo "===> Building local environment. Loading, please wait...\n"
	@docker build \
		--no-cache \
		-t da-api:develop -f build/Dockerfile.dev .
	@echo "\n===> Starting local environment!\n"
	@docker run -ti --rm -p 8081:8080 da-api:develop

# tests: run all tests on the application
tests:
	@echo "===> Test execution started!\n"
	@docker run \
		-ti --rm \
		-v $(shell pwd):/app \
		-w /app \
		golang:1.20-alpine \
		sh -c "go test -v -timeout 300ms -coverprofile=testcov/cover.out ./... \
			&& go tool cover -html=testcov/cover.out -o=testcov/cover.html"

# lint: run GolangCI-Lint on the codebase
lint:
	@echo "===> Started GolangCI-Lint execution!\n"
	@docker run \
		-ti --rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:latest \
		golangci-lint run -v

# vet: execute 'go vet' on the codebase
vet:
	@echo "===> Started 'go vet' execution!\n"
	@docker run \
		-ti --rm \
		-v $(shell pwd):/app \
		-w /app \
		golang:1.20-alpine \
		go vet ./...

# fmt: execute 'go fmt' on the codebase
fmt:
	@echo "===> Started 'go fmt' execution!\n"
	@docker run \
		-ti --rm \
		-v $(shell pwd):/app \
		-w /app \
		golang:1.20-alpine \
		go fmt ./...

# Create a ./env/.env file if it not exists, based on ./env/.env.sample
create-env-file:
	@cp -u -p ./env/.env.sample ./env/.env

# create-db: creates the project's database
create-db:
	@echo "===> Creating database, please wait...\n"
	@docker-compose exec -T db \
		psql < ./env/db/up.sql -U $(shell docker-compose exec db bash -c 'printenv -0 POSTGRES_USER')

# destroy-env: stops and destroy Docker Compose environment
destroy-env:
	@echo "===> Nuking environment, please wait...\n"
	@docker-compose kill -s SIGKILL
	@docker-compose rm -v -f
	@echo "\n===> Done, environment is destroyed!"


.PHONY: build run run-prd run-dev tests lint vet fmt create-env-file create-db destroy-env
