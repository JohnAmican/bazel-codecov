"""This module implements the language-specific toolchain rule."""

CodecovInfo = provider(
    doc = "Information about how to invoke the codecov executable.",
    fields = {
        "binary": "Executable codecov binary",
        "version": "Codecov version",
    },
)

def _codecov_toolchain_impl(ctx):
    binary = ctx.executable.codecov
    template_variables = platform_common.TemplateVariableInfo({
        "CODECOV_BIN": binary.path,
    })
    default = DefaultInfo(
        files = depset([binary]),
        runfiles = ctx.runfiles(files = [binary]),
    )
    codecov_info = CodecovInfo(
        binary = binary,
        version = ctx.attr.version.removeprefix("v"),
    )
    toolchain_info = platform_common.ToolchainInfo(
        codecov_info = codecov_info,
        template_variables = template_variables,
        default = default,
    )
    return [
        default,
        toolchain_info,
        template_variables,
    ]

codecov_toolchain = rule(
    implementation = _codecov_toolchain_impl,
    attrs = {
        "codecov": attr.label(
            doc = "A hermetically downloaded executable target for the target platform.",
            mandatory = True,
            allow_single_file = True,
            executable = True,
            cfg = "exec",
        ),
        "version": attr.string(mandatory = True, doc = "Version of the codecov binary"),
    },
    doc = "Defines a codecov toolchain. See: https://docs.bazel.build/versions/main/toolchains.html#defining-toolchains.",
)
