package sys_model

import (
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysOrganizationInfo struct {
	Id          int64  `json:"id"          dc:""`
	Name        string `json:"name"        dc:"名称" v:"required|length:2,32#名称不能为空|名称长度仅限2~32个字符"`
	ParentId    int64  `json:"parentId"    dc:"父级ID" v:"min:0"`
	Description string `json:"description" dc:"描述"`
}

type SysOrganizationTree struct {
	SysOrganizationInfo
	CascadeDeep int                    `json:"cascadeDeep" dc:"级联深度" v:"min:0"`
	Children    []*SysOrganizationTree `json:"children" orm:"-" dc:"下级组织架构"`
}

type OrganizationInfoListRes base_model.CollectRes[*sys_entity.SysOrganization]

type OrganizationInfo sys_entity.SysOrganization
