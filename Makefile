.PHONY: run
run:
	/usr/local/go/bin/go run cmd/main.go

.PHONY: up-db
up-db:
	docker-compose up -d