start-db:
	docker-compose up -d

stop-db:
	docker-compose down

create-db:
	docker-compose exec psql_database createdb --username=root --owner=root chat-app

drop-db:
	docker-compose exec psql_database dropdb chat-app

db-shell:
	docker-compose  exec -it psql_database psql --username=root -d chat-app


db-script:
	@echo "Enter the psql script you want to run:"
	@read script; \
	echo "Script: $$script"; \
	docker-compose exec -T psql_database psql --username=root -d chat-app -c "$$script"


tidy:
	go mod tidy

build: tidy
	go build -o bin/chat-app cmd/server/main.go

run: build
	./bin/chat-app

clean:
	rm -rf bin/*

.PHONY: start-db stop-db create-db drop-db db-shell db-script tidy build run clean
