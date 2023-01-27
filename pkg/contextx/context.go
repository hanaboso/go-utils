package contextx

import (
	"context"
	"time"
)

func WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func WithTimeoutCtx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

func WithTimeoutSeconds(seconds int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
}

func WithTimeoutSecondsCtx(seconds int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	return ctx
}
