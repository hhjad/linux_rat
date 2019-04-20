/*************************************************************
     FileName: gcmdb->websocket->websocket.go
         Date: 2018/3/15 下午2:29
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http:// blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package wssaction

import (
	"gcmdb/common/utils"
	"gcmdb/controllers/websocket/agentInfo"
	"gcmdb/controllers/websocket/wss"
	"gcmdb/models/hardware"
	"gcmdb/models/response"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

func (this *WebSocketController) Get() {
	beego.Info(this.Ctx.Input.CruSession.SessionID())
	this.SetSession("WebSocketSession", this.Ctx.Input.CruSession.SessionID())
	this.Data["WebSocketSession"] = this.Ctx.Input.CruSession.SessionID()
	this.TplName = "websocket.html"
	this.Data["IsWebSocket"] = true
}

func (this *WebSocketController) Info() {
	r := response.NewSuccess()
	defer func() {
		this.Data["json"] = &r
		this.ServeJSON()
		return
	}()
	agentDown, sm, er := hardware.QueryAgentDown()
	var Agent AgentData
	var AgentDown AgentDict
	if er != nil {
		r.Msg = "请求异常"
		r.Status = "403"
		return
	}
	for _, v := range agentDown {
		AgentDown.Eth1 = v.Eth1
		AgentDown.ID = v.ID
		Agent = append(Agent, AgentDown)
	}
	var result AgentResult
	result.Up = &agentInfo.AgentMap.Data
	result.Down = Agent
	r.Data = result
	r.Msg = "在线主机数量: " + strconv.Itoa(agentInfo.AgentMap.LenMap()) + " 台" + " 异常主机: (" + strconv.Itoa(sm) + ") 台"
	r.Status = "200"
	beego.Debug("ActiveCount: ", utils.RedisCli.Stats().ActiveCount, "IdleCount:", utils.RedisCli.Stats().IdleCount)
	return
}

func (this *WebSocketController) ServeWs() {
	var gcmdbLoginStatus bool
	ssid := this.Ctx.Input.CruSession.SessionID()
	cookieSec := beego.AppConfig.String("cookie_secret")
	uid, _ := this.Ctx.GetSecureCookie(cookieSec, "user_id")

	if ssid != "" && uid != "" {
		gcmdbLoginStatus = true
	}
	//beego.Debug("连接websocket ip:", this.Ctx.Request.RemoteAddr)

	//beego.Debug("连接websocket ip: ", this.Ctx.Request.Header.Get("X-Forwarded-For-Agent"))
	r := response.NewSuccess()
	ws, err := wss.Upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(this.Ctx.Request.Header.Get("X-Forwarded-For-Agent"), "连接服务器失败")
		this.Data["json"] = &r
		this.ServeJSON()
		return
	}
	c := &wss.Connection{}
	if ssid != "" && uid != "" {
		c = &wss.Connection{Send: make(chan []byte, 256), Ws: ws, Auth: true}
	} else {
		c = &wss.Connection{Send: make(chan []byte, 256), Ws: ws, Auth: false}
	}
	webSocketId := this.Ctx.Input.CruSession.SessionID()
	// 加入注册通道，意思是只要连接的人都加入Register通道
	beego.Debug("111")
	wss.H.Register <- c
	go c.WritePump() // 服务器端发送消息给客户端

	remoteIp := this.Ctx.Request.Header.Get("X-Forwarded-For-Agent")
	clientIp := ""
	if remoteIp != "" {
		clientIp = remoteIp
		c.Username = this.Ctx.Request.Header.Get("X-Forwarded-For-Agent")
	} else {
		clientIp = strings.Split(c.Ws.RemoteAddr().String(), ":")[0]
		c.Username = clientIp
	}
	// 添加全局websocket
	//wss.C = c.Ws
	//var ClientIp string
	//remoteIp := this.Ctx.Request.Header.Get("X-Forwarded-For-Agent")

	c.Username = this.Ctx.Request.Header.Get("X-Forwarded-For-Agent")

	if ssid != "" && uid != "" {
		c.ReadPump(webSocketId, gcmdbLoginStatus, clientIp)
	} else {
		c.ReadPump(this.Ctx.Input.Header("Sec-WebSocket-Key"), gcmdbLoginStatus, clientIp)
	}
	beego.Debug("is over")
	defer func() {
		//http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		if this.GetString("rstType") == "xml" {
			this.Data["xml"] = &r
			this.ServeXML()
		} else {
			this.Data["json"] = &r
			this.ServeJSON()
		}
		return
	}()
}
