// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsLogs is the golang structure for table sys_sms_logs.
type SysSmsLogs struct {
	Id        float64     `json:"id"        description:""`
	Type      string      `json:"type"      description:"短信平台：qyxs：企业信使"`
	Context   string      `json:"context"   description:"短信内容"`
	Mobile    string      `json:"mobile"    description:"手机号"`
	State     string      `json:"state"     description:"发送状态"`
	Result    string      `json:"result"    description:"短信接口返回内容"`
	UserId    int64       `json:"userId"    description:"用户ID"`
	LicenseId int64       `json:"licenseId" description:"主体ID"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
