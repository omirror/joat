workspace(name = "joat")

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

load("@io_bazel_rules_go//go:def.bzl", "go_download_sdk", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    go_image_repositories = "repositories",
)

go_image_repositories()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_armon_go_metrics",
    commit = "783273d703149aaeb9897cf58613d5af48861c25",
    importpath = "github.com/armon/go-metrics",
)

go_repository(
    name = "com_github_docker_distribution",
    commit = "48294d928ced5dd9b378f7fd7c6f5da3ff3f2c89",
    importpath = "github.com/docker/distribution",
)

go_repository(
    name = "com_github_docker_docker",
    commit = "092cba3727bb9b4a2f0e922cd6c0f93ea270e363",
    importpath = "github.com/docker/docker",
)

go_repository(
    name = "com_github_docker_go_connections",
    commit = "3ede32e2033de7505e6500d6c868c2b9ed9f169d",
    importpath = "github.com/docker/go-connections",
)

go_repository(
    name = "com_github_docker_go_units",
    commit = "47565b4f722fb6ceae66b95f853feed578a4a51c",
    importpath = "github.com/docker/go-units",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    commit = "c2828203cd70a50dcccfb2761f8b1f8ceef9a8e9",
    importpath = "github.com/fsnotify/fsnotify",
)

go_repository(
    name = "com_github_go_chi_chi",
    commit = "e83ac2304db3c50cf03d96a2fcd39009d458bc35",
    importpath = "github.com/go-chi/chi",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "b4deda0973fb4c70b50d226b1af49f3da59f5265",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    commit = "7554cd9344cec97297fa6649b055a8c98c2a1e55",
    importpath = "github.com/hashicorp/errwrap",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    commit = "7f3cd4390caab3250a57f30efdb2a65dd7649ecf",
    importpath = "github.com/hashicorp/go-immutable-radix",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    commit = "fa3f63826f7c23912c15263591e65d54d080b458",
    importpath = "github.com/hashicorp/go-msgpack",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    commit = "b7773ae218740a7be65057fc60b366a49b538a44",
    importpath = "github.com/hashicorp/go-multierror",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    commit = "6d291a969b86c4b633730bfc6b8b9d64c3aafed9",
    importpath = "github.com/hashicorp/go-sockaddr",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    commit = "0fb14efe8c47ae851c0034ed7a448854d3d34cf3",
    importpath = "github.com/hashicorp/golang-lru",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    commit = "ef8a98b0bbce4a65b5aa4c368430a80ddc533168",
    importpath = "github.com/hashicorp/hcl",
)

go_repository(
    name = "com_github_hashicorp_memberlist",
    commit = "ce8abaa0c60c2d6bee7219f5ddf500e0a1457b28",
    importpath = "github.com/hashicorp/memberlist",
)

go_repository(
    name = "com_github_hashicorp_serf",
    commit = "0df3e3df1703f838243a7f3f12bf0b88566ade5a",
    importpath = "github.com/hashicorp/serf",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    commit = "76626ae9c91c4f2a10f34cad8ce83ea42c93bb75",
    importpath = "github.com/inconshreveable/mousetrap",
)

go_repository(
    name = "com_github_magiconair_properties",
    commit = "c3beff4c2358b44d0493c7dda585e7db7ff28ae6",
    importpath = "github.com/magiconair/properties",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    commit = "7da180ee92d8bd8bb8c37fc560e673e6557c392f",
    importpath = "github.com/Microsoft/go-winio",
)

go_repository(
    name = "com_github_miekg_dns",
    commit = "eac804ceef194db2da6ee80c728d7658c8c805ff",
    importpath = "github.com/miekg/dns",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    commit = "b8bc1bf767474819792c23f32d8286a45736f1c6",
    importpath = "github.com/mitchellh/go-homedir",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "bb74f1db0675b241733089d5a1faa5dd8b0ef57b",
    importpath = "github.com/mitchellh/mapstructure",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    commit = "acdc4509485b587f5e675510c4f2c63e90ff68a8",
    importpath = "github.com/pelletier/go-toml",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "645ef00459ed84a119197bfb8d8205042c6df63d",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_rs_zerolog",
    commit = "05eafee0eb17d0150591a8f30f0fa592cc9b7471",
    importpath = "github.com/rs/zerolog",
)

go_repository(
    name = "com_github_satori_go_uuid",
    commit = "f58768cc1a7a7e77a3bd49e98cdd21419399b6a3",
    importpath = "github.com/satori/go.uuid",
)

go_repository(
    name = "com_github_sean__seed",
    commit = "e2103e2c35297fb7e17febb81e49b312087a2372",
    importpath = "github.com/sean-/seed",
)

go_repository(
    name = "com_github_spf13_afero",
    commit = "63644898a8da0bc22138abf860edaf5277b6102e",
    importpath = "github.com/spf13/afero",
)

go_repository(
    name = "com_github_spf13_cast",
    commit = "8965335b8c7107321228e3e3702cab9832751bac",
    importpath = "github.com/spf13/cast",
)

go_repository(
    name = "com_github_spf13_cobra",
    commit = "ef82de70bb3f60c65fb8eebacbb2d122ef517385",
    importpath = "github.com/spf13/cobra",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    commit = "7c0cea34c8ece3fbeb2b27ab9b59511d360fb394",
    importpath = "github.com/spf13/jwalterweatherman",
)

go_repository(
    name = "com_github_spf13_pflag",
    commit = "583c0c0531f06d5278b7d917446061adc344b5cd",
    importpath = "github.com/spf13/pflag",
)

go_repository(
    name = "com_github_spf13_viper",
    commit = "b5e8006cbee93ec955a89ab31e0e3ce3204f3736",
    importpath = "github.com/spf13/viper",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "7bb2a897381c9c5ab2aeb8614f758d7766af68ff",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "41344da2231b913fa3d983840a57a6b1b7b631a1",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "1a580b3eff7814fc9b40602fd35256c63b50f491",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "2491c5de3490fced2f6cff376127c667efeed857",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "7c87d13f8e835d2fb3a70a2912c811ed0c1d241b",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)
