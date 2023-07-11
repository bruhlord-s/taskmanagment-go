# ✅ Task Managment App
API Server for task managment application, that will help developers to easily manage tasks for their projects.  Written in Go + Gin


*My pet-project to learn Go, tests, websockets, etc*

## Prerequisites
- Docker + docker-compose
- Go >=v1.20

## Installation
Create `.env` file in the root directory and paste database password, i.e
```
DB_PASSWORD=obpassword
```

Create `docker-compose` file from sample
```
$ cp docker-compose.yml.sample docker-compose.yml   // creates file from sample
```

Run
```
$ make up-db        // starts postgres container
$ make run          // starts webserver
```

If you don't have `make`, run
```
$ docker-compose up -d      // starts postgres container
$ go run cmd/main.go        // starts webserver
```

🎉 Now API Server is running at port `8080`

