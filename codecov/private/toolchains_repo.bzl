"""Create a repository to hold the toolchains

This follows guidance here:
https://docs.bazel.build/versions/main/skylark/deploying.html#registering-toolchains
"
Note that in order to resolve toolchains in the analysis phase
Bazel needs to analyze all toolchain targets that are registered.
Bazel will not need to analyze all targets referenced by toolchain.toolchain attribute.
If in order to register toolchains you need to perform complex computation in the repository,
consider splitting the repository with toolchain targets
from the repository with <LANG>_toolchain targets.
Former will be always fetched,
and the latter will only be fetched when user actually needs to build <LANG> code.
"
The "complex computation" in our case is simply downloading large artifacts.
This guidance tells us how to avoid that: we put the toolchain targets in the alias repository
with only the toolchain attribute pointing into the platform-specific repositories.
"""

# Add more platforms as needed to mirror all the binaries
# published by the upstream project.
PLATFORMS = {
    "darwin_amd64": struct(
        compatible_with = [
            "@platforms//os:macos",
            "@platforms//cpu:x86_64",
        ],
    ),
    #    "darwin_arm64": struct(
    #        compatible_with = [
    #            "@platforms//os:macos",
    #            "@platforms//cpu:aarch64",
    #        ],
    #    ),
    #    "linux_arm64": struct(
    #        compatible_with = [
    #            "@platforms//os:linux",
    #            "@platforms//cpu:aarch64",
    #        ],
    #    ),
    #    "linux_armv6": struct(
    #        compatible_with = [
    #            "@platforms//os:linux",
    #            "@platforms//cpu:arm64",
    #        ],
    #    ),
    #    "linux_i386": struct(
    #        compatible_with = [
    #            "@platforms//os:linux",
    #            "@platforms//cpu:x86_32",
    #        ],
    #    ),
    #    "linux_s390x": struct(
    #        compatible_with = [
    #            "@platforms//os:linux",
    #            "@platforms//cpu:s390x",
    #        ],
    #    ),
    #    "linux_amd64": struct(
    #        compatible_with = [
    #            "@platforms//os:linux",
    #            "@platforms//cpu:x86_64",
    #        ],
    #    ),
    #    "windows_armv6": struct(
    #        compatible_with = [
    #            "@platforms//os:windows",
    #            "@platforms//cpu:arm64",
    #        ],
    #    ),
    #    "windows_amd64": struct(
    #        compatible_with = [
    #            "@platforms//os:windows",
    #            "@platforms//cpu:x86_64",
    #        ],
    #    ),
}

DEFS_TMPL = """\
# Generated by toolchains_repo.bzl for {toolchain_type}
load("@bazel_skylib//lib:structs.bzl", "structs")

# Forward all the providers
def _resolved_toolchain_impl(ctx):
    toolchain_info = ctx.toolchains["{toolchain_type}"]
    return [toolchain_info] + structs.to_dict(toolchain_info).values()

# Copied from java_toolchain_alias
# https://cs.opensource.google/bazel/bazel/+/master:tools/jdk/java_toolchain_alias.bzl
resolved_toolchain = rule(
    implementation = _resolved_toolchain_impl,
    toolchains = ["{toolchain_type}"],
)
"""

TOOLCHAIN_TMPL = """\
toolchain(
    name = "{platform}_toolchain",
    exec_compatible_with = {compatible_with},
    toolchain = "{toolchain}",
    toolchain_type = "{toolchain_type}",
)
"""

BUILD_HEADER_TMPL = """\
# Generated by toolchains_repo.bzl
#
# These can be registered in the workspace file or passed to --extra_toolchains flag.
# By default all of these toolchains are registered by the oci_register_toolchains macro
# so you don't normally need to interact with these targets.

load(":defs.bzl", "resolved_toolchain")

resolved_toolchain(name = "current_toolchain", visibility = ["//visibility:public"])
"""

def _toolchains_repo_impl(repository_ctx):
    # Expose a concrete toolchain which is the result of Bazel resolving the toolchain
    # for the execution or target platform.
    # Workaround for https://github.com/bazelbuild/bazel/issues/14009
    defs_content = DEFS_TMPL.format(
        toolchain_type = repository_ctx.attr.toolchain_type,
    )
    repository_ctx.file("defs.bzl", defs_content)

    build_content = BUILD_HEADER_TMPL

    for [platform, meta] in PLATFORMS.items():
        build_content += TOOLCHAIN_TMPL.format(
            platform = platform,
            name = repository_ctx.attr.name,
            compatible_with = meta.compatible_with,
            toolchain_type = repository_ctx.attr.toolchain_type,
            toolchain = repository_ctx.attr.toolchain.format(platform = platform),
        )

    repository_ctx.file("BUILD.bazel", build_content)

toolchains_repo = repository_rule(
    _toolchains_repo_impl,
    doc = "Creates a repository with toolchain definitions for all known platforms which can be registered or selected.",
    attrs = {
        "toolchain": attr.string(doc = "Label of the toolchain with {platform} left as placeholder. example; @container_crane_{platform}//:crane_toolchain"),
        "toolchain_type": attr.string(doc = "Label of the toolchain_type. example; //oci:crane_toolchain_type"),
    },
)