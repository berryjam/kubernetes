package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fake_certificates_client.go",
        "fake_certificatesigningrequest.go",
        "fake_certificatesigningrequest_expansion.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/apis/certificates:go_default_library",
        "//pkg/apis/meta/v1:go_default_library",
        "//pkg/client/clientset_generated/internalclientset/typed/certificates/internalversion:go_default_library",
        "//pkg/client/restclient:go_default_library",
        "//pkg/client/testing/core:go_default_library",
        "//pkg/labels:go_default_library",
        "//pkg/runtime/schema:go_default_library",
        "//pkg/watch:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
