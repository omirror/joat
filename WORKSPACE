workspace(name = "joat")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "build_bazel_rules_nodejs",
    remote = "https://github.com/bazelbuild/rules_nodejs.git",
    tag = "0.9.1",
)

load("@build_bazel_rules_nodejs//:defs.bzl", "node_repositories", "yarn_install")

yarn_install(
    name = "build_bazel_rules_nodejs_managed_deps",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

node_repositories(
    package_json = ["//:package.json"],
    preserve_symlinks = True,
)

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.12.0/rules_go-0.12.0.tar.gz",
    sha256 = "c1f52b8789218bb1542ed362c4f7de7052abcf254d865d96fb7ba6d44bc15ee3",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz",
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
)

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.4.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    go_image_repositories = "repositories",
)

go_image_repositories()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

git_repository(
    name = "io_bazel_rules_webtesting",
    remote = "https://github.com/bazelbuild/rules_webtesting.git",
    commit = "ca7b8062d9cf4ef2fde9193c7d37a0764c4262d7",
)

load("@io_bazel_rules_webtesting//web:repositories.bzl", "browser_repositories", "web_test_repositories")

web_test_repositories()

browser_repositories(
    chromium = True,
    firefox = True,
)

git_repository(
    name = "build_bazel_rules_typescript",
    remote = "https://github.com/bazelbuild/rules_typescript.git",
    tag = "0.14.0",
)

load("@build_bazel_rules_typescript//:defs.bzl", "ts_setup_workspace")

ts_setup_workspace()
