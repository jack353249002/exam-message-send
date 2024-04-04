package sys_hook

import (
	"context"
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
	"github.com/jack353249002/exam-message-send/sys_model/sys_enum"
)

//type RoleMemberHookFunc func(ctx context.Context, event sys_enum.RoleMemberChange, role *sys_model.RoleInfo, sysUser *sys_model.SysUser) (bool, error)

// RoleMemberChangeHookFunc 团队成员发生更改处理逻辑
type RoleMemberChangeHookFunc func(ctx context.Context, event sys_enum.RoleMemberChange, role sys_entity.SysRole, sysUser *sys_model.SysUser) (bool, error)
