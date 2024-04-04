package internal

import (
	"github.com/jack353249002/exam-message-send/internal/boot"
	"github.com/jack353249002/exam-message-send/utility/idgen"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/kysion/base-library/utility/env"
)

func init() {
	env.LoadEnv()

	idgen.InitIdGenerator()
	boot.InitIp2region()
	boot.InitCustomRules()
	boot.InitGlobal()
	boot.InitRedisCache()
	boot.InitLogLevelToDatabase()
	boot.InitPermission()
	boot.InitEmail()

}
