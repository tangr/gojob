package main

import (
	_ "gojob/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gojob/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
