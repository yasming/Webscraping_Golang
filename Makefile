#!make
include .env

start:
	docker-compose up

install-new-package:
	docker-compose run --rm server go get -u $(package_name)

run-all-tests:
	docker-compose run --rm server go test -coverprofile=coverage.out ./...

stop:
	docker-compose down