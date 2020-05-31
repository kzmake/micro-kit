
SERVICE_DIR := service

SOURCES := $(shell find . -name "*.go" -type f)

__symbol := 🧊

$(warning SERVICE_DIR = $(SERVICE_DIR))

.DEFAULT_GOAL := help

.PHONY: tools
tools: ## 開発に必要なツールをインストールします (./bin/)
	@echo "\033[31m"
	@echo "$$ brew install protobuf"
	@echo "\033[0m"
	go generate -tags=tools ./...
	go install github.com/golang/protobuf/protoc-gen-go
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	ls -l ./bin

.PHONY: fmt
fmt: ## goファイルをフォーマットします
	bin/goimports -l -w .

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
		echo "${__symbol} generating $$f"; \
	done

.PHONY: proto/lint
proto/lint: ## protoファイルを検証します
	@echo "${__symbol} linting protos"
	./bin/buf check lint

.PHONY: proto/fmt
proto/fmt: proto/lint ## protoファイルのフォーマットを行います
	@echo "${__symbol} formating protos"
	@prototool format -d .
	@prototool format -w .

.PHONY: module
module: ## mod.go / mod.sum を整理します
	go get google.golang.org/grpc@v1.26
	go mod edit -require=google.golang.org/grpc@v1.26.0
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
	./bin/bazelisk run service/task/cmd/srv:srv

.PHONY: bazel/gazelle
bazel/gazelle: ## gazelleをセットアップします
	./bin/bazelisk run :gazelle

.PHONY: bazel/update/repo
bazel/update/repo: ## go.modからbazelのgo-repositoriesを更新します
	./bin/bazelisk run :gazelle -- update-repos -from_file=go.mod -build_file_proto_mode=disable_global

.PHONY: bazel/update/repo/prune
bazel/update/repo/prune: ## go.modからbazelのgo-repositoriesを更新します(不要な依存パッケージは削除されます)
	go mod tidy
	./bin/bazelisk run :gazelle -- update-repos -from_file=go.mod -build_file_proto_mode=disable_global -prune=true

.PHONY: bazel/clean
bazel/clean: ## bazelで生成された中間ファイルやBUILDなどをすべて削除します
	./bin/bazelisk clean --expunge
	find . -path "./BUILD*" -prune -o -type f -name 'BUILD*' -delete

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
