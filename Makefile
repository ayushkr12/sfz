APP_NAME := "sfz"
CMD_DIR := "cmd/$(APP_NAME)/main.go"

.PHONY: run
run:
	go run $(CMD_DIR) $(ARGS)