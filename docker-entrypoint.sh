#!/bin/sh
set -e

while true ; do
	nc -z ${TODO_DB_HOST} ${TODO_DB_PORT} 2>&1 && break
	echo "Waiting for postgress..."
	sleep 3
done

exec "$@"
