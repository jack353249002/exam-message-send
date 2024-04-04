package sys_hook

import (
	"context"
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_model/sys_enum"
)

type JwtHookFunc func(ctx context.Context, claims *sys_model.JwtCustomClaims) (*sys_model.JwtCustomClaims, error)

type JwtHookInfo sys_model.HookEventType[sys_enum.UserType, JwtHookFunc]
