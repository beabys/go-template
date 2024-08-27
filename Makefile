SHELL:=/bin/bash

PROJECT_PATH := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
TEST_OUTPUT_PATH := $(PROJECT_PATH)/.output

FILE=./.env
ifneq ("$(wildcard $(FILE))","")
	include $(FILE)
	export $(shell sed 's/=.*//' $(FILE))
endif

.PHONY: run
run:
	go mod tidy && go run cmd/server/main.go
	
.PHONY: up
up: 
	docker build -t test-server -f builds/app/Dockerfile . && docker-compose -f ./builds/docker-compose.yml up --build -d

.PHONY: down
down:
	docker-compose -f ./builds/docker-compose.yml down

# redis
.PHONY: redis-up
redis-up: 
	docker-compose -f ./builds/redis/docker-compose.yml up --build -d

.PHONY: redis-down
redis-down:
	docker-compose -f ./builds/redis/docker-compose.yml down

# mysql
.PHONY: mysql-up
mysql-up: 
	docker-compose -f ./builds/mysql/docker-compose.yml up --build -d

.PHONY: mysql-down
mysql-down:
	docker-compose -f ./builds/mysql/docker-compose.yml down

.PHONY: mockery
mockery:
	mockery && go mod tidy

.PHONY: unit
unit:
	go mod tidy && go test $(shell go list ./internal/... | grep -v /mocks) -race -coverprofile .testCoverage.txt

.PHONY: unit-coverage
unit-coverage: unit ## Runs unit tests and generates a html coverage report
	go tool cover -html=.testCoverage.txt -o unit.html

.PHONY: gen-api
gen-api: ## generates public api interfaces
	docker container run --rm -v $(PWD):/app golang:1.22.0 sh -c "go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest && \
    cd /app && mkdir -p ./internal/api/v1 && \
    oapi-codegen --package=v1 --generate="types,client,spec,chi-server,skip-prune" ./openapi.yaml | sed 's/V1/v1/g' | sed 's/Id$(WORD_END)/ID/g' | sed 's/Guid/GUID/g' | sed 's/Sku/SKU/g' | sed 's/Qoh/QOH/g' | sed 's/float32/float64/g' | sed 's/Url/URL/g' > ./internal/api/v1/v1.go"

.PHONY: gen-api-doc
gen-api-doc: ## generates public api document
	docker run --rm -v $(PWD):/app -w /docs node:hydrogen-slim sh -c "npm i -g @redocly/cli@latest && \
	redocly build-docs -o /app/docs/index.html /app/openapi.yaml"

.PHONY: proto-gen
proto-gen:
	docker run -w /proto/defs --rm  -v ${CURDIR}/proto:/proto ealves/buf:0.1.0-rc4 generate
