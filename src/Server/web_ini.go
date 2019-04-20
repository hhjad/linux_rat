package Server

import (
	"github.com/astaxie/beego"
	"strconv"
	"g"
	//"gcmdb/controllers/websocket/wssaction"
	"fmt"
)

//=====
var (
	page_index = 20 //每页多少条数据
	Web_debug = false //false为调试模式
	//Web_debug = true //true为发布模式
)

func Web_run(Server, port string){  //启动网站
	beego.BConfig.Listen.ServerTimeOut = 15 //设置 HTTP 的超时时间，默认是 0，不超时。
	//beego.BConfig.Listen.EnableHTTP = true //是否启用 HTTP 监听，默认是 true。
	if Server != "" {
		beego.BConfig.Listen.HTTPAddr = Server //应用监听地址，默认为空，监听所有的网卡 IP。
	}

	i, err := strconv.Atoi(port)
	if err != nil {
		beego.BConfig.Listen.HTTPPort = 800 //应用监听端口,默认为 8080。
	} else {
		beego.BConfig.Listen.HTTPPort = i
	}
	//fmt.Printf("========beego.VERSION  %v======\n", beego.VERSION)
	beego.BConfig.AppName = "jq"                               //应用名称，默认是 beego。通过 bee new 创建的是创建的项目名。
	beego.BConfig.ServerName = "xxxx.zzkey.com"                       //beego 服务器默认在请求的时候输出 server 为 beego。
	beego.BConfig.WebConfig.Session.SessionName = "seosessionID"      //存在客户端的 cookie 名称，默认值是 beegosessionID。
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 10  //session 过期时间，默认值是 3600 秒。
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600 * 10 //session 默认存在客户端的 cookie 的时间，默认值是 3600 秒。
	beego.BConfig.WebConfig.Session.SessionDomain = "" //session cookie 存储域名, 默认空。
	//beego.BConfig.WebConfig.ViewsPath = "myview"    //允许跨域
	//跨域设置


	g.Log_msgbox("==Debug mode==")
	//运行模式，可选值为 prod, dev 或者 test. 默认是 dev, 为开发模式
	beego.BConfig.RunMode = "dev" //"test"
	beego.BConfig.WebConfig.ViewsPath = "www"             //模板路径，默认值是 views。
	beego.SetStaticPath("/public/", "./www/public")             //静态目录
	beego.Router("/wss", &WebSocket_Server{}, "*:ServeWs")
	beego.Router("/*", &Iindex{}, "*:Login")
	beego.Router("/close/", &Iindex{}, "*:Close")

	go manager.start()  //启动webSOCKET

	go beego.Run()

	open_url:=""
	if Server != "" {
		open_url=fmt.Sprintf("%s:%d",Server, i)
		//g.Log_msgbox("")
		g.Log_msgbox(fmt.Sprintf("run server===%v", open_url))
	} else {
		open_url=fmt.Sprintf("127.0.0.1:%d", i)
		//g.Log_msgbox("")
		g.Log_msgbox( fmt.Sprintf("run server===localhost:%d", i))
		g.Log_msgbox( fmt.Sprintf("run server===%v", open_url))
	}


}

func (this *Iindex) Close() {
	defer g.Panic_Err()           //异常处理
	this.Ctx.SetCookie("username", "", "/") //设置Cookie
	this.Ctx.SetCookie("password", "", "/") //设置Cookie
	//this.return_alert_href("退出网站后台","/Login.html")   //返回信息
	this.Redirect("/", 302) //支持 401、403、404、500、503
	return
}

