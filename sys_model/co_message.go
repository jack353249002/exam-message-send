// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_model

import (
)

// CoMessage is the golang structure of table co_message for DAO operations like Where/Data.
type CoMessage struct {
	Title     string `json:"title" v:"required#请输入标题"  dc:"标题"`
	Body      string `json:"body" v:"required#请输入内容"  dc:"内容"`
}