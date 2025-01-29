#!/bin/sh
set -x #echo on
#set +x #echo off

source ./docker/local_vars.sh
source ./docker/private_local_vars.sh

docker-compose up db -d
docker-compose logs db -f
docker compose down db
