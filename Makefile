.DEFAULT_GOAL := help
APP_ID_CLIENT := com.github.kaginawa.KVNCClient
APP_ID_AGENT := com.github.kaginawa.KVNCAgent

.PHONY: setup
setup: ## Resolve dependencies using Go Modules
	go mod download

.PHONY: clean
clean: ## Remove build artifact directory
	-rm -rfv build

.PHONY: test
test: ## Tests all code
	go test -cover -race ./...

.PHONY: lint
lint: ## Runs static code analysis
	command -v golint >/dev/null 2>&1 || { go get golang.org/x/lint/golint; }
	golint -set_exit_status ./...

.PHONY: run-client
run-client: ## Run client without build artifact generation
	go run ./cmd/kvnc-client

.PHONY: run-agent
run-agent: ## Run agent without build artifact generation
	go run ./cmd/kvnc-agent

.PHONY: build
build: ## Build executable binaries for local execution
	go build -ldflags "-s -w" -o build/kvnc-client ./cmd/kvnc-client
	go build -ldflags "-s -w" -o build/kvnc-agent ./cmd/kvnc-agent

.PHONY: build-linux
build-linux: ## Build linux package
	command -v fyne >/dev/null 2>&1 || { go get fyne.io/fyne/v2/cmd/fyne; }
	fyne package -os linux -icon icon.png -release -sourceDir ./cmd/kvnc-client -appID $(APP_ID_CLIENT)
	fyne package -os linux -icon icon.png -release -sourceDir ./cmd/kvnc-agent -appID $(APP_ID_AGENT)

.PHONY: build-mac
build-mac: ## Build mac package
	command -v fyne >/dev/null 2>&1 || { go get fyne.io/fyne/v2/cmd/fyne; }
	fyne package -os darwin -icon icon.png -release -sourceDir ./cmd/kvnc-client -appID $(APP_ID_CLIENT)
	fyne package -os darwin -icon icon.png -release -sourceDir ./cmd/kvnc-agent -appID $(APP_ID_AGENT)

.PHONY: build-win
build-win: ## Build windows package
	if not exist fyne go get fyne.io/fyne/v2/cmd/fyne
	fyne package -os windows -icon icon.png -release -sourceDir ./cmd/kvnc-client -appID $(APP_ID_CLIENT)
	fyne package -os windows -icon icon.png -release -sourceDir ./cmd/kvnc-agent -appID $(APP_ID_AGENT)

.PHONY: count-go
count-go: ## Count number of lines of all go codes
	find . -name "*.go" -type f | xargs wc -l | tail -n 1

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
