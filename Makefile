PROJECT = github.com/slavamuravey/wbppc
GO_PACKAGES = $(shell go list $(PROJECT)/pkg/...)
PRIVATE_REPOS ?=
LINTER_VERSION ?= v1.32.2
BIN = wbppc
GOOS ?= linux
GOARCH ?= amd64
LDFLAGS = "-s -w"

.PHONY: all
all: compile

.PHONY: compile
compile: vendor test lint
	@echo "+ $@"
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -installsuffix cgo \
		-ldflags $(LDFLAGS) -o bin/$(GOOS)-$(GOARCH)/$(BIN) $(PROJECT)/cmd

.PHONY: vendor
vendor: bootstrap
	@echo "+ $@"
ifeq ("$(wildcard go.mod)","")
	@GOPRIVATE=$(PRIVATE_REPOS) go mod init $(PROJECT)
endif
	@GOPRIVATE=$(PRIVATE_REPOS) go mod tidy

.PHONY: test
test:
	@echo "+ $@"
	@go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}}"go test -race -cover {{.Dir}}"{{end}}' $(GO_PACKAGES) | xargs -L 1 sh -c

.PHONY: fmt
fmt:
	@echo "+ $@"
	@go list -f '"gofmt -w -s -l {{.Dir}}"' $(GO_PACKAGES) | xargs -L 1 sh -c

.PHONY: imports
imports: bootstrap
	@echo "+ $@"
	@go list -f '"goimports -w {{.Dir}}"' ${GO_PACKAGES} | xargs -L 1 sh -c

.PHONY: lint
lint: bootstrap
	@echo "+ $@"
	@GOPRIVATE=$(PRIVATE_REPOS) golangci-lint run ./...

HAS_LINT := $(shell command -v golangci-lint;)
HAS_IMPORTS := $(shell command -v goimports;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_LINT
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINTER_VERSION)
endif
ifndef HAS_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif