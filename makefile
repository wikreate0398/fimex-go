PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/.bin
CMD_DIR = $(PROJECT_DIR)/cmd
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
APP_NAME = app
PATH := $(PROJECT_BIN):$(PATH)

#git
COMMIT ?= "fix"
BRANCH ?= "dev"

#lint
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

push:
	git add .
	git commit -m "$(COMMIT)"
	git push origin $(BRANCH)

run:
	go run ./cmd/main.go

act:
	act --container-architecture linux/amd64 --secret-file ./.github/.secrets

#BUILD
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(PROJECT_BIN)/$(APP_NAME) $(CMD_DIR)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(PROJECT_BIN)/$(APP_NAME).exe $(CMD_DIR)

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(PROJECT_BIN)/$(APP_NAME) $(CMD_DIR)

clean:
	rm -rf $(PROJECT_BIN)

#LINTER
.PHONY: .install-linter
.install-linter:
	### INSTALL GOLANGCI-LINT ###
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.63.4

.PHONY: lint
lint: .install-linter
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY: lint-fast
lint-fast: .install-linter
	$(GOLANGCI_LINT) run ./... --fast --config=./.golangci.yml
