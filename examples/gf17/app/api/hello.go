package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	id, _ := g.Model("user").Data(g.Map{"name": "john", "age": 1}).InsertAndGetId()
	r.Response.Writeln("Hello World!", id)
}
