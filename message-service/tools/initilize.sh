#!/bin/bash

CONTAINER_NAME="postgres"
DB_USER="postgres"
DB_PASSWORD="test123456"

docker exec -it $CONTAINER_NAME psql -U $DB_USER -c "CREATE DATABASE messages;"
docker exec -it $CONTAINER_NAME psql -U $DB_USER -d messages -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'


