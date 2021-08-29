# mini-pos-app

Mini Pos APP

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Docker-compose](https://docs.docker.com/compose/) for running MySQL Database.

## Setup

Build Database Environment Container

```bash
docker-compose up
```

## Run the service

Prepare necessary environemt by rename .env.example to .env

```bash
APP_HOST=localhost
APP_PORT=7000
DB_HOST=localhost
DB_PORT=7010
DB_USER=root
DB_PASSWORD=mypass
DB_NAME=minipos
```

Get Go packages

```bash
go get .
```

Then run the proggram

```bash
go run cmd/main.go
```

## Documentation

- [Database Scheme](https://dbdiagram.io/d/612b12a3825b5b0146e93d14) 
- [API Documentation](https://www.postman.com/danisbagus/workspace/miniposapp/request/8996756-218551d2-1532-4ad3-a28d-01b19a144870)
