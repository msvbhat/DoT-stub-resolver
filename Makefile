GO_FILES := $(shell find . -type f -name '*.go')
PKG_LIST := $(shell go list ./...)
BINARY := dot-resolver

.PHONY: all dep lint build docker clean

all: build

dep: ## Get the dependencies
	@go get -v -d ./...

lint: ## Run the golint
	@golint -set_exit_status ${PKG_LIST}

build: ## Build the static binary
	@env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -i -o ${BINARY}

docker: ## Build Docker image assuming static binary has been built
	@docker build -t dot-resolver .

clean: ## Remove the previous build
	@rm -f ${BINARY}
