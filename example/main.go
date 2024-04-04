package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "github.com/jack353249002/exam-message-send"
	"github.com/jack353249002/exam-message-send/example/internal/boot"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	boot.Main.Run(gctx.New())
}
