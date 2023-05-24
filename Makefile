.PHONY: build run run-prd run-dev tests lint vet fmt create-env-file create-db destroy-env

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

# run: starts application and it's dependencies in local environment
run:
	@echo "===> Starting local environment!\n"
	@docker-compose up -d db
	@sleep 5 # Just some time to wait before the database is ready for the next command
	@make create-db
	@docker-compose up

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
	@docker-compose run \
		--entrypoint="" \
		app \
		sh -c "go test -v -timeout 500ms -cover ./..."
	@echo "\n===> Done!"

# lint: run GolangCI-Lint on the codebase
lint:
	@echo "===> Started GolangCI-Lint execution!\n"
	@docker run \
		-ti --rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:latest \
		golangci-lint run -v
	@echo "\n===> Done!"

# vet: execute 'go vet' on the codebase
vet:
	@echo "===> Started 'go vet' execution!\n"
	@docker-compose run \
		--entrypoint="" \
		app \
		sh -c "go vet ./..."
	@echo "\n===> Done!"

# fmt: execute 'go fmt' on the codebase
fmt:
	@echo "===> Started 'go fmt' execution!\n"
	@docker-compose run \
		--entrypoint="" \
		app \
		sh -c "go fmt ./..."
	@echo "\n===> Done!"

# Create a ./env/.env file if it not exists, based on ./env/.env.sample
create-env-file:
	@cp -u -p ./env/.env.sample ./env/.env

# create-db: creates the project's database
create-db:
	@echo "===> Creating database, please wait...\n"
	@docker-compose exec -T db \
		psql < ./env/db/up.sql -U $(shell docker-compose exec db bash -c 'printenv -0 POSTGRES_USER')
	@echo "\n===> Done!"

# destroy-env: stops and destroy Docker Compose environment
destroy-env:
	@echo "===> Nuking environment, please wait...\n"
	@docker-compose kill -s SIGKILL
	@docker-compose rm -v -f
	@echo "\n===> Done, environment is destroyed!"
