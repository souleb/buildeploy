PROJECTNAME = $(shell basename "$(PWD)")
DBNAME = $(PROJECTNAME)db

# Go related variables.
GOBASE = $(shell pwd)
GOBIN = $(GOBASE)/bin
GOFILES = $(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR = /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will store the server process id when it's running on development mode
PID = /tmp/.$(PROJECTNAME)-api-server.pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

REGISTRY ?= gcr.io/images
COMMIT_SHA = $(shell git rev-parse --short HEAD)

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o $(APP) ./cmd/main.go

.PHONY: run
## run: runs go run main.go
run:
	go run -race ./cmd/main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@go clean

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy
	
# helper rule for deployment
check-environment:
ifndef ENV
    $(error ENV not set, allowed values - `development`, `staging` or `production`)
endif

.PHONY: docker-build
## docker-build: builds the project's docker image
docker-build: build
	docker build -t $(PROJECTNAME):$(COMMIT_SHA) .

.PHONY: docker-push
## docker-push: pushes the project's docker image to registry
docker-push: check-environment docker-build
	docker push $(REGISTRY)/$(ENV)/$(PROJECTNAME):$(COMMIT_SHA)

.PHONY: help
## help: prints this help message
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: gen
## gen: generate all proto files
gen:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src ./proto/workflow/v1/workflow.proto
	protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
	./proto/workflow/v1/workflow.proto

.PHONY: docker-postgres
## docker-postgres: launch a new postgres server with a default db set to project's name
docker-postgres:
	@docker pull postgres
	@mkdir -p $(HOME)/docker/volumes/postgres
	@docker run --rm --name pg-docker -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=$(DBNAME) \
	-d -p 5432:5432 \
	-v $(HOME)/docker/volumes/postgres:/var/lib/postgresql/data  postgres

.PHONY: migrateup
migrateup:
    migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/$(DBNAME)?sslmode=disable" -verbose up

.PHONY: migratedown
migratedown:
    migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/$(DBNAME)?sslmode=disable" -verbose down