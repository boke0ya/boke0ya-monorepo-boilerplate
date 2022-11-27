.PHONY: docker front dbup dbdown

docker:
	docker-compose up

front:
	cd front; yarn dev

setup:
	docker-compose build --no-cache
	cd front; yarn

dbup:
	migrate -database postgres://develop:password@localhost:5432/develop?sslmode=disable -source file://./api/migrations up

dbdown:
	migrate -database postgres://develop:password@localhost:5432/develop?sslmode=disable -source file://./api/migrations down
