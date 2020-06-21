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
- [x] Microservices with micro/go-micro
- [ ] gRPC microservices
- [x] Support HTTP-API
- [ ] Support CLI
- [ ] Support NATS
- [ ] Kubernetes with Kustomize
- [x] Input Validation with PGV
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
# grpc
go run service/task/cmd/srv/main.go service/task/cmd/srv/plugin.go --registry=etcd --server_address=localhost:3001

# api
go run service/task/cmd/api/main.go --registry=etcd --server_address=localhost:3000
```

### Request

use grpc
```sh
export MICRO_REGISTRY=etcd

# use micro
micro call kzmake.microkit.task.v1 TaskService.Create '{"description": "hoge"}'
{
        "result": {
                "id": "01EBC02Q7WJTTNT29CT2FVTZDG",
                "description": "hoge",
                "created_at": "2020-06-21T17:51:07Z",
                "updated_at": "2020-06-21T17:51:07Z"
        }
}

micro call kzmake.microkit.task.v1 TaskService.Get '{"id": "01EBC02Q7WJTTNT29CT2FVTZDG"}'
{
        "result": {
                "id": "01EBC02Q7WJTTNT29CT2FVTZDG",
                "description": "hoge",
                "created_at": "2020-06-21T17:51:07Z",
                "updated_at": "2020-06-21T17:51:07Z"
        }
}

micro call kzmake.microkit.task.v1 TaskService.List
{
        "results": [
                {
                        "id": "01EBC02Q7WJTTNT29CT2FVTZDG",
                        "description": "hoge",
                        "created_at": "2020-06-21T17:51:07Z",
                        "updated_at": "2020-06-21T17:51:07Z"
                }
        ]
}

micro call kzmake.microkit.task.v1 TaskService.Delete '{"id": "01EBC02Q7WJTTNT29CT2FVTZDG"}'
{}


# use grpcurl
grpcurl -plaintext -proto service/task/interface/proto/service.proto -d '{"id": "01EBC02Q7WJTTNT29CT2FVTZDG"}' localhost:3001 kzmake.microkit.task.v1.TaskService/Get
{
  "result": {
    "id": "01EBBZ73Y8FW8S02Z4ZWGAMQD1",
    "description": "ppp",
    "createdAt": "2020-06-21T17:36:03Z",
    "updatedAt": "2020-06-21T17:36:03Z"
  }
}
```

use http
```sh
http http://localhost:3000/tasks description=fuga
HTTP/1.1 201 Created
Content-Length: 56
Content-Type: application/json; charset=utf-8
Date: Sun, 21 Jun 2020 17:59:27 GMT

{
    "description": "fuga",
    "id": "01EBC0HZ2F481ENFBNE0FW0JXK"
}

http http://localhost:3000/tasks/01EBC0HZ2F481ENFBNE0FW0JXK
HTTP/1.1 200 OK
Content-Length: 56
Content-Type: application/json; charset=utf-8
Date: Sun, 21 Jun 2020 18:01:16 GMT

{
    "description": "fuga",
    "id": "01EBC0HZ2F481ENFBNE0FW0JXK"
}

http http://localhost:3000/tasks
HTTP/1.1 200 OK
Content-Length: 178
Content-Type: application/json; charset=utf-8
Date: Sun, 21 Jun 2020 18:00:35 GMT

{
    "tasks": [
        {
            "description": "fuga",
            "id": "01EBC0HZ2F481ENFBNE0FW0JXK"
        }
    ]
}

http DELETE http://localhost:3000/tasks/01EBC0HZ2F481ENFBNE0FW0JXK
HTTP/1.1 204 No Content
Date: Sun, 21 Jun 2020 18:01:37 GMT
```

## Credits

- Inspired by [@xmlking's micro-starter-kit](https://github.com/xmlking/micro-starter-kit)
- Inspired by [@kgrzybek's modular-monolith-with-ddd](https://github.com/kgrzybek/modular-monolith-with-ddd)
