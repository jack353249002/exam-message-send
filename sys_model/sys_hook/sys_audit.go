package sys_hook

import (
	"context"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
	"github.com/jack353249002/exam-message-send/sys_model/sys_enum"
)

type AuditHookFunc func(ctx context.Context, state sys_enum.AuditEvent, info *sys_entity.SysAudit) error
type AuditHookInfo struct {
	Key      sys_enum.AuditEvent
	Value    AuditHookFunc
	Category int `json:"category" dc:"业务类别"`
}
