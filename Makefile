pkgs=$(shell go list ./... | grep -v /vendor/)

all: vet format style test build

.PHONY: style
style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

.PHONY: test
test:
	@echo ">> running tests"
	@go test -short $(pkgs)
	
.PHONY: format
format:
	@echo ">> formatting code"
	@go fmt $(pkgs)

.PHONY: vet
vet:
	@echo ">> vetting code"
	@go vet $(pkgs)

.PHONY: build
build: dep
	@echo ">> building binaries"
	@dep ensure -vendor-only
	@mkdir -p bin
	@CGO_ENABLED=0 go build -ldflags "-X main.version=`git rev-parse --short HEAD`" -o bin/terradays main.go

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u -v github.com/golang/dep/cmd/dep
endif

.PHONY: release
release: goreleaser
	goreleaser --rm-dist

.PHONY: goreleaser
goreleaser:
	go get github.com/goreleaser/goreleaser && go install github.com/goreleaser/goreleaser