// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CoMessage is the golang structure of table co_message for DAO operations like Where/Data.
type CoMessage struct {
	g.Meta    `orm:"table:co_message, do:true"`
	Id        interface{} //
	Title     interface{} //
	Body      interface{} //
	Attach    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}