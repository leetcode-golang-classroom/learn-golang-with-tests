.PHONY=build
# include .env
# export $(shell sed 's/=.*//' .env)

coverage:
	@go test -v -cover ./...

test:
	@go test -v ./...

coverage-test:
	@go test -cover github.com/leetcode-golang-classroom/learn-golang-with-tests/arrays-and-slices -run ^TestSumAllTails$