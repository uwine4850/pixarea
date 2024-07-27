#!/bin/bash

# sudo docker exec -i mysql mysql --defaults-extra-file=/schema/mysql.cnf -e "CREATE DATABASE IF NOT EXISTS pixarea;"
# sudo docker exec -i mysql mysql --defaults-extra-file=./schema/mysql.cnf < ./schema/auth.sql
# sudo docker exec -i mysql mysql --defaults-extra-file=./schema/mysql.cnf < ./schema/publication.sql

sudo docker-compose run --rm node npm install