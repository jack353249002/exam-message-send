// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CoSend is the golang structure for table co_send.
type CoSend struct {
	Id            int         `json:"id"            description:""`
	Title         string      `json:"title"         description:""`
	MessageId     int         `json:"messageId"     description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:""`
	SendModel     int         `json:"sendModel"     description:""`
	SendServerId  string      `json:"sendServerId"  description:""`
	DispatchModel int         `json:"dispatchModel" description:""`
	Status        int         `json:"status"        description:""`
	SendListCount int         `json:"sendListCount" description:""`
	SuccessCount  int         `json:"successCount"  description:""`
}
