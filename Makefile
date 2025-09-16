ifndef PRJ_DIR
	PRJ_DIR := $(PWD)../../
	export
endif
.SILENT: up down logs restart

up:
	docker compose up -d

down:
	docker compose down --remove-orphans

logs:
	docker compose logs -f api

restart:
	docker compose restart

migrate:
	docker compose exec -it api sh script/migrate.sh