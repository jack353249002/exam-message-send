package sys_model


type CoSend struct {
	Title  string `json:"title" v:"required#请输入标题"  dc:"标题"`
	MessageId   int `json:"message_id" v:"required#请输入消息id"  dc:"消息id"`
	SendServerId   string  `json:"send_server_id" v:"required#请输入服务器id"  dc:"服务器id"`
	Receive   string  `json:"receive" v:"required#请输入接收者信息"  dc:"接收者信息 信息1,信息2"`
}
type  UpdateCoSend struct {
	*CoSend
	Id  int `json:"id" v:"required#请输入send_id"  dc:"send_id"`
}