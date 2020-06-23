# go-xo-practice

## Prerequisite

See https://github.com/kyanny/postgresql-tutorial-sample-database-docker for PostgreSQL setup.

Note that this project uses PostgreSQL 11 because `xo` seems not working for PostgreSQL 12.

```
$ docker-compose up -d
```

## Run `xo`

```
$ mkdir models
$ xo "pgsql://postgres:@localhost/dvdrental?sslmode=disable" -o models
```
