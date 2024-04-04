package sys_hook

import (
	"context"
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_model/sys_enum"
)

type AuthHookFunc func(ctx context.Context, actionType sys_enum.AuthActionType, info *sys_model.SysUser) error
type AuthHookInfo struct {
	Key      sys_enum.AuthActionType
	Value    AuthHookFunc
	UserType sys_enum.UserType `json:"userType" dc:"用户类型"`
}
