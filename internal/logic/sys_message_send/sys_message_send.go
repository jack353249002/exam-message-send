package sys_area

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/jack353249002/exam-message-send/sys_model/sys_dao"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
	"github.com/jack353249002/exam-message-send/sys_service"
	"strings"
	"time"
)

type sMessageSend struct {
	conf gdb.CacheOption
}
func init() {
	sys_service.RegisterMessageSend(New())
}
func New() *sMessageSend {
	return &sMessageSend{
		conf: gdb.CacheOption{
			Duration: time.Hour * 24 * 7,
			Force:    false,
		},
	}
}
// 添加消息
func (s *sMessageSend) SetMessage(ctx context.Context,title string,body string) (bool,error) {
	data:=&sys_entity.CoMessage{
		Title: title,
		Body: body,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	_,err:=sys_dao.CoMessage.Ctx(ctx).OmitEmpty().OmitNilData().Data(data).Insert()
	if err==nil{
		return true,nil
	}else{
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, err.Error()), "", sys_dao.CoMessage.Table())
	}
}
// 添加发送规则
func (s *sMessageSend) AddSend(ctx context.Context,title string,messageId int,sendServerId string,receive string) (bool,error) {
	receives := strings.Split(receive, ",")
	err:=sys_dao.CoSend.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		send:=&sys_entity.CoSend{
			Title: title,
			MessageId: messageId,
			SendServerId: sendServerId,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
			SendListCount: len(receives),
		}
		res,err:=sys_dao.CoSend.Ctx(ctx).OmitEmpty().OmitNilData().Data(send).Insert()
		lastInsertId, err := res.LastInsertId()
		var sendList []map[string]interface{}
		for _,val:=range receives{
			dataRow:=map[string]interface{}{"send_id":lastInsertId,"receive":val,"status":0,"send_server_id":0}
			sendList=append(sendList,dataRow)
		}
		if len(sendList)>0 {
			res, err = sys_dao.CoSendList.Ctx(ctx).OmitEmpty().OmitNilData().Data(sendList).Insert()
		}
		return err
	})
	if err==nil{
		return true,nil
	}else{
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, err.Error()), "", sys_dao.CoSendList.Table())
	}
}