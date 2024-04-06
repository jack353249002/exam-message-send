// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/kysion/base-library/base_model"
)

type (
	IMessageSend interface {
		// 添加消息
		SetMessage(ctx context.Context, title string, body string) (bool, error)
		// 消息列表
		QueryMessageList(ctx context.Context, info *base_model.SearchParams) (response *sys_model.CoMessageListRes, err error)
		// 添加发送规则
		AddSend(ctx context.Context, title string, messageId int, sendServerId string, receive string) (bool, error)
		// 设置发送消息
		SetSendInfoAction(ctx context.Context, sendId int, status int8) (bool, error)
	}
)

var (
	localMessageSend IMessageSend
)

func MessageSend() IMessageSend {
	if localMessageSend == nil {
		panic("implement not found for interface IMessageSend, forgot register?")
	}
	return localMessageSend
}

func RegisterMessageSend(i IMessageSend) {
	localMessageSend = i
}
