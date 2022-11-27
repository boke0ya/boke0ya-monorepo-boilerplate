# BeatHub

## Requirements
- Docker + docker-compose
- Node.js 18.12.1
- yarn
- make
- Go 1.19
- golang-migrate

## Setup
```sh
$ make setup
```

## Start developping
### Start docker
```sh
$ make docker
```

### Start frontend
```sh
$ make front
```

## Migrate DB
Install [golang-migrate](https://github.com/golang-migrate/migrate) first.

```sh
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Migrate up

```sh
$ make dbup
```

### Migrate down

```sh
$ make dbdown
```
