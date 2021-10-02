#!/bin/sh

set -o errexit

dir=$(dirname "$0")

random_port() {
    # Using IANA suggested ephemeral port range 49152 - 65535
    awk 'BEGIN{srand('$$');print(int(49152 + rand() * 16383))}'
}

usage() {
    echo "Usage: $0 <env> <service> <command>

where <env> is:
    local      -- To run migrations in localhost instance.

<service> is the name of a  service backed by a Postgres database (e.g.: offset, service1, service2)

and <command> is a valid goose command"
    exit 1
}

if [ "$#" -ne 3 ]; then
  echo "wasdasd" $@
    usage
fi

port=$(random_port)
env="$1"
service="$2"
direction="$3"

case $env in
    local)
        port=$(docker-compose port $service-db 5432 | sed -e 's/.*://')
        if ! nc -z localhost $port
        then
            echo "localhost $port is not open. See README.md for more information."
            exit 1
        fi
        PGUSER=dbuser
        PGPASSWORD=dbpass
        ;;

    *)
        usage
        ;;
esac

echo "Running $service migrations to database in $env"
export PGUSER PGPASSWORD
goose -dir "./db/migrations" postgres "postgres://localhost:$port/$service?sslmode=disable" "$direction"
