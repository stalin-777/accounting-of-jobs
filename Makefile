run:
	go run cmd/aoj/main.go
upbuild:
	docker-compose up -d --build aoj-srv
up:
	docker-compose up -d aoj-srv
build:
	docker build -t aoj-srv .
prune:
	docker container prune
dockerrun:
	docker run --name=aoj-srv -p 8088:8088 --rm aoj-srv
down:
	docker-compose down aoj-srv
pgrun:
	docker run --rm --name postgres-aoj -e POSTGRES_PASSWORD=qwerty -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -d -p 5432:5432 postgres

.PHONY: run upbuild up build prune dockerrun down pgrun