// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

// CoSendList is the golang structure for table co_send_list.
type CoSendList struct {
	Id           int    `json:"id"           description:""`
	SendId       int    `json:"sendId"       description:""`
	Receive      string `json:"receive"      description:""`
	Status       int    `json:"status"       description:""`
	SendServerId int    `json:"sendServerId" description:""`
}