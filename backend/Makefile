ENV ?= local
export ENV

run:
	docker build -f dev/Dockerfile.godev -t godev .
	docker-compose up --build --remove-orphans

db/up:
	@utils/db_run_migrations.sh $(env) $(service) up

db/down:
	@utils/db_run_migrations.sh $(env) $(service) down

db/connect:
	@utils/db_connect.sh $(env) $(service)

db/version:
	@utils/db_migration_version.sh $(env) $(service)

