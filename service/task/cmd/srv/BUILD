load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "plugin.go",
    ],
    importpath = "github.com/kzmake/micro-kit/service/task/cmd/srv",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/logger:go_default_library",
        "//service/task/infrastructure/grpc:go_default_library",
        "//service/task/interface/controller:go_default_library",
        "//service/task/interface/presenter:go_default_library",
        "//service/task/usecase/interactor:go_default_library",
        "@com_github_micro_go_plugins_registry_consul_v2//:go_default_library",
        "@com_github_micro_go_plugins_registry_etcd_v2//:go_default_library",
        "@com_github_micro_go_plugins_registry_kubernetes_v2//:go_default_library",
        "@org_golang_x_xerrors//:go_default_library",
    ],
)

go_binary(
    name = "srv",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
