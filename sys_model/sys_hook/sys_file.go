package sys_hook

import (
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
	"github.com/jack353249002/exam-message-send/sys_model/sys_enum"
)

type FileHookFunc sys_model.HookFunc[sys_enum.UploadEventState, sys_entity.SysFile]
type FileHookInfo sys_model.HookEventType[sys_enum.UploadEventState, FileHookFunc]
