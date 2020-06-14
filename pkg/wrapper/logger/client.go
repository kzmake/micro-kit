package logger

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"

	"github.com/kzmake/micro-kit/pkg/logger/technical"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	technical.Debugf("ClientCaller: Service: %s, Method: %s, ctx: %v", req.Service(), req.Method(), md)
	return c.Client.Call(ctx, req, rsp, opts...)
}

func (c *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {
	md, _ := metadata.FromContext(ctx)
	technical.Debugf("ClientPublisher: Topic: %s, ContentType: %s, Payload: %v, ctx: %v", p.Topic(), p.ContentType(), p.Payload(), md)
	return c.Client.Publish(ctx, p, opts...)
}
