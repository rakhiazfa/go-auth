.PHONY: build

create-migration:
	migrate create -ext sql -dir migrations -seq $(name)

database-seed:
	go run cmd/seed/main.go

generate-wire:
	wire gen ./wire

build:
	go build -o bin/api cmd/api/main.go

dev:
	go run cmd/api/main.go

start:
	./bin/api