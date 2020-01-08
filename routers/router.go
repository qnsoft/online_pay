package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//加入跨域权限
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		AllowCredentials: true}))
	beego.SetStaticPath("/upload", "upload") //重定向静态路径
	beego.SetStaticPath("/views", "views")   //重定向静态路径
	//-------------------------基础接口开始-----------------------------------//
	// //根目录
	// beego.Router("/", &controllers.Default_Controller{}, "get,post:Get")
	// //先获取access_token
	// beego.Router("/access_token", &Token.AccesstokenController{}, "post:Access_Token")
	// //token测试
	// beego.Router("/testtoken", &controllers.Default_Controller{}, "get,post:TestToken")
	//------------------------在线支付--工易付接口-------------------------------------//
	//router_pay.GongYiFu()
}
