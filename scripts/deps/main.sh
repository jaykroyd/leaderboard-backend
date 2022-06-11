#!/bin/bash

WHITE='\033[0m'
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'

set -e

case $ENV in
  test)
    export DOCKER_ENV='docker/test/docker-compose.yaml' ;;
  *)
    export DOCKER_ENV='docker/dev/docker-compose.yaml' ;;
esac

echo -e "${GREEN}=> Starting database${WHITE}"
docker-compose -f $DOCKER_ENV up -d postgres

until docker-compose -f $DOCKER_ENV exec -T postgres pg_isready; do
  echo -e "${BLUE}=> Waiting for Postgres...${WHITE}"
  sleep 1
done

echo -e "${GREEN}=> Adding the necessary extensions to the database${WHITE}"
docker-compose -f $DOCKER_ENV exec -T postgres psql -U user -d postgres \
  -c 'CREATE EXTENSION IF NOT EXISTS "pgcrypto";'
