ENV ?= local
export ENV

run:
	docker build -f dev/Dockerfile.godev -t godev .
	docker-compose up --build --remove-orphans

db/up:
	@scripts/db_migrations.sh $(env) $(service) up

db/down:
	@scripts/db_migrations.sh $(env) $(service) down
