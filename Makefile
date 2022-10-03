init:
	$(eval APP_NAME := $(shell echo "hoge"))

build: init
	go build -o $(APP_NAME) ./main.go
