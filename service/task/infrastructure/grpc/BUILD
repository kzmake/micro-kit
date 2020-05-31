load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    importpath = "github.com/kzmake/micro-kit/service/task/infrastructure/grpc",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/constant:go_default_library",
        "//pkg/logger:go_default_library",
        "//pkg/wrapper/log:go_default_library",
        "//service/task/interface/proto:go_default_library",
        "@com_github_micro_cli_v2//:go_default_library",
        "@com_github_micro_go_micro_v2//:go_default_library",
        "@com_github_micro_go_micro_v2//server:go_default_library",
        "@com_github_micro_go_plugins_registry_etcdv3_v2//:go_default_library",
        "@org_golang_x_xerrors//:go_default_library",
    ],
)
