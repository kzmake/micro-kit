
SERVICE_DIR := service

SOURCES := $(shell find . -name "*.go" -type f)

__symbol := 🧊

$(warning SERVICE_DIR = $(SERVICE_DIR))

.DEFAULT_GOAL := help
.PHONY: proto proto/lint proto/fmt help

fmt:
	@hash goimports > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		env GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports; \
	fi
	@goimports -l -w $(SOURCES)

proto:
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

proto/lint:
	@echo "${__symbol} linting protos"
	@buf check lint

# I prefer VS Code's proto plugin to format my code then prototool
proto/fmt: proto/lint
	@echo "${__symbol} formating protos"
	@prototool format -d .
	@prototool format -w .

__:
	@echo "\033[33m"
	@echo "kzmake/micro-kit"
	@echo "\tfor @kzmake using micro/micro"
	@echo "\033[0m"

help: __ ## Show help
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@cat $(MAKEFILE_LIST) \
	| grep -e "^[a-zA-Z_/\-]*: *.*## *" \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}' \
	| sed "s/\(.*\/.*\)/  \1/"
