load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "log.go",
        "logger.go",
    ],
    importpath = "github.com/kzmake/micro-kit/pkg/logger",
    visibility = ["//visibility:public"],
    deps = ["@com_github_rs_zerolog//:go_default_library"],
)
