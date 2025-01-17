MAIN_PACKAGE_PATH := .
BINARY_NAME := git-classrooms
APP_VERSION ?= $(shell git describe --tags --always --dirty)
APP_GIT_COMMIT ?= $(shell git rev-parse HEAD)
APP_GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
APP_GIT_REPOSITORY ?= https://github.com/git-classrooms/git-classrooms
APP_BUILD_TIME ?= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
define GOFLAGS
-ldflags=" \
  -s -w \
  -X main.version=$(APP_VERSION) \
"
endef
MOCKERY_VERSION := v2.42.2
#   -X github.com/git-classrooms/git-classrooms/internal/storage/local/info.version=$(APP_VERSION) \
#   -X github.com/git-classrooms/git-classrooms/internal/storage/local/info.gitCommit=$(APP_GIT_COMMIT) \
#   -X github.com/git-classrooms/git-classrooms/internal/storage/local/info.gitBranch=$(APP_GIT_BRANCH) \
#   -X github.com/git-classrooms/git-classrooms/internal/storage/local/info.gitRepository=$(APP_GIT_REPOSITORY) \
#   -X github.com/git-classrooms/git-classrooms/internal/storage/local/info.buildTime=$(APP_BUILD_TIME) \

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build                             Build"
	@echo "  generate                          Generate"
	@echo "  generate/client                   Generate client pkg"
	@echo "  setup                             Install dependencies"
	@echo "  setup/ci                          Install dependencies for CI"
	@echo "  clean                             Clean"
	@echo "  run                               Run"
	@echo "  run/dev                           Run development environment"
	@echo "  build/docker                      Build docker image"
	@echo "  run/docker                        Run docker container"
	@echo "  infra/up                          Run infrastructure in docker compose (postgres, (pgadmin) and mailpit)"
	@echo "  infra/stop                        Run infrastructure stop"
	@echo "  infra/down                        Run infrastructure down (delete)"
	@echo "  migrate/new name=<name>           Create new migration"
	@echo "  migrate/check                     Check and print outstanding migrations"
	@echo "  migrate/status                    Migrate status"
	@echo "  seed/up                           Seed up"
	@echo "  seed/reset                        Seed reset"
	@echo "  tidy                              Fmt and Tidy"
	@echo "  lint                              Lint"
	@echo "  test                              Test"
	@echo "  test/verbose                      Test verbose"
	@echo "  config                            TODO: Edit config"
	@echo "  debug                             Debug"


.PHONY: run/dev
run/dev: generate
	@echo "Starting development environment..."
	@concur || true

.PHONY: run
run: build
	@echo "Running binary..."
	@./bin/$(BINARY_NAME)

.PHONY: build
build: generate
	@echo "Building binary..."
	@CGO_ENABLED=0 go build $(GOFLAGS) -o ./bin/$(BINARY_NAME) $(MAIN_PACKAGE_PATH)

.PHONY: setup
setup:
	# Backend
	@echo "Setting up environment..."
	go install github.com/akatranlp/concur@latest
	go install github.com/air-verse/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)
	go install github.com/golangci/golangci-lint@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	# TODO: maybe go install github.com/mikefarah/yq/v4@latest
	go mod download
	# TODO: with go 1.24 we can simply use go install tool

	# Frontend
	@cd frontend
	yarn install


.PHONY: setup/ci
setup/ci:
	@echo "Installing..."
	go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)
	go install github.com/swaggo/swag/cmd/swag@latest
	go mod download

.PHONY: setup/frontend
setup/frontend:
	@cd frontend
	yarn install

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin ./tmp
	@rm -f docs/docs.go docs/swagger.json docs/swagger.yaml
	@rm -rf $(shell yq '.packages | to_entries | map(.value.config.dir) | .[]' .mockery.yaml)
	@cd frontend && yarn clean

.PHONY: generate
generate:
	@echo "Generating code..."
	@go generate ./...
	$(MAKE) generate/client

.PHONY: generate/client
generate/client:
	@echo "Generating client..."
	@./script/swagger-codegen.sh

.PHONY: docker/build
docker/build:
	@echo "Running docker..."
	docker build -t $(BINARY_NAME)-docker \
		--build-arg APP_VERSION=$(APP_VERSION) \
		--build-arg APP_GIT_COMMIT=$(APP_GIT_COMMIT) \
		--build-arg APP_GIT_BRANCH=$(APP_GIT_BRANCH) \
		--build-arg APP_GIT_REPOSITORY=$(APP_GIT_REPOSITORY) \
		--build-arg APP_BUILD_TIME=$(APP_BUILD_TIME) \
		.

.PHONY: docker/run
docker/run:
	@echo "Running docker..."
	docker run -it --env-file .env --rm -p 3000:3000 $(BINARY_NAME)-docker

.PHONY: infra/up
infra/up:
	@docker compose -f docker-compose.local.yaml up -d

.PHONY: migrate/new
migrate/new:
	@echo "Migrating up..."
	@if [ -z "$(name)" ]; then \
		echo "error: name is required"; \
		echo "usage: make migrate/new name=name_of_migration"; \
		exit 1; \
	fi
	go run ./code_gen/goose/. normal create $(name) sql

.PHONY: migrate/status
migrate/status:
	@echo "Migrating status..."
	go run ./code_gen/goose/. normal status

.PHONY: seed/up
seed/up:
	@echo "Seeding up..."
	go run ./code_gen/goose/. seed -no-versioning up

.PHONY: seed/reset
seed/reset:
	@echo "Seeding reset..."
	go run ./code_gen/goose/. seed -no-versioning reset

.PHONY: migrate/check
migrate/check:
	go run ./code_gen/migrations/.

.PHONY: tidy
tidy:
	@echo "Tidying up..."
	go fmt ./...
	swag fmt --exclude frontend
	go mod tidy

.PHONY: lint
lint:
	@echo "Linting..."
	golangci-lint run

.PHONY: test
test:
	@echo "Testing..."
	go test -cover ./...

.PHONY: test/verbose
test/verbose:
	@echo "Testing..."
	go test -v -cover ./...

.PHONY: infra/logs
infra/logs:
	@docker compose -f docker-compose.local.yaml logs -n 10 -f

.PHONY: infra/stop
infra/stop:
	@docker compose -f docker-compose.local.yaml stop

.PHONY: infra/down
infra/down:
	@docker compose -f docker-compose.local.yaml down --volumes

.PHONY: debug
debug:
	@echo "Debugging..."
	dlv debug

.PHONY: config
config:
	@echo "TODO: Implement config with sops"
