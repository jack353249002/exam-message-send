package sys_message_send

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/jack353249002/exam-message-send/internal/logic/send/sendunit"
	"github.com/jack353249002/exam-message-send/sys_model"
	"github.com/jack353249002/exam-message-send/sys_model/sys_dao"
	"github.com/jack353249002/exam-message-send/sys_model/sys_do"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
	"github.com/jack353249002/exam-message-send/sys_service"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"math"
	"strconv"
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
func (s *sMessageSend) SetMessage(ctx context.Context, title string, body string) (bool, error) {
	data := &sys_entity.CoMessage{
		Title:     title,
		Body:      body,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	_, err := sys_dao.CoMessage.Ctx(ctx).OmitEmpty().OmitNilData().Data(data).Insert()
	if err == nil {
		return true, nil
	} else {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, err.Error()), "", sys_dao.CoMessage.Table())
	}
}

// 消息列表
func (s *sMessageSend) QueryMessageList(ctx context.Context, info *base_model.SearchParams) (response *sys_model.CoMessageListRes, err error) {
	if info != nil {
		newFields := make([]base_model.FilterInfo, 0)

		// 移除类型筛选条件
		for _, field := range info.Filter {
			newFields = append(newFields, field)
		}
	}

	result, err := daoctl.Query[*sys_model.CoMessage](sys_dao.CoMessage.Ctx(ctx), info, false)

	newList := make([]*sys_model.CoMessage, 0)
	if result != nil && result.Records != nil && len(result.Records) > 0 {
		for _, message := range result.Records {
			newList = append(newList, message)
		}
	}

	if newList != nil && len(newList) > 0 {
		result.Records = newList
	}

	return (*sys_model.CoMessageListRes)(result), err

}

// 添加发送规则
func (s *sMessageSend) AddSend(ctx context.Context, title string, messageId int, sendServerId string, receive string) (bool, error) {
	receives := strings.Split(receive, ",")
	err := sys_dao.CoSend.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		send := &sys_entity.CoSend{
			Title:         title,
			MessageId:     messageId,
			SendServerId:  sendServerId,
			CreatedAt:     gtime.Now(),
			UpdatedAt:     gtime.Now(),
			SendListCount: len(receives),
		}
		res, err := sys_dao.CoSend.Ctx(ctx).OmitEmpty().OmitNilData().Data(send).Insert()
		lastInsertId, err := res.LastInsertId()
		var sendList []map[string]interface{}
		for _, val := range receives {
			dataRow := map[string]interface{}{"send_id": lastInsertId, "receive": val, "status": 0, "send_server_id": 0}
			sendList = append(sendList, dataRow)
		}
		if len(sendList) > 0 {
			res, err = sys_dao.CoSendList.Ctx(ctx).OmitEmpty().OmitNilData().Data(sendList).Insert()
		}
		return err
	})
	if err == nil {
		return true, nil
	} else {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, err.Error()), "", sys_dao.CoSendList.Table())
	}
}

// 设置发送消息
func (s *sMessageSend) SetSendInfoAction(ctx context.Context, sendId int, status int8) (bool, error) {
	var coSend sys_entity.CoSend
	var coSendList []sys_entity.CoSendList
	sys_dao.CoSend.Ctx(ctx).Where(sys_do.CoSend{Id: sendId}).Scan(&coSend)
	if coSend.Id == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "发送消息规则不存在"), "", sys_dao.CoSend.Table())
	} else {
		if status == 1 || status == -1 || status == -2 {
			key := strconv.Itoa(sendId)
			emailSend, sendUnitHave := sendunit.GetSendEmailUnitPool(key)
			switch status {
			case 1:
				if sendUnitHave {
					emailSend.Start()
				} else {
					sys_dao.CoSendList.Ctx(ctx).Where("send_id=? AND status=0", sendId).Scan(&coSendList)
					num := math.Ceil(float64(len(coSendList)) / 2.0)
					sendUnit := sendunit.SendFactory(coSend.SendModel)
					if sendUnit == nil {
						return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "发送模型不存在"), "", sys_dao.CoSend.Table())
					}
					sendUnit.Init(int(num), &coSend, &coSendList, context.Background())
				}
			case -2:
				if sendUnitHave {
					emailSend.Pause()
				} else {
					return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "控制单元不存在"), "", sys_dao.CoSend.Table())
				}
			case -1:
				if sendUnitHave {
					emailSend.Stop()
				} else {
					return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "控制单元不存在"), "", sys_dao.CoSend.Table())
				}
			}
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "控制单元不存在"), "", sys_dao.CoSend.Table())
		}
	}
	return true, nil
}
