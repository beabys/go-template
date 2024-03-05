# go-http-template

## Requirements

* [docker and docker-compose for local usage / development](https://www.docker.com/get-started)
* [MariaDB](https://mariadb.org/) or [MySQL](https://www.mysql.com/) (optional only for run on host machine) 
* [Redis](https://redis.io/download//) (optional only for run on host machine) 

## Project structure

```
go-http-template
├── builds
├── cmd
│   └── server              # Main File(s)
├── config.yaml             # configuration file used in development
├── deployment              # k8s deployment files
├── internal
│   ├── api                 # api handler implementation
│   │   └── v1              # openapi generated files
│   ├── app                 # internal application 
│   │   ├── config          # structs and configuration implementation
│   │   └── database        # database clients
│   ├── domain              # domain folder for entities
│   ├── hello_world         # hello_world domain service implementation
│   ├── utils               # helper functions
│   └── mocks               # mockery generated files
├── migrations              # sql migration files
├── openapi.yaml            # openapi
├── pkg                     # packages expeted to be separated in the future
└── proto                   # Proto contracts and definitions
```

## Configurations

First copy the template file `./config/config.yaml.sample` to `./config/config.yaml`

```
cp ./config.yaml.sample ./config.yaml

```

Copy the env template file `.env.local` to `.env`

```
cp .env.local .env

```

Edit the corresponding fields on each file if required


*Note: The make script will bind .env variables on each run as a ENVIRONMENT variables, for production is required to inject those env variables into the host

## Commands

        make up            - run build and docker-compose up
        make down          - run docker-compose down
        make run           - run the app locally
        make mysql-up      - run docker-compose up of the mysql showing the logs
        make mysql-down    - run docker-compose down for mysql
        make redis-up      - run docker-compose up of the redis showing the logs
        make redis-down    - run docker-compose down for redis
        make mockery       - run the mockery command to create mock interfaces(requires mockery installed)
        make unit          - run Unit test of the app
        make unit-coverage - run Unit test and generating html report
        gen-api-v1         - generate v1 code from apenapi.yaml file
        gen-api-doc        - generate api html documentation from apenapi.yaml file
        proto-gen          - genetate proto contracts in php and go

## Quick start running Local with Docker

Assume you already installed docker.

For the first time to start the service, just run `make up` this command will start the service at port `80` by default, if need to change the port of the application change it on the `./config.yaml` file.


## Run by local golang

If you do not install docker but you have installed golang and `GO111MODULE` enabled, Mysql and Redis

You can run `make run` to run by local golang.

*Note: this requires a local instance of mysql and redis

Created By Alfonso Rodriguez(beabys@gmail.com)
