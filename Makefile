# Usage:
# make help -- to show help

.PHONY = help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

app: ## Build applicatiom.
	@echo "Build applicatiom."
	go build -v -o bin/sensor ./sensor

unittest: ## Build and run golang unittest.
	@echo "Building go unittets"
	go test -cover ./...

run: ## Build and run local sensor with fake sensor.
	@echo "Build and run local sensor with fake sensor"
	CONFIG=./testdata DEVICE=./fejksensor/ go run ./sensor
