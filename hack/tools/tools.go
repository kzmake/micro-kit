// +build tools

package tools

// Package tools is used to track binary dependencies with go modules
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
import (
	// protoc tools
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-check-breaking"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-check-lint"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/micro/micro/v2/cmd/protoc-gen-micro"

	// lint tools
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"

	// fmt tools
	_ "golang.org/x/tools/cmd/goimports"

	// bazel tools
	_ "github.com/bazelbuild/bazel-gazelle/cmd/gazelle"
	_ "github.com/bazelbuild/bazelisk"
)
