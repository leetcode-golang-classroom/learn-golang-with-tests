.PHONY=build
# include .env
# export $(shell sed 's/=.*//' .env)

coverage:
	@go test -v -cover ./...

test:
	@go test -v ./...
