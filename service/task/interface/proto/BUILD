load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "entity.pb.go",
        "entity.pb.micro.go",
        "entity.pb.validate.go",
        "service.pb.go",
        "service.pb.micro.go",
        "service.pb.validate.go",
    ],
    importpath = "github.com/kzmake/micro-kit/service/task/interface/proto",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_envoyproxy_protoc_gen_validate//validate:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
        "@com_github_golang_protobuf//ptypes/wrappers:go_default_library",
        "@com_github_micro_go_micro_v2//api:go_default_library",
        "@com_github_micro_go_micro_v2//client:go_default_library",
        "@com_github_micro_go_micro_v2//server:go_default_library",
        "@org_golang_google_protobuf//reflect/protoreflect:go_default_library",
        "@org_golang_google_protobuf//runtime/protoimpl:go_default_library",
    ],
)
