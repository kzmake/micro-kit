load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/kzmake/micro-kit/api/task/cmd/api",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/logger/technical:go_default_library",
        "//service/task/infrastructure/grpc:go_default_library",
        "//service/task/interface/controller:go_default_library",
        "//service/task/usecase/interactor:go_default_library",
        "//vendor/golang.org/x/xerrors:go_default_library",
    ],
)

go_binary(
    name = "api",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
