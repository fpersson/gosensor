# Usage:
# make help -- to show help

.PHONY = help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

app: ## Build applicatiom.
	@echo "Build applicatiom."
	go test -cover ./...

unittest: ## Build and run golang unittest.
	@echo "Building go unittets"
	go test -cover ./...

run: ## Build and run local sensor with fake sensor.
	@echo "Build and run local sensor with fake sensor"
	CONFIG=./testdata DEVICE=./fejksensor/ go run ./sensor

deploy-zero: ## Build and deploy on RPI zero, use DEST=<user>@<ip> to deploy
	@echo "Build and deploy on RPI zero"
	mkdir -p bin/arm
	GOOS=linux GOARCH=arm go build -o ./bin/arm/sensor ./sensor
	scp ./bin/arm/sensor $(DEST):/home/pi/bin
	scp -r ./templates $(DEST):/home/pi/bin

deploy-pi34: ## Build and deploy on RPI 3/4 use DEST=<user>@<ip> to deploy
	@echo "Build and deploy on RPI 3/4"
	mkdir -p bin/arm64
	GOOS=linux GOARCH=arm64 go build -o ./bin/arm64/sensor ./sensor
	scp ./bin/arm/sensor $(DEST):/home/pi/bin
	scp -r ./templates $(DEST):/home/pi/bin