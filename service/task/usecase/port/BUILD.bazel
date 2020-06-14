load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["create_task.go"],
    importpath = "github.com/kzmake/micro-kit/service/task/usecase/port",
    visibility = ["//visibility:public"],
    deps = ["//service/task/domain/aggregate:go_default_library"],
)
