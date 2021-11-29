# URL_Shorter
Application to create new shorter URL using GoLang.

## Environment Variables
The project require a file `/.env` to contain
following environment variables:
* `REDIS_HOST`
* `REDIS_PORT`
* `POSTGRES_HOST`
* `POSTGRES_PORT`
* `POSTGRES_PASSWORD`
* `POSTGRES_USER`
* `POSTGRES_DB`


## How to start
1. Clone the repository.
2. Go to repo folder.
3. `cd ./app`
4. `go get ./...` to install all requirements
5. `docker-compose up --build` to build and start project

## Configurations
In file `configurations.json` you can change DB which use and server port. Possible valies for db_name: `redis` or `postgreSQL`.

## Urls
1. To  create Short Url you have to call `/create_url` with body `{"url_origin": <string: url>}`. And you will get `shortLink`.
2. To redirect you just have to call `/{shortLink}`.

## Tests
To start test:
1. `cd ./app/tests`
2. `go tests -v`

## File System
```
├── app
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   └── url_controllers.go
│   ├── db
│   │   ├── db_universtal.go
│   │   ├── models.go
│   │   ├── postgresql.go
│   │   └── redis.go
│   ├── hasher
│   │   └── hasher.go
│   ├── migrations
│   │   └── 0001_create_url_table.sql
│   ├── tests
│   │   ├── hasher_test.sql
│   │   ├── postgresql_test.sql
│   │   └── redis_test.sql
│   ├── configurations.json
│   ├── go.mod
│   └── main.go
├── data
│   └── .gitignore
├── .env
├── .gitignore
├── docker-compose.yml
├── Dockerfile
└── README.md
```