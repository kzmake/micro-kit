package api

import (
	"time"

	"github.com/gin-gonic/gin"
	cli "github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	"github.com/kzmake/micro-kit/pkg/constant"
)

var (
	service = constant.Service.Task
	version = "v0.1.0"
)

// Server はサーバーとして動作するアプリケーションです。
type Server interface {
	Run() error
}

// New はサーバーを生成します。
func New() Server {
	service := web.NewService(
		web.Name("api."+service),
		web.Version(version),

		web.RegisterInterval(10*time.Second), // nolint:gomnd
		web.Registry(etcd.NewRegistry()),
	)
	err := service.Init(
		// init
		web.Action(func(c *cli.Context) {
			// load config
		}),
	)
	if err != nil {
		return nil
	}

	task := NewTaskRouter()
	router := gin.Default()
	router.POST("/tasks/:description", task.CreateTask)
	router.POST("/tasks/", task.CreateTask)

	// Register Handler
	service.Handle("/", router)

	return service
}