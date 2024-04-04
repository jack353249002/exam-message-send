// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/jack353249002/exam-message-send/sys_model/sys_dao/internal"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
)

type CoSendListDao = dao_interface.TIDao[internal.CoSendListColumns]

func NewCoSendList(dao ...dao_interface.IDao) CoSendListDao {
	return (CoSendListDao)(internal.NewCoSendListDao(dao...))
}

var (
	CoSendList = NewCoSendList()
)