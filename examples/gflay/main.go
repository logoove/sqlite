package main

import (
	_ "gflay/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"gflay/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
