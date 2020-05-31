// +build tools

package tools

import (
	// protoc
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/micro/micro/v2/cmd/protoc-gen-micro"

	// buf
	_ "github.com/bufbuild/buf/cmd/buf"

	// goimports
	_ "golang.org/x/tools/cmd/goimports"

	// godoc
	_ "golang.org/x/tools/cmd/godoc"

	// golangci-lint
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"

	// bazel
	_ "github.com/bazelbuild/bazelisk"
)

//go:generate go build -v -o bin/protoc-gen-validate github.com/envoyproxy/protoc-gen-validate
//go:generate go build -v -o bin/protoc-gen-micro github.com/micro/micro/v2/cmd/protoc-gen-micro
//go:generate go build -v -o bin/goimports golang.org/x/tools/cmd/goimports
//go:generate go build -v -o bin/buf github.com/bufbuild/buf/cmd/buf
//go:generate go build -v -o bin/godoc golang.org/x/tools/cmd/godoc
//go:generate go build -v -o bin/bazelisk github.com/bazelbuild/bazelisk
