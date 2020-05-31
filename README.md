# Monorepo and Microservices kit using micro/micro for [@kzmake](https://github.com/kzmake)

※あくまで自分用です

- クリーンアーキテクチャ(を意識したレイヤードアーキテクチャ)
- マイクロサービス
- モノレポ
- 軽量DDD

---

- Inspired by [@xmlking's micro-starter-kit](https://github.com/xmlking/micro-starter-kit)
- Inspired by [@kgrzybek's modular-monolith-with-ddd](https://github.com/kgrzybek/modular-monolith-with-ddd)

## Roadmap

- [ ] Monorepo
- [ ] Microservices with micro/go-micro
- [ ] gRPC microservices
- [ ] Support HTTP-API
- [ ] Support CLI
- [ ] Support NATS
- [ ] Kubernetes with Kustomize
- [ ] Input Validation with PGV
- [x] Linting Protos with Buf
- [ ] Observability
- [x] Building/Testing with Bazel
- [ ] Support gRPC reflection API 

## Getting Started
### Init

```sh
git clone https://github.com/kzmake/micro-kit.git $GOPATH/src/github.com/kzmake/micro-kit
cd $GOPATH/src/github.com/kzmake/micro-kit
go mod download
```

### Service Discovery

```sh
# install
go get -v go.etcd.io/etcd
# or
brew install etcd

# run
etcd
# or
brew servces start etcd
```

### Client
```sh
# install
go get -u github.com/micro/micro/v2

# run
micro --version
```

### Run

```sh
make task/run
```

### Request

```sh
# use micro
micro call task Task.Create '{"description": "hoge"}'

# use grpcurl
grpcurl -plaintext -proto service/task/interface/proto/service.proto -d '{"description": "hogehoge"}' localhost:3000 task.Task/Create
```

## Credits

- Inspired by [@xmlking's micro-starter-kit](https://github.com/xmlking/micro-starter-kit)
- Inspired by [@kgrzybek's modular-monolith-with-ddd](https://github.com/kgrzybek/modular-monolith-with-ddd)
