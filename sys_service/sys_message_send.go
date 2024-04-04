// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"
)

type (
	IMessageSend interface {
		// 添加消息
		SetMessage(ctx context.Context, title string, body string) (bool, error)
		// 添加发送规则
		AddSend(ctx context.Context, title string, messageId int, sendServerId string, receive string) (bool, error)
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
