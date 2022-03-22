package main

import (
	_ "gf17/boot"
	_ "gf17/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
