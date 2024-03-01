load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# TODO: rename
def codecov_cli(name, cli, coverage_files, **kwargs):
    go_library(
        name = "codecov_lib",
        srcs = [
            "debug.go",
            "main.go",
        ],
        importpath = "github.com/JohnAmican/bazel-codecov/codecov",
        visibility = ["//visibility:private"],
        deps = [
            "@com_github_davecgh_go_spew//spew",
            "@io_bazel_rules_go//go/runfiles:go_default_library",
        ],
    )

    go_binary(
        name = "codecov",
        data = coverage_files + [cli],
        embed = [":codecov_lib"],
        visibility = ["//visibility:public"],
        x_defs = {
            "cli": cli,
            #            "cliPkg": Label(entrypoint).repo_name,
            #            "cliName": Label(entrypoint).package + ":" + Label(entrypoint).name,
        },
    )
