package logger

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"

	"github.com/kzmake/micro-kit/pkg/logger/technical"
)

// NewHandlerWrapper は server handler の wrapper を生成します。
func NewHandlerWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			md, _ := metadata.FromContext(ctx)
			technical.Debugf("ServerHandler: Service: %s, Method: %s, ctx: %v", req.Service(), req.Method(), md)
			return fn(ctx, req, rsp)
		}
	}
}

// NewSubscriberWrapper は subscribe の wrapper を生成します。
func NewSubscriberWrapper() server.SubscriberWrapper {
	return func(fn server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, p server.Message) error {
			md, _ := metadata.FromContext(ctx)
			technical.Debugf("ServerSubscriber: Topic: %s, ContentType: %s, Payload: %v, ctx: %v", p.Topic(), p.ContentType(), p.Payload(), md)
			return fn(ctx, p)
		}
	}
}

// NewClientWrapper は client の wrapper を生成します。
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
