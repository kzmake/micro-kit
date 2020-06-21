package api

import (
	"context"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/errors"

	"github.com/kzmake/micro-kit/pkg/constant"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
)

var client = proto.NewTaskService(constant.Service.Task, grpc.NewClient())

// TaskRouter は task に関するルーターの定義です。
type TaskRouter struct{}

// NewTaskRouter はルーターを生成します。
func NewTaskRouter() *TaskRouter {
	return &TaskRouter{}
}

// CreateTask はタスク生成のリクエストを処理します。
func (r *TaskRouter) CreateTask(c *gin.Context) {
	type postRequest struct {
		Description string `json:"description"`
	}
	body := postRequest{}
	if err := c.Bind(&body); err != nil {
		renderError(c, err)
		return
	}

	response, err := client.Create(context.Background(),
		&proto.CreateRequest{Description: &wrappers.StringValue{Value: body.Description}},
	)
	if err != nil {
		renderError(c, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"id":          response.GetResult().GetId().GetValue(),
		"description": response.GetResult().GetDescription().GetValue(),
	})
}

// ListTasks はタスク一覧取得のリクエストを処理します。
func (r *TaskRouter) ListTasks(c *gin.Context) {
	response, err := client.List(context.Background(),
		&proto.ListRequest{},
	)
	if err != nil {
		renderError(c, err)
		return
	}

	tasks := []map[string]string{}
	for _, res := range response.GetResults() {
		task := map[string]string{
			"id":          res.GetId().GetValue(),
			"description": res.GetDescription().GetValue(),
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"tasks": tasks})
}

// GetTask はタスク取得のリクエストを処理します。
func (r *TaskRouter) GetTask(c *gin.Context) {
	taskID := c.Param("task_id")

	response, err := client.Get(context.Background(),
		&proto.GetRequest{Id: &wrappers.StringValue{Value: taskID}},
	)
	if err != nil {
		renderError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"id":          response.GetResult().GetId().GetValue(),
		"description": response.GetResult().GetDescription().GetValue(),
	})
}

// DeleteTask はタスク削除のリクエストを処理します。
func (r *TaskRouter) DeleteTask(c *gin.Context) {
	taskID := c.Param("task_id")

	_, err := client.Delete(context.Background(),
		&proto.DeleteRequest{Id: &wrappers.StringValue{Value: taskID}},
	)
	if err != nil {
		renderError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func renderError(c *gin.Context, err error) {
	var e *errors.Error
	if xerrors.As(err, &e) {
		merr := errors.Parse(err.Error())
		c.JSON(int(merr.GetCode()), map[string]string{
			"code":    merr.GetId(),
			"message": merr.GetDetail(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]string{
		"code":    "InternalServerError",
		"message": "An internal error has occurred. Please try your query again at a later time.",
	})
}
