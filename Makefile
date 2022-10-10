PID      = /tmp/noteappbackend-golang-project.pid
GO_FILES = $(wildcard *.go)
APP      = ./app
serve: restart
	@make build
	@watchman-make -p '**/*.go' 'Makefile*' -t build

stop:
	@docker-compose down

build:
	@go build .
	@./noteappbackend


.PHONY: serve restart kill before # let's go to reserve rules names
