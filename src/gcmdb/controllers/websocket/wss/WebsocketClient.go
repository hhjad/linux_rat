/*************************************************************
     FileName: gcmdb->controllers->websocket.go
         Date: 2018/3/26 上午10:17
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http://blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package wss

import (
	"net/url"
	"github.com/astaxie/beego"
)

func WsClient() {

	u := url.URL{Scheme: "ws", Host: beego.AppConfig.String("httpaddr") + ":" + beego.AppConfig.String("httpport"), Path: "/wss"}
	beego.Info("connecting to %s", u.String())
	/*var rst DialData
	// var rst RstData
	_ = json.Unmarshal([]byte(string(message)), &rst)
	H.Broadcast <- &rst*/
	/*//连接服务器
	for {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			wss.SleepRandomDurationCount()
			beego.Info("连接websocket失败，重试...")
		} else {
			C = c
			var rst DialData
			rst.Token = beego.AppConfig.String("managerToken")
			rst.Active = "login"
			rst.Msg = "登录验证"
			rst.Manager = true
			rs, _ := json.Marshal(rst)
			auth_err := C.WriteMessage(websocket.TextMessage, []byte(rs))
			if auth_err != nil {
				beego.Info("write------------:", auth_err)
			}
			//break
			for{
				time.Sleep(time.Duration(10) * time.Second)
				if c.LocalAddr().String() == ""{
					beego.Error("连接websocket service 异常")
					break
				}
				C.WriteMessage(websocket.TextMessage, []byte("setPing"))
			}

		}
	}*/

}