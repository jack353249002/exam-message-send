package sys_controller

import (
	"context"
	"github.com/jack353249002/exam-message-send/api_v1/sys_api"
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_service"
)

// SysArea 地区
var SysArea = cSysArea{}

type cSysArea struct{}

// GetAreaListByParentId 获取属于父级ID的地区列表
func (c *cSysArea) GetAreaListByParentId(ctx context.Context, req *sys_api.GetAreaListByParentIdReq) (*sys_model.AreaListRes, error) {
	return sys_service.Area().GetAreaListByParentId(ctx, req.ParentId)
}
