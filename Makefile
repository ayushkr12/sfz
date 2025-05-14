default: build

APP_NAME := "sfz"
CMD_DIR := "cmd/$(APP_NAME)/main.go"

.PHONY: run
run:
	go run $(CMD_DIR) $(ARGS)

.PHONY: build
build:
	go build -o $(APP_NAME) $(CMD_DIR)

.PHONY: bin
bin:
	mv $(APP_NAME) ~/go/bin/$(APP_NAME)
