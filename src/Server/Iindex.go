package Server

import (
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"g"
	"github.com/yinheli/qqwry"
	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
)
const (
	username="admin"   //用户名
	password="29295842"		//879301117密码
)

var(
	IP_qqwry = qqwry.NewQQwry("qqwry.dat")
)

type Iindex struct {
	beego.Controller
}

func (this *Iindex) Login() {
	this.TplName = "template/login.html"
	//验证登陆
	Form_username := this.Ctx.GetCookie("username") //获取表单
	Form_password := this.Ctx.GetCookie("password") //获取表单
	if Form_username==username && Form_password==password{
		this.Get()
		return
	}

	Form_username = this.Input().Get("username") //获取表单
	Form_password = this.Input().Get("password") //获取表单
	if Form_username==username && Form_password==password{
		this.Ctx.SetCookie("username", Form_username, "/") //设置Cookie
		this.Ctx.SetCookie("password", Form_password, "/") //设置Cookie
		this.Get()
		return
	}else{
		this.Ctx.SetCookie("username", "", "/") //设置Cookie
		this.Ctx.SetCookie("password", "", "/") //设置Cookie
	}
	return
}

func (this *Iindex) Get() {  //
	path_url := this.Ctx.Request.URL.Path
	s := fmt.Sprintf("========%s========\r\n", path_url)
	fmt.Printf(s)
	this.Data["User_username"] = username   //用户名
	if len(path_url) <= 4 {
		this.Iindex()
		return
	}
	if path_url == "/Login.html" || path_url == "/index.html" { //首页      登陆窗口  path_url == "/" ||
		this.Iindex()
		return
	}
	if path_url=="/shell.html"{
		this.shell()
		return
	}
	if path_url=="/get_shell.html"{
		this.get_shell()
		return
	}
	if path_url=="/send_shll.html"{
		this.send_shll()
		return
	}

	this.Ctx.WriteString("")   //必须返回一个数据
	return
}

func (this *Iindex) send_shll() {   //发送数据
	uuid:=g.Url_Parse(this.Input().Get("uuid"))
	id:=g.Url_Parse(this.Input().Get("id"))
	shell:=g.Url_Parse(this.Input().Get("zl"))
	send_bool:=false
	for value := range manager.clients {
		if value.uuid == uuid && value.id == id {
			jsonMessage, _ := json.Marshal(&Message{Action: "Shell",Data:g.Base64_Encode(shell)})   ///
			value.socket.WriteMessage(websocket.TextMessage, jsonMessage)   //发送数据
			send_bool=true
			break //#跳出
		}
	}
	if(send_bool){
		this.Ctx.WriteString("ok")   //必须返回一个数据
	}else{
		this.Ctx.WriteString("发送失败")   //必须返回一个数据
	}
	return
}

func (this *Iindex) get_shell() {  //读取数据
	uuid:=g.Url_Parse(this.Input().Get("uuid"))
	id:=g.Url_Parse(this.Input().Get("id"))
	data:=""
	for value := range manager.clients {
		if value.uuid == uuid && value.id == id {
			//value.socket.WriteMessage(websocket.TextMessage, message)   //发送数据
			for _, value := range value.receive_data {
				//fmt.Printf(value)
				data = fmt.Sprintf("%s\r\n%s\r\n", data,value)
			}
			value.receive_data=[]string{}
			break //#跳出
		}
	}
	srcDecoder := mahonia.NewDecoder("gbk")
	desDecoder := mahonia.NewDecoder("utf-8")
	resStr:= srcDecoder.ConvertString(data)
	_, resBytes, _ := desDecoder .Translate([]byte(resStr), true)
	data = string(resBytes)
	this.Ctx.WriteString(data)   //必须返回一个数据
	return
}

func (this *Iindex) shell() {
	uuid:=g.Url_Parse(this.Input().Get("uuid"))
	id:=g.Url_Parse(this.Input().Get("id"))
	this.Data["uuid"] = uuid   ////远程ID
	this.Data["id"] = id   // //本地ID

	this.Data["ip_wlwz"] =""    //物理位置
	this.Data["ip"] = ""	//IP地址
	this.Data["u_os"] = ""  //操作系统
	this.Data["did"] = ""   //代理Id
	this.Data["bz_name"] = ""  //备注名称
	this.Data["update_time"] = ""      //更新时间
	this.Data["add_time"] = ""        //添加时间
	//================
	//查询机器是否存在
	cx_bool:=false
	for value := range manager.clients {
		if value.uuid==uuid && value.id==id {
			cx_bool=true
			//===============
			IP_qqwry.Find(value.ip)
			this.Data["ip_wlwz"] = fmt.Sprintf("%v-%v", IP_qqwry.Country,IP_qqwry.City)  //物理位置
			this.Data["ip"] = value.ip	//IP地址
			this.Data["u_os"] = value.u_os  //操作系统
			this.Data["did"] = value.did   //代理Id
			this.Data["bz_name"] = value.bz_name  //备注名称
			this.Data["update_time"] = g.Unix_Time(value.update_time)      //更新时间
			this.Data["add_time"] = g.Unix_Time(value.add_time)       //添加时间
			//===============
			break //#跳出
		}
	}
	if cx_bool==false{
		this.Ctx.WriteString("当前主机不存在")   //必须返回一个数据
		return
	}
	//================
	//s := fmt.Sprintf("====%s====%s===", uuid,id)
	//this.Ctx.WriteString(s)   //必须返回一个数据
	this.TplName = "template/shell.html"
	return
}

func (this *Iindex) Iindex() {

	//this.Ctx.WriteString("xxxxxxxxxxxxx")   //必须返回一个数据
	//========================================
	//主机列表
	host_map := make(map[string]interface{}) //
	//fmt.Println(reflect.TypeOf(manager.register))
	i:=0
	for value := range manager.clients {
		//fmt.Println(reflect.TypeOf(value.add_time))
		rows_map := make(map[string]string)
		rows_map["add_time"] = g.Unix_Time(value.add_time)   //添加时间
		rows_map["update_time"] = g.Unix_Time(value.update_time)  //更新时间
		rows_map["ip"] = value.ip	 //IP地址
		IP_qqwry.Find(value.ip)
		rows_map["ip_wlwz"] = fmt.Sprintf("%v-%v", IP_qqwry.Country,IP_qqwry.City)  //物理位置
		rows_map["id"] = value.id //本地ID
		rows_map["uuid"] = value.uuid //远程ID
		rows_map["u_os"] = value.u_os //操作系统
		//rows_map["u_bh"] = value.u_bh //编号
		rows_map["did"] = value.did //代理ID
		rows_map["bz_name"] = value.bz_name //备注名称
		host_map[string(i)] = rows_map
		i++
	}
	this.Data["host_map"] = host_map   //主机列表
	//========================================
	this.TplName = "template/index.html"
	return
}


/*
func (this *Iindex) Login() {
	defer g.Panic_Err() //异常处理
	host_map := make(map[string]interface{}) //
	//fmt.Println(reflect.TypeOf(manager.register))
	i:=0
	for value := range manager.clients {
		//fmt.Println(reflect.TypeOf(value.add_time))
		rows_map := make(map[string]string)
		rows_map["add_time"] = g.Unix_Time(value.add_time)   //添加时间
		rows_map["update_time"] = g.Unix_Time(value.update_time)  //更新时间
		rows_map["ip"] = value.ip	 //IP地址
		rows_map["id"] = value.id //本地ID
		rows_map["uuid"] = value.uuid //远程ID
		host_map[string(i)] = rows_map
		i++
	}
	s := fmt.Sprintf("%+v", host_map)
	fmt.Printf(s)
	this.Ctx.WriteString("")   //必须返回一个数据
}
*/




