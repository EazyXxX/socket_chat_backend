#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

echo "PGPASSWORD=$DB_PASSWORD psql -h \"$host\" -U \"postgres\" -c '\q'; do"

until PGPASSWORD=postgres psql -h "db" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 2
done

>&2 echo "Postgres is up - executing command"
exec $cmd