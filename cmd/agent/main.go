package main

import (
	_ "gojob/internal/boot"
	_ "gojob/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gojob/internal/cmd"
)

func main() {
	cmd.AgentMain.Run(gctx.GetInitCtx())
}
