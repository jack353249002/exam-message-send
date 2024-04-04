package sys_hook

import (
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_model/sys_enum"
)

type UserHookFunc sys_model.HookFuncRes[sys_enum.UserEvent, sys_model.SysUser]
type UserHookInfo sys_model.HookEventType[sys_enum.UserEvent, UserHookFunc]
