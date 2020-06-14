package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v2/client/grpc"

	"github.com/kzmake/micro-kit/pkg/constant"
	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
)

// TaskRouter は task に関するルーターの定義です。
type TaskRouter struct{}

// NewTaskRouter はルーターを生成します。
func NewTaskRouter() *TaskRouter {
	return &TaskRouter{}
}

// CreateTask はタスク生成のリクエストを処理します。
func (r *TaskRouter) CreateTask(c *gin.Context) {
	technical.Infof("Task.Createへのリクエストを受け付けました")

	description := c.Param("description")

	client := proto.NewTaskService(constant.Service.Task, grpc.NewClient())
	response, err := client.Create(context.Background(),
		&proto.CreateRequest{
			Description: &wrappers.StringValue{Value: description},
		},
	)
	if err != nil {
		c.JSON(http.StatusOK, map[string]string{
			"code":    "InternalError",
			"message": "だめ",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"id":          response.GetResult().GetId().GetValue(),
		"description": response.GetResult().GetDescription().GetValue(),
	})
}
