.PHONY: up down logs ps restart rebuild

up:
	docker compose up -d

down:
	docker compose down

ps:
	docker compose ps

restart:
	docker compose down && docker compose up -d

rebuild:
	docker compose build --no-cache

books:
	docker compose up -d books-db books-service

users:
	docker compose up -d users-db

loans:
	docker compose up -d loans-db loans-service


