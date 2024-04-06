package sys_controller

import (
	"context"
	"github.com/jack353249002/exam-message-send/api_v1"
	"github.com/jack353249002/exam-message-send/api_v1/sys_api"
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_service"
)

// SysMessageSend 消息发送
var SysMessageSend = cMessageSend{}

type cMessageSend struct{}

// SetMessage 设置消息
func (c *cMessageSend) SetMessage(ctx context.Context, req *sys_api.SetMessageInfoReq) (api_v1.BoolRes, error) {
	ok, err := sys_service.MessageSend().SetMessage(ctx, req.Title, req.Body)
	return ok == true, err
}

// SetSend 设置发送规则
func (c *cMessageSend) SetSend(ctx context.Context, req *sys_api.AddSendReq) (api_v1.BoolRes, error) {
	ok, err := sys_service.MessageSend().AddSend(ctx, req.Title, req.MessageId, req.SendServerId, req.Receive)
	return ok == true, err
}

// SetSend 发送消息
func (c *cMessageSend) SetSendInfoAction(ctx context.Context, req *sys_api.SetSendActionReq) (api_v1.BoolRes, error) {
	ok, err := sys_service.MessageSend().SetSendInfoAction(ctx, req.Id, int8(req.Status))
	return ok == true, err
}

// QueryUserList 获取用户|列表
func (c *cMessageSend) QueryMessageList(ctx context.Context, req *sys_api.QueryMessageListReq) (*sys_model.CoMessageListRes, error) {
	return sys_service.MessageSend().QueryMessageList(ctx, &req.SearchParams)
}
