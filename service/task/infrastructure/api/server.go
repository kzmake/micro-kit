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
	service = constant.TaskAPI
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
		web.Action(func(c *cli.Context) {
		}),
	)
	if err != nil {
		return nil
	}

	task := NewTaskRouter()
	router := gin.Default()
	router.GET("/tasks", task.ListTasks)
	router.POST("/tasks", task.CreateTask)
	router.GET("/tasks/:task_id", task.GetTask)
	router.DELETE("/tasks/:task_id", task.DeleteTask)

	// Register Handler
	service.Handle("/", router)

	return service
}
