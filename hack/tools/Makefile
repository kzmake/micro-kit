# https://github.com/kubernetes/kubernetes/blob/v1.19.0-beta.0/hack/tools/Makefile

export GO111MODULE=on

GOBIN ?= $(shell go env GOPATH)/bin

TOOLS= \
	$(GOBIN)/buf \
	$(GOBIN)/protoc-gen-buf-check-breaking \
	$(GOBIN)/protoc-gen-buf-check-lint \
	$(GOBIN)/protoc-gen-validate \
	$(GOBIN)/protoc-gen-go \
	$(GOBIN)/protoc-gen-micro \
	$(GOBIN)/golangci-lint \
	$(GOBIN)/goimports \
	$(GOBIN)/bazelisk \
	$(GOBIN)/gazelle \

.PHONY: all
all: uninstall clean install

.PHONY: uninstall
uninstall:
	rm -f $(TOOLS)

.PHONY: clean
clean:
	rm -rf vendor/

.PHONY: install
install: $(TOOLS) clean

.PHONY: deps
deps:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	go mod vendor

# protoc tools
$(GOBIN)/buf: deps
	go install github.com/bufbuild/buf/cmd/buf

$(GOBIN)/protoc-gen-buf-check-breaking: deps
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-check-breaking

$(GOBIN)/protoc-gen-buf-check-lint: deps
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-check-lint

$(GOBIN)/protoc-gen-go: deps
	go install github.com/golang/protobuf/protoc-gen-go

$(GOBIN)/protoc-gen-validate: deps
	go install github.com/envoyproxy/protoc-gen-validate

$(GOBIN)/protoc-gen-micro: deps
	go install github.com/micro/micro/v2/cmd/protoc-gen-micro

# lint tools
$(GOBIN)/golangci-lint: deps
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

# fmt tools
$(GOBIN)/goimports: deps
	go install golang.org/x/tools/cmd/goimports

# bazel tools
$(GOBIN)/bazelisk: deps
	go install github.com/bazelbuild/bazelisk

$(GOBIN)/gazelle: deps
	go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle
