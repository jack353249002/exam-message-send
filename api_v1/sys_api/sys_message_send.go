package sys_api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/kysion/base-library/base_model"
)

type SetMessageInfoReq struct {
	g.Meta `path:"/setMessage" method:"post" summary:"设置消息信息" tags:"消息推送系统"`
	Title  string `json:"title"    description:"标题"  default:"" `
	Body   string `json:"body"    description:"内容"  default:""`
	sys_model.CoMessage
}
type QueryMessageListReq struct {
	g.Meta `path:"/queryMessageList" method:"post" summary:"获取消息|列表" tags:"消息推送系统"`
	base_model.SearchParams
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}
type SendReq struct {
	Title        string `json:"title"    description:"标题"  default:"" `
	MessageId    int    `json:"message_id"    description:"消息id"  default:""`
	SendServerId string `json:"send_server_id"    description:"发送服务器id"  default:""`
	Receive      string `json:"receive"    description:"接收账号"  default:""`
	sys_model.CoSend
}
type AddSendReq struct {
	g.Meta `path:"/addSend" method:"post" summary:"设置发送" tags:"消息推送系统"`
	*SendReq
}
type UpdateSendReq struct {
	g.Meta `path:"/updateSend" method:"post" summary:"设置发送" tags:"消息推送系统"`
	*SendReq
	Id int64 `json:"id"    description:"id"  default:"" `
}
type SetSendActionReq struct {
	g.Meta `path:"/setSendAction" method:"post" summary:"发送指令" tags:"消息推送系统"`
	*sys_model.SetCoSendAction
}
