module github.com/kzmake/micro-kit

go 1.14

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3

require (
	github.com/certifi/gocertifi v0.0.0-20200211180108-c7c1fbc02894 // indirect
	github.com/cockroachdb/errors v1.2.4
	github.com/cockroachdb/logtags v0.0.0-20190617123548-eb05cc24525f // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.12
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/registry/consul/v2 v2.8.0
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.8.0
	github.com/oklog/ulid/v2 v2.0.2
	github.com/opentracing/opentracing-go v1.1.0
	github.com/rs/zerolog v1.19.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/stretchr/testify v1.6.1
	github.com/uber/jaeger-client-go v2.24.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/protobuf v1.23.0
)
