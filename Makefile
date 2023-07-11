.PHONY: run
run:
	go run cmd/main.go

.PHONY: up-db
up-db:
	docker-compose up -d