## http_archive : https://docs.bazel.build/versions/master/repo/http.html#http_archive のロード
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

## bazelbuild/rules_go : https://github.com/bazelbuild/rules_go の設定
## https://github.com/bazelbuild/rules_go#initial-project-setup を参考
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "0c10738a488239750dbf35336be13252bad7c74348f867d30c3c3e0001906096",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.2/rules_go-v0.23.2.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.2/rules_go-v0.23.2.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

## bazelbuild/bazel-gazelle : https://github.com/bazelbuild/bazel-gazelle/tree/master#running-gazelle-with-bazel の設定
## https://github.com/bazelbuild/bazel-gazelle/tree/master#running-gazelle-with-bazel
## bazel/rules_go の version に注意
http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

## com_google_protobuf : https://github.com/protocolbuffers/protobuf の設定
## https://github.com/bazelbuild/rules_go#protobuf-and-grpc を参考
http_archive(
    name = "com_google_protobuf",
    sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
    strip_prefix = "protobuf-3.11.4",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

## go_repositoris の設定
# gazelle:repository_macro repositories.bzl%go_repositories

load("//:repositories.bzl", "go_repositories")

go_repositories()

go_repository(
    name = "cc_mvdan_interfacer",
    build_file_proto_mode = "disable_global",
    importpath = "mvdan.cc/interfacer",
    sum = "h1:WX1yoOaKQfddO/mLzdV4wptyWgoH/6hwLs7QHTixo0I=",
    version = "v0.0.0-20180901003855-c20040233aed",
)

go_repository(
    name = "cc_mvdan_lint",
    build_file_proto_mode = "disable_global",
    importpath = "mvdan.cc/lint",
    sum = "h1:DxJ5nJdkhDlLok9K6qO+5290kphDJbHOQO1DFFFTeBo=",
    version = "v0.0.0-20170908181259-adc824a0674b",
)

go_repository(
    name = "cc_mvdan_unparam",
    build_file_proto_mode = "disable_global",
    importpath = "mvdan.cc/unparam",
    sum = "h1:Cq7MalBHYACRd6EesksG1Q8EoIAKOsiZviGKbOLIej4=",
    version = "v0.0.0-20190720180237-d51796306d8f",
)

go_repository(
    name = "com_github_armon_consul_api",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/armon/consul-api",
    sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version = "v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name = "com_github_bazelbuild_bazelisk",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/bazelbuild/bazelisk",
    sum = "h1:N64d5HZoyGy6UxTTxDNp59dgF1vi10CvF9GWTJHmsZo=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_boltdb_bolt",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/boltdb/bolt",
    sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_bombsimon_wsl_v3",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/bombsimon/wsl/v3",
    sum = "h1:w9f49xQatuaeTJFaNP4SpiWSR5vfT6IstPtM62JjcqA=",
    version = "v3.0.0",
)

go_repository(
    name = "com_github_bufbuild_buf",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/bufbuild/buf",
    sum = "h1:jj+dZ6BiAK9dzvWPDUWVu8WglVyDR37Gys1/3WRlYAU=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_bufbuild_cli",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/bufbuild/cli",
    sum = "h1:ZcSiFwhd5bVIT3eRDdWfW50tLsi9wZpkAg4ucC/cw5M=",
    version = "v0.0.0-20200130190020-2009ccb4e7a8",
)

go_repository(
    name = "com_github_cespare_xxhash",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_chzyer_logex",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/coreos/go-etcd",
    sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/cpuguy83/go-md2man",
    sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_dgryski_go_sip13",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/dgryski/go-sip13",
    sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
    version = "v0.0.0-20181026042036-e10d5fee7954",
)

go_repository(
    name = "com_github_djarvur_go_err113",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/Djarvur/go-err113",
    sum = "h1:hY39LwQHh+1kaovmIjOrlqnXNX6tygSRfLkkK33IkZU=",
    version = "v0.0.0-20200410182137-af658d038157",
)

go_repository(
    name = "com_github_go_critic_go_critic",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-critic/go-critic",
    sum = "h1:4DTQfT1wWwLg/hzxwD9bkdhDQrdJtxe6DUTadPlrIeE=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_go_lintpack_lintpack",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-lintpack/lintpack",
    sum = "h1:DI5mA3+eKdWeJ40nU4d6Wc26qmdG8RCi/btYq0TuRN0=",
    version = "v0.5.2",
)

go_repository(
    name = "com_github_go_ole_go_ole",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-ole/go-ole",
    sum = "h1:2lOsA72HgjxAuMlKpFiCbHTvu44PIVkZ5hqm3RSdI/E=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:7LxgVwFb2hIQtMm87NdgAVfXjnt4OePseqT1tKx+opk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_go_toolsmith_astcast",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/astcast",
    sum = "h1:JojxlmI6STnFVG9yOImLeGREv8W2ocNUM+iOhR6jE7g=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astcopy",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/astcopy",
    sum = "h1:OMgl1b1MEpjFQ1m5ztEO06rz5CUd3oBv9RF7+DyvdG8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astequal",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/astequal",
    sum = "h1:4zxD8j3JRFNyLN46lodQuqz3xdKSrur7U/sr0SDS/gQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astfmt",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/astfmt",
    sum = "h1:A0vDDXt+vsvLEdbMFJAUBI/uTbRw1ffOPnxsILnFL6k=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astinfo",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/astinfo",
    sum = "h1:wP6mXeB2V/d1P1K7bZ5vDUO3YqEzcvOREOxZPEu3gVI=",
    version = "v0.0.0-20180906194353-9809ff7efb21",
)

go_repository(
    name = "com_github_go_toolsmith_astp",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/astp",
    sum = "h1:alXE75TXgcmupDsMK1fRAy0YUzLzqPVvBKoyWV+KPXg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_pkgload",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/pkgload",
    sum = "h1:4DFWWMXVfbcN5So1sBNW9+yeiMqLFGl1wFLTL5R0Tgg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_strparse",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/strparse",
    sum = "h1:Vcw78DnpCAKlM20kSbAyO4mPfJn/lyYA4BJUDxe2Jb4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_typep",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-toolsmith/typep",
    sum = "h1:zKymWyA1TRYvqYrYDrfEMZULyrhcnGY3x7LDKU2XQaA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_xmlfmt_xmlfmt",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-xmlfmt/xmlfmt",
    sum = "h1:khEcpUM4yFcxg4/FHQWkvVRmgijNXRfzkIDHh23ggEo=",
    version = "v0.0.0-20191208150333-d5b6f63a941b",
)

go_repository(
    name = "com_github_gobwas_glob",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gobwas/glob",
    sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_gofrs_flock",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gofrs/flock",
    sum = "h1:ekuhfTjngPhisSjOJ0QWKpPQE8/rbknHaes6WVJj5Hw=",
    version = "v0.0.0-20190320160742-5135e617513b",
)

go_repository(
    name = "com_github_golangci_check",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/check",
    sum = "h1:23T5iq8rbUYlhpt5DB4XJkc6BU31uODLD1o1gKvZmD0=",
    version = "v0.0.0-20180506172741-cfe4005ccda2",
)

go_repository(
    name = "com_github_golangci_dupl",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/dupl",
    sum = "h1:w8hkcTqaFpzKqonE9uMCefW1WDie15eSP/4MssdenaM=",
    version = "v0.0.0-20180902072040-3e9179ac440a",
)

go_repository(
    name = "com_github_golangci_errcheck",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/errcheck",
    sum = "h1:YYWNAGTKWhKpcLLt7aSj/odlKrSrelQwlovBpDuf19w=",
    version = "v0.0.0-20181223084120-ef45e06d44b6",
)

go_repository(
    name = "com_github_golangci_go_misc",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/go-misc",
    sum = "h1:9kfjN3AdxcbsZBf8NjltjWihK2QfBBBZuv91cMFfDHw=",
    version = "v0.0.0-20180628070357-927a3d87b613",
)

go_repository(
    name = "com_github_golangci_goconst",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/goconst",
    sum = "h1:pe9JHs3cHHDQgOFXJJdYkK6fLz2PWyYtP4hthoCMvs8=",
    version = "v0.0.0-20180610141641-041c5f2b40f3",
)

go_repository(
    name = "com_github_golangci_gocyclo",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/gocyclo",
    sum = "h1:J2XAy40+7yz70uaOiMbNnluTg7gyQhtGqLQncQh+4J8=",
    version = "v0.0.0-20180528134321-2becd97e67ee",
)

go_repository(
    name = "com_github_golangci_gofmt",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/gofmt",
    sum = "h1:iR3fYXUjHCR97qWS8ch1y9zPNsgXThGwjKPrYfqMPks=",
    version = "v0.0.0-20190930125516-244bba706f1a",
)

go_repository(
    name = "com_github_golangci_golangci_lint",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/golangci-lint",
    sum = "h1:VYLx63qb+XJsHdZ27PMS2w5JZacN0XG8ffUwe7yQomo=",
    version = "v1.27.0",
)

go_repository(
    name = "com_github_golangci_ineffassign",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/ineffassign",
    sum = "h1:gLLhTLMk2/SutryVJ6D4VZCU3CUqr8YloG7FPIBWFpI=",
    version = "v0.0.0-20190609212857-42439a7714cc",
)

go_repository(
    name = "com_github_golangci_lint_1",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/lint-1",
    sum = "h1:MfyDlzVjl1hoaPzPD4Gpb/QgoRfSBR0jdhwGyAWwMSA=",
    version = "v0.0.0-20191013205115-297bf364a8e0",
)

go_repository(
    name = "com_github_golangci_maligned",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/maligned",
    sum = "h1:kNY3/svz5T29MYHubXix4aDDuE3RWHkPvopM/EDv/MA=",
    version = "v0.0.0-20180506175553-b1d89398deca",
)

go_repository(
    name = "com_github_golangci_misspell",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/misspell",
    sum = "h1:EL/O5HGrF7Jaq0yNhBLucz9hTuRzj2LdwGBOaENgxIk=",
    version = "v0.0.0-20180809174111-950f5d19e770",
)

go_repository(
    name = "com_github_golangci_prealloc",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/prealloc",
    sum = "h1:leSNB7iYzLYSSx3J/s5sVf4Drkc68W2wm4Ixh/mr0us=",
    version = "v0.0.0-20180630174525-215b22d4de21",
)

go_repository(
    name = "com_github_golangci_revgrep",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/revgrep",
    sum = "h1:HVfrLniijszjS1aiNg8JbBMO2+E1WIQ+j/gL4SQqGPg=",
    version = "v0.0.0-20180526074752-d9c87f5ffaf0",
)

go_repository(
    name = "com_github_golangci_unconvert",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golangci/unconvert",
    sum = "h1:zwtduBRr5SSWhqsYNgcuWO2kFlpdOZbP0+yRjmvPGys=",
    version = "v0.0.0-20180507085042-28b1c447d1f4",
)

go_repository(
    name = "com_github_gookit_color",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gookit/color",
    sum = "h1:xOYBan3Fwlrqj1M1UN2TlHOCRiek3bGzWf/vPnJ1roE=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_gordonklaus_ineffassign",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gordonklaus/ineffassign",
    sum = "h1:vc7Dmrk4JwS0ZPS6WZvWlwDflgDTA26jItmbSj83nug=",
    version = "v0.0.0-20200309095847-7953dde2c7bf",
)

go_repository(
    name = "com_github_gostaticanalysis_analysisutil",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gostaticanalysis/analysisutil",
    sum = "h1:JVnpOZS+qxli+rgVl98ILOXVNbW+kb5wcxeGx8ShUIw=",
    version = "v0.0.0-20190318220348-4088753ea4d3",
)

go_repository(
    name = "com_github_hashicorp_go_version",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/hashicorp/go-version",
    sum = "h1:3vNe/fWF5CBgRIguda1meWhsZHy3m8gCJ5wx+dIzX/E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_iancoleman_strcase",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/iancoleman/strcase",
    sum = "h1:VHgatEHNcBFEB7inlalqfNqw65aNkM1lGX2yt3NmbS8=",
    version = "v0.0.0-20191112232945-16388991a334",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jhump_protoreflect",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/jhump/protoreflect",
    sum = "h1:2KMeMX92EK+pyOHQQcFpV77MlGtDQIdmKXqKnJqu3D8=",
    version = "v1.6.2-0.20200529043331-b0c184e143c4",
)

go_repository(
    name = "com_github_jingyugao_rowserrcheck",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/jingyugao/rowserrcheck",
    sum = "h1:GmsqmapfzSJkm28dhRoHz2tLRbJmqhU86IPgBtN3mmk=",
    version = "v0.0.0-20191204022205-72ab7603b68a",
)

go_repository(
    name = "com_github_jirfag_go_printf_func_name",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/jirfag/go-printf-func-name",
    sum = "h1:jNYPNLe3d8smommaoQlK7LOA5ESyUJJ+Wf79ZtA7Vp4=",
    version = "v0.0.0-20191110105641-45db9963cdd3",
)

go_repository(
    name = "com_github_jmoiron_sqlx",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/jmoiron/sqlx",
    sum = "h1:lrdPtrORjGv1HbbEvKWDUAy97mPpFm4B8hp77tcCUJY=",
    version = "v1.2.1-0.20190826204134-d7d95172beb5",
)

go_repository(
    name = "com_github_klauspost_compress",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/klauspost/compress",
    sum = "h1:8VMb5+0wMgdBykOV96DwNwKFQ+WTI4pzYURP99CcB9E=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_logrusorgru_aurora",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/logrusorgru/aurora",
    sum = "h1:9MlwzLdW7QSDrhDjFlsEYmxpFyIoXmYRon3dt0io31k=",
    version = "v0.0.0-20181002194514-a7b3b318ed4e",
)

go_repository(
    name = "com_github_lyft_protoc_gen_star",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/lyft/protoc-gen-star",
    sum = "h1:quC0MYv1hc+s8Wz9qCdPPI+DjhEPVD9gHZn2So2jbwc=",
    version = "v0.4.15",
)

go_repository(
    name = "com_github_magiconair_properties",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/magiconair/properties",
    sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_maratori_testpackage",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/maratori/testpackage",
    sum = "h1:QtJ5ZjqapShm0w5DosRjg0PRlSdAdlx+W6cCKoALdbQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_matoous_godox",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/matoous/godox",
    sum = "h1:RHba4YImhrUVQDHUCe2BNSOz4tVy2yGyXhvYDvxGgeE=",
    version = "v0.0.0-20190911065817-5d6d842e92eb",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:pDRiWfl+++eC2FEFRy6jXmQlvp4Yh3z1MJKg4UeYM/4=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_mattn_goveralls",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mattn/goveralls",
    sum = "h1:7eJB6EqsPhRVxvwEXGnqdO2sJI0PTsrWoTMXEk9/OQc=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_micro_micro_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/micro/micro/v2",
    sum = "h1:8gF0SIHMkdsXWt4ItmnUKvEuYF4U8BSkcAXp+Li+4O8=",
    version = "v2.7.0",
)

go_repository(
    name = "com_github_mitchellh_go_ps",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mitchellh/go-ps",
    sum = "h1:9+ke9YJ9KGWw5ANXK6ozjoK47uI3uNbXv4YVINBnGm8=",
    version = "v0.0.0-20190716172923-621e5597135b",
)

go_repository(
    name = "com_github_mozilla_tls_observatory",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mozilla/tls-observatory",
    sum = "h1:1xJ+Xi9lYWLaaP4yB67ah0+548CD3110mCPWhVVjFkI=",
    version = "v0.0.0-20200317151703-4fa42e1c2dee",
)

go_repository(
    name = "com_github_nakabonne_nestif",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/nakabonne/nestif",
    sum = "h1:+yOViDGhg8ygGrmII72nV9B/zGxY188TYpfolntsaPw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_nbutton23_zxcvbn_go",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/nbutton23/zxcvbn-go",
    sum = "h1:AREM5mwr4u1ORQBMvzfzBgpsctsbQikCVpvC+tX285E=",
    version = "v0.0.0-20180912185939-ae427f1e4c1d",
)

go_repository(
    name = "com_github_netdata_go_orchestrator",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/netdata/go-orchestrator",
    sum = "h1:jSzujNrzCHAxa05SDRsjNx/OBbElK3a6yUtGqJXr32w=",
    version = "v0.0.0-20190905093727-c793edba0e8f",
)

go_repository(
    name = "com_github_nishanths_predeclared",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/nishanths/predeclared",
    sum = "h1:3f0nxAmdj/VoCGN/ijdMy7bj6SBagaqYg1B0hu8clMA=",
    version = "v0.0.0-20200524104333-86fad755b4d3",
)

go_repository(
    name = "com_github_oklog_ulid",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/oklog/ulid",
    sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_openpeedeep_depguard",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/OpenPeeDeeP/depguard",
    sum = "h1:VlW4R6jmBIv3/u1JNlawEvJMM4J+dPORPaZasQee8Us=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_pborman_uuid",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/pborman/uuid",
    sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/pelletier/go-toml",
    sum = "h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_phayes_checkstyle",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/phayes/checkstyle",
    sum = "h1:CdDQnGF8Nq9ocOS/xlSptM1N3BbrA6/kmaep5ggwaIA=",
    version = "v0.0.0-20170904204023-bfd46e6a821d",
)

go_repository(
    name = "com_github_pkg_profile",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/pkg/profile",
    sum = "h1:042Buzk+NhDI+DeSAA62RwJL8VAuZUMQZUjCsRz1Mug=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_prometheus_tsdb",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/prometheus/tsdb",
    sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
    version = "v0.7.1",
)

go_repository(
    name = "com_github_quasilyte_go_consistent",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/quasilyte/go-consistent",
    sum = "h1:JoUA0uz9U0FVFq5p4LjEq4C0VgQ0El320s3Ms0V4eww=",
    version = "v0.0.0-20190521200055-c6f3937de18c",
)

go_repository(
    name = "com_github_russross_blackfriday",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_ryancurrah_gomodguard",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/ryancurrah/gomodguard",
    sum = "h1:oCreMAt9GuFXDe9jW4HBpc3GjdX3R/sUEcLAGh1zPx8=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_securego_gosec_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/securego/gosec/v2",
    sum = "h1:y/9mCF2WPDbSDpL3QDWZD3HHGrSYw0QSHnCqTfs4JPE=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_serenize_snaker",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/serenize/snaker",
    sum = "h1:ofR1ZdrNSkiWcMsRrubK9tb2/SlZVWttAfqUjJi6QYc=",
    version = "v0.0.0-20171204205717-a683aaf2d516",
)

go_repository(
    name = "com_github_shirou_gopsutil",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/shirou/gopsutil",
    sum = "h1:WokF3GuxBeL+n4Lk4Fa8v9mbdjlrl7bHuneF4N1bk2I=",
    version = "v0.0.0-20190901111213-e4ec7b275ada",
)

go_repository(
    name = "com_github_shirou_w32",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/shirou/w32",
    sum = "h1:udFKJ0aHUL60LboW/A+DfgoHVedieIzIXE8uylPue0U=",
    version = "v0.0.0-20160930032740-bb4de0191aa4",
)

go_repository(
    name = "com_github_shurcool_go",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/shurcooL/go",
    sum = "h1:MZM7FHLqUHYI0Y/mQAt3d2aYa0SiNms/hFqC9qJYolM=",
    version = "v0.0.0-20180423040247-9e1955d9fb6e",
)

go_repository(
    name = "com_github_shurcool_go_goon",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/shurcooL/go-goon",
    sum = "h1:llrF3Fs4018ePo4+G/HV/uQUqEI1HMDjCeOf2V6puPc=",
    version = "v0.0.0-20170922171312-37c2f522c041",
)

go_repository(
    name = "com_github_sourcegraph_go_diff",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/sourcegraph/go-diff",
    sum = "h1:gO6i5zugwzo1RVTvgvfwCOSVegNuvnNi6bAD1QCmkHs=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
    version = "v0.0.0-20180118202830-f09979ecbc72",
)

go_repository(
    name = "com_github_spf13_afero",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/spf13/afero",
    sum = "h1:m8/z1t7/fwjysjQRYbP0RD+bUIF/8tJwPdEZsI83ACI=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_spf13_cast",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/spf13/cast",
    sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/spf13/cobra",
    sum = "h1:6m/oheQuQ13N9ks4hubMG6BnvwOeaJrqSPLahSnczz8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/spf13/jwalterweatherman",
    sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/spf13/pflag",
    sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_spf13_viper",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/spf13/viper",
    sum = "h1:pDDu1OyEDTKzpJwdq4TiuLyMsUgRa/BT5cn5O62NoHs=",
    version = "v1.6.3",
)

go_repository(
    name = "com_github_stackexchange_wmi",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/StackExchange/wmi",
    sum = "h1:fLjPD/aNc3UIOA6tDi6QXUemppXK3P9BI7mr2hd6gx8=",
    version = "v0.0.0-20180116203802-5d049714c4a6",
)

go_repository(
    name = "com_github_subosito_gotenv",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/subosito/gotenv",
    sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_tdakkota_asciicheck",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/tdakkota/asciicheck",
    sum = "h1:Xr9gkxfOP0KQWXKNqmwe8vEeSUiUj4Rlee9CMVX2ZUQ=",
    version = "v0.0.0-20200416190851-d7f85be797a2",
)

go_repository(
    name = "com_github_tetafro_godot",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/tetafro/godot",
    sum = "h1:+mecr7RKrUKB5UQ1gwqEMn13sDKTyDR8KNIquB9mm+8=",
    version = "v0.3.7",
)

go_repository(
    name = "com_github_timakin_bodyclose",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/timakin/bodyclose",
    sum = "h1:RumXZ56IrCj4CL+g1b9OL/oH0QnsF976bC8xQFYUD5Q=",
    version = "v0.0.0-20190930140734-f7f2e9bca95e",
)

go_repository(
    name = "com_github_tommy_muehle_go_mnd",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/tommy-muehle/go-mnd",
    sum = "h1:RC4maTWLKKwb7p1cnoygsbKIgNlJqSYBeAFON3Ar8As=",
    version = "v1.3.1-0.20200224220436-e6f9a994e8fa",
)

go_repository(
    name = "com_github_ugorji_go",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/ugorji/go",
    sum = "h1:j4s+tAvLfL3bZyefP2SEWmhBzmuIlH/eqNuPdFPgngw=",
    version = "v1.1.4",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:3SVOIvH7Ae1KRYyQWRjXWJEA9sS/c/pjvH++55Gr648=",
    version = "v0.0.0-20181204163529-d75b2dcb6bc8",
)

go_repository(
    name = "com_github_ultraware_funlen",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/ultraware/funlen",
    sum = "h1:Av96YVBwwNSe4MLR7iI/BIa3VyI7/djnto/pK3Uxbdo=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ultraware_whitespace",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/ultraware/whitespace",
    sum = "h1:If7Va4cM03mpgrNH9k49/VOicWpGoG70XPBFFODYDsg=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_uudashr_gocognit",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/uudashr/gocognit",
    sum = "h1:MoG2fZ0b/Eo7NXoIwCVFLG5JED3qgQz5/NEE+rOsjPs=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasthttp",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/valyala/fasthttp",
    sum = "h1:dzZJf2IuMiclVjdw0kkT+f9u4YdrapbNyGAN47E/qnk=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_valyala_quicktemplate",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/valyala/quicktemplate",
    sum = "h1:BaO1nHTkspYzmAjPXj0QiDJxai96tlcZyKcI9dyEGvM=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_valyala_tcplisten",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/valyala/tcplisten",
    sum = "h1:0R4NLDRDZX6JcmhJgXi5E4b8Wg84ihbmUKp/GvSPEzc=",
    version = "v0.0.0-20161114210144-ceec8f93295a",
)

go_repository(
    name = "com_github_xlab_treeprint",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/xlab/treeprint",
    sum = "h1:1CFlNzQhALwjS9mBAUkycX616GzgsuYUOCHA5+HSlXI=",
    version = "v0.0.0-20181112141820-a009c3971eca",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/xordataexchange/crypt",
    sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "com_sourcegraph_sqs_pbtypes",
    build_file_proto_mode = "disable_global",
    importpath = "sourcegraph.com/sqs/pbtypes",
    sum = "h1:JPJh2pk3+X4lXAkZIk2RuE/7/FoK9maXw+TNPJhVS/c=",
    version = "v0.0.0-20180604144634-d3ebe8f20ae4",
)
