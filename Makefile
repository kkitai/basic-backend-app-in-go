SRC = main.go

init:
	$(eval APP_NAME := $(shell basename $(CURDIR)))

build: init
	go build -o $(APP_NAME) $(SRC)

test:
	go test ./...

