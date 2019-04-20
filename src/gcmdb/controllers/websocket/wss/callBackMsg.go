/*************************************************************
     FileName: gcmdb->wss->callBackMsg.go
         Date: 2018/4/8 下午11:15
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http://blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package wss

import (
	"gcmdb/controllers/websocket/agentInfo"
	"github.com/astaxie/beego"
	"time"
	"encoding/json"
)

func CallBackMsg(m *DialData, c *Connection, send_msg []byte) {

	var rstMsg DialDataRst
	rstMsg.Data = m.Data
	rstMsg.Code = m.Code
	rstMsg.Msg = m.Msg
	rstMsg.Time = time.Now().Format("2006-01-02 15:04:05")
	beego.Error("websocket set callBaskMsg")
	beego.Info(m.CallBackUser)
	beego.Info(c.Username)
	beego.Info(m.Data)
	rest, err := json.Marshal(rstMsg)
	if err != nil {
		beego.Error("数据格式错误", c.Username)
	} else {
		if m.CallBackUser == c.Username {
			send_msg = []byte(rest)
			select {
			// 发送数据给连接
			case c.Send <- send_msg:
				// 关闭连接
			default:

			}
		}

	}
}

func CallBackShellMsg(m *DialData, c *Connection, send_msg []byte) {
	var rstMsg DialDataRst
	rstMsg.Data = m.Data
	rstMsg.Code = m.Code
	rstMsg.Msg = m.Msg
	rstMsg.Time = time.Now().Format("2006-01-02 15:04:05")
	beego.Error("websocket set callBaskpush")
	beego.Info(m.CallBackUser)
	beego.Info(c.Username)
	rest, err := json.Marshal(rstMsg)
	if err != nil {
		beego.Error("数据格式错误", c.Username)
	} else {
		if m.CallBackUser == c.Username {
			send_msg = []byte(rest)
			select {
			// 发送数据给连接
			case c.Send <- send_msg:
				// 关闭连接
			default:

			}
		}

	}
}

func CheckHttpUrl(m *DialData, c *Connection, send_msg []byte) {
	beego.Debug("初始化Http检测数据")
	beego.Debug(m)

	rest, err := json.Marshal(m)
	if err != nil {
		beego.Error("数据格式错误", c.Username)
	} else {
		_, oker := agentInfo.AgentMap.ReadMap(c.Username)
		if oker != nil {
			return
		}
		send_msg = []byte(rest)
		select {
		// 发送数据给连接
		case c.Send <- send_msg:
			// 关闭连接
		default:

		}

	}
}

func CheckTools(m *DialData, c *Connection, send_msg []byte) {
	beego.Debug("初始化Check方法")

	rest, err := json.Marshal(m)
	if err != nil {
		beego.Error("数据格式错误", c.Username)
	} else {
		_, oker := agentInfo.AgentMap.ReadMap(c.Username)
		if oker != nil {
			return
		}
		send_msg = []byte(rest)
		select {
		// 发送数据给连接
		case c.Send <- send_msg:
			// 关闭连接
		default:
			return
		}
	}
}

func CallBackCheckHttp(m *DialData, c *Connection, send_msg []byte) {

	var rstMsg DialDataRst
	rstMsg.Data = m.Data
	rstMsg.Code = m.Code
	rstMsg.Msg = m.Msg
	rstMsg.Active = "callBackCheckHttp"
	rstMsg.Time = time.Now().Format("2006-01-02 15:04:05")
	beego.Error("websocket set callBackCheckHttp")
	beego.Info(m.CallBackUser)
	beego.Info(c.Username)
	beego.Info(m.Data)
	rest, err := json.Marshal(rstMsg)
	if err != nil {
		beego.Error("数据格式错误", c.Username)
	} else {
		if m.CallBackUser == c.Username {
			beego.Info(rstMsg.Data)
			send_msg = []byte(rest)
			select {
			// 发送数据给连接
			case c.Send <- send_msg:
				// 关闭连接
			default:

			}
		}

	}
}

func CallBackCheckActive(m *DialData, c *Connection, send_msg []byte) {

	var rstMsg DialDataRst
	rstMsg.Data = m.Data
	rstMsg.Code = m.Code
	rstMsg.Msg = m.Msg
	rstMsg.Active = m.Active
	rstMsg.Time = time.Now().Format("2006-01-02 15:04:05")
	rest, err := json.Marshal(rstMsg)
	if err != nil {
		beego.Error("数据格式错误", c.Username)
	} else {
		if m.CallBackUser == c.Username {
			beego.Info(rstMsg.Data)
			send_msg = []byte(rest)
			select {
			// 发送数据给连接
			case c.Send <- send_msg:
				// 关闭连接
			default:

			}
		}

	}
}
