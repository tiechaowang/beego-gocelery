package routers

import (
	"firstproject/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
	 _, ok := ctx.Input.Session("uid").(int)
    if !ok && ctx.Request.RequestURI != "/" {
        ctx.Redirect(302, "/")
    }
}


func init() {
	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
    // beego.Router("/", &controllers.MainController{})
    beego.Router("/", &controllers.TestController{})
    beego.Router("/user", &controllers.MainController{})
}
