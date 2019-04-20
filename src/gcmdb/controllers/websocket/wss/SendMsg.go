/*************************************************************
     FileName: src->wss->SendMsg.go
         Date: 2018/4/18 下午10:03
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http://blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package wss

import "github.com/astaxie/beego"

func SendMsgManager(c *Connection, rst DialData) {
	defer func() {
		return
	}()
	beego.Debug("开始执行shell")
	beego.Debug(rst.Secret)
	if !rst.Secret {
		if c.Manager && c.Auth == true {
				H.Broadcast <- &rst
		}
	} else {
		beego.Debug(c.Manager)
		beego.Debug(c.Auth)
		if c.Manager && c.Auth == true {
			beego.Error("私聊2")
			var s = &rst
			if len(s.ToUserList) > 0 {
				for _,v := range s.ToUserList{
					beego.Error("私聊to", v)
					s.ToUser = v
					H.Broadcast <- s
				}
				return
			}else{
				beego.Info(s.ToUser)
				beego.Info(s.CallBackUser)
				beego.Warning(s)
				H.Broadcast <- s
				return
			}
		}
	}

}
