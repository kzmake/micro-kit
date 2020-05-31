load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "log.go",
    ],
    importpath = "github.com/kzmake/micro-kit/pkg/wrapper/log",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/logger:go_default_library",
        "@com_github_micro_go_micro_v2//client:go_default_library",
        "@com_github_micro_go_micro_v2//metadata:go_default_library",
        "@com_github_micro_go_micro_v2//server:go_default_library",
    ],
)
