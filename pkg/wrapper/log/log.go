package log

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"

	"github.com/kzmake/micro-kit/pkg/logger"
)

// NewHandlerWrapper ...
func NewHandlerWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			md, _ := metadata.FromContext(ctx)
			logger.Debugf("ServerHandler: Service: %s, Method: %s, ctx: %v", req.Service(), req.Method(), md)
			return fn(ctx, req, rsp)
		}
	}
}

// NewSubscriberWrapper ...
func NewSubscriberWrapper() server.SubscriberWrapper {
	return func(fn server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, p server.Message) error {
			md, _ := metadata.FromContext(ctx)
			logger.Debugf("ServerSubscriber: Topic: %s, ContentType: %s, Payload: %v, ctx: %v", p.Topic(), p.ContentType(), p.Payload(), md)
			return fn(ctx, p)
		}
	}
}

// NewClientWrapper ...
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
