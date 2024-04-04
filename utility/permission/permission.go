package permission

import (
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
	"github.com/jack353249002/exam-message-send/utility/idgen"
	"github.com/kysion/base-library/utility/base_permission"
)

func initFactory() {
	if base_permission.Factory == nil {
		base_permission.Factory = func() base_permission.IPermission {
			return &sys_model.SysPermissionTree{
				SysPermission: &sys_entity.SysPermission{},
			}
		}
	}
}

func New(id int64, identifier string, name string, description ...string) base_permission.IPermission {
	var desc string
	if len(description) > 0 {
		desc = description[0]
	}

	initFactory()

	return base_permission.Factory().SetId(id).SetIdentifier(identifier).SetName(name).SetDescription(desc)
}

// NewInIdentifier 构造权限信息
func NewInIdentifier(identifier string, name string, description ...string) base_permission.IPermission {
	var desc string

	if len(description) > 0 {
		desc = description[0]
	}

	initFactory()

	return base_permission.Factory().SetId(idgen.NextId()).SetIdentifier(identifier).SetName(name).SetDescription(desc)
}
