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
	mockery --dir=internal --all --recursive --keeptree --output=./internal/mocks && go mod tidy

.PHONY: unit
unit:
	go mod tidy && go test -race ./... -v -coverprofile .testCoverage.txt

.PHONY: unit-coverage
unit-coverage: unit ## Runs unit tests and generates a html coverage report
	go tool cover -html=.testCoverage.txt -o unit.html
