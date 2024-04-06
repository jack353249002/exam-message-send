package dispatch

import (
	"context"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
)

const ErrorCheckCoolTime = 120

type EmailDispatcher interface {
	MoveIndex()
	FillServer([]sys_entity.CoSmtpServer)
	GetServer() (SmtpServerInfo, int, bool)
	GetServerListen() (sys_entity.CoSmtpServer, int, bool, bool)
	RemoveServer(int, int, string)
	Init(int, context.Context)
	SetIsWaitServer()
	SetNotWaitServer()
	GetIsWaitServer() bool
}
