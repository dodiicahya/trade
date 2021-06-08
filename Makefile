#!/bin/bash

build: binary

binary:
	@echo "building binary.."
	@go build -tags static_all .


clean:
	@echo "cleaning ..."
	@rm -f go-trade
	@rm -rf vendor
	@rm -f go.sum


install:
	@echo "Installing dependencies...."
	@rm -rf vendor
	@rm -f Gopkg.lock
	@rm -f glide.lock
	@go mod tidy && go mod download && go mod vendor

test:
	@go test $$(go list ./... | grep -v /vendor/) -cover

test-cover:
	@go test $$(go list ./... | grep -v /vendor/) -coverprofile=cover.out && go tool cover -html=cover.out ; rm -f cover.out

coverage:
	@go test ./... --cover | awk '{if ($$1 != "?") print $$2 " " $$5;}' | sed 's/\%//g' | awk '{ print $$1 " | coverage: " $$2 "%"; sum += $$2; n++ } END { if (n > 0) printf "AVG coverage project directory = %.2f%%\n", sum/n }'

local:
	@docker-compose -f docker-compose.yml --project-directory . up --build