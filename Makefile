.PHONY: docker front

docker:
	docker-compose up

front:
	cd front; yarn dev

setup:
	docker-compose build --no-cache
	cd front; yarn
