package sys_model

import (
	"context"
)

type HookFunc[S any, T any] func(ctx context.Context, state S, info T) error

type HookFuncRes[S any, T any] func(ctx context.Context, state S, info T) (T, error)

type HookEventType[S any, F any] KeyValueT[S, F]
