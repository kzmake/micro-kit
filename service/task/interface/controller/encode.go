package controller

import (
	"context"

	merrors "github.com/micro/go-micro/v2/errors"

	"github.com/kzmake/micro-kit/service/task/domain/errors"
)

// https://pfs.nifcloud.com/api/hatoba/errors.htm を参考に設定
func encodeError(_ context.Context, err error) error {
	code := errors.GetCode(err)

	switch code {
	case errors.IllegalInputBody:
		return merrors.BadRequest(code.String(), "The request body is not appropriate.")
	case errors.IllegalInputTaskID:
		return merrors.BadRequest(code.String(), "The requested id is invalid.")
	case errors.IllegalInputDescription:
		return merrors.BadRequest(code.String(), "The requested id is invalid.")
	case errors.NotFoundTask:
		return merrors.NotFound(code.String(), "The task does not found.")
	case errors.DuplicateTask:
		return merrors.BadRequest(code.String(), "The task already exists.")
	}

	return merrors.InternalServerError("InternalServerError", "An internal error has occurred. Please try your query again at a later time.")
}
