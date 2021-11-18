# Guide

***
##Postgres DB

Docker

- `docker run --name pg-cont -e POSTGRES_PASSWORD=1234 -d -p 5432:5432 postgres`

Create DB

- `$docker exec -it pg-cont bash` to enter into container
- `$su postgres` to login as the default db user
- `$createdb testdb` to the database
- `$psql` to connect to DB engine
- `\l` to list databases
- `\c testdb` to select database. Then run the sql in the data.setup.sql
- `\dt` To list the database tables.