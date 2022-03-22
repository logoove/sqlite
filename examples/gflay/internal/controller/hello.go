package controller

import (
	"context"
	"gflay/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (h *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	id, _ := g.Model("user").Data(g.Map{"name": "john", "age": 1}).InsertAndGetId()
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!", id)
	return
}
