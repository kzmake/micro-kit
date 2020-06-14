
SERVICE_DIR := service

SOURCES := $(shell find . \
	-type f -name '*.go' \
	-not -name '*.pb*.go' \
	-not -path "./vendor/*" \
	-not -path "./hack/tools/vendor/*")

$(warning SERVICE_DIR = $(SERVICE_DIR))

.DEFAULT_GOAL := help

.PHONY: tools
tools: ## 開発に必要なツールをインストールします
	@echo "\033[31m"
	@echo "$$ brew install protobuf"
	@echo "\033[0m"
	make -C hack/tools install

.PHONY: lint
lint: ## コードを検証します
	golangci-lint run

.PHONY: fmt
fmt: ## コードをフォーマットします
	@goimports -l -w $(SOURCES)

.PHONY: mod
mod: ## mod.go / mod.sum を整理します
	go mod tidy
	go mod vendor

.PHONY: bazel
bazel: ## bazel
	bazelisk run //:vendor
	bazelisk run //:gazelle -- update
	@echo "\033[31m"
	@echo "$$ bazelisk run //:vendor"
	@echo "$$ bazelisk run //:gazelle -- update"
	@echo "$$ bazelisk test //..."
	@echo "$$ bazelisk run //service/task/cmd/srv:srv"
	@echo "$$ bazelisk build //service/task/cmd/srv:srv"
	@echo "\033[0m"

.PHONY: proto
proto: ## protoファイルからgoファイルを生成します
	@for f in ${SERVICE_DIR}/**/interface/proto/*.proto; do \
		protoc \
		--proto_path=.:. \
		--proto_path=.:${GOPATH}/src \
		--proto_path=.:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=paths=source_relative:. \
		--micro_out=paths=source_relative:. \
		--validate_out=lang=go,paths=source_relative:. \
		$$f; \
		echo "generating $$f"; \
	done

.PHONY: proto/lint
proto/lint: ## protoファイルを検証します
	@echo "linting protos"
	buf check lint

.PHONY: proto/fmt
proto/fmt: proto/lint ## protoファイルのフォーマットを行います
	@echo "formating protos"
	prototool format -d . || true
	prototool format -w .

.PHONY: module
module: ## mod.go / mod.sum を整理します
	go get google.golang.org/grpc@v1.26
	go mod tidy

.PHONY: task/service
task/service: # taskサービスについて
	@echo "\033[31m"
	@echo "$$ micro list services"
	@echo "\033[0m"
	@echo "\033[33m"
	@echo "$$ grpcurl -plaintext -proto service/task/interface/proto/service.proto -d '{\"description\": \"hogehoge\"}' localhost:3000 task.Task/Create"
	@echo "$$ micro call task Task.Create '{\"description\": \"hogehoge\"}'"
	@echo "\033[0m"

.PHONY: task/run
task/run: task/service module ## taskサービスを起動します by go run
	go run service/task/cmd/srv/main.go service/task/cmd/srv/plugin.go --server_address=0.0.0.0:3000 --registry=etcd

.PHONY: task/bazel/run ## taskサービスを起動します by bazel
task/bazel/run: task/service bazel/gazelle bazel/update/repo
	bazelisk run service/task/cmd/srv:srv

.PHONY: bazel/clean
bazel/clean: ## bazelで生成された中間ファイルを削除します
	bazelisk clean --expunge

.PHONY: __
__:
	@echo "\033[33m"
	@echo "kzmake/micro-kit"
	@echo "\tMonorepo and Microservices kit using micro/micro for @kzmake"
	@echo "\033[0m"

.PHONY: help
help: __ ## ヘルプを表示します
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@cat $(MAKEFILE_LIST) \
	| grep -e "^[a-zA-Z_/\-]*: *.*## *" \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'
