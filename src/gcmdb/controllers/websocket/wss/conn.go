// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wss

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

const (
	// 对方写入会话等待时间
	// Time allowed to write a message to the peer.
	writeWait = 86400 * time.Second

	// 对方读取下次消息等待时间
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// 对方ping周期
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// 对方最大写入字节数
	// Maximum message size allowed from peer.
	maxMessageSize = 102400
)

var C *Connection
//var C *websocket.Conn

// 服务器配置信息
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024000,
	WriteBufferSize: 1024000,
}

// connection 是websocket的conntion和hub的中间人
// connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// websocket的连接
	Ws *websocket.Conn
	// Buffered channel of outbound messages.
	// 出站消息缓存通道
	Send chan []byte

	// 验证状态，普通用户，用于上报数据和自动更新
	Auth bool

	// Manager
	Manager bool
	// 是否管理员，添加帐号，主动推送信息，执行命令
	Agent bool
	// 所属用户
	Username string
}

// 读取connection中的数据导入到hub中，实则发广播消息
// 服务器读取的所有客户端的发来的消息
// readPump pumps messages from the websocket connection to the hub.
func (c *Connection) ReadPump(key string, logStatus bool, remoteIp string) {
	defer func() {
		H.UnRegister <- c
		_ = c.Ws.Close()
	}()
	c.Ws.SetReadLimit(maxMessageSize)
	_ = c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {

		_, message, err := c.Ws.ReadMessage()

		if err != nil {
			break
		}
		/*
			获取并处理message数据
			DialData
		*/
		if logStatus {
			c.Auth = true
			c.Username = key
		}
		//beego.Debug(c.Ws.RemoteAddr().String())
		//beego.Debug(remoteIp)
		clientIp := ""
		if remoteIp != "" {
			clientIp = remoteIp
		} else {
			clientIp = strings.Split(c.Ws.RemoteAddr().String(), ":")[0]
		}

		var rst DialData
		// var rst RstData

		resuerr := json.Unmarshal([]byte(string(message)), &rst)
		if resuerr != nil {
			if string(message) != "setPing" {
				beego.Info("数据格式有误", resuerr, remoteIp)
			}
		} else {
			// 获取字段格式
			switch rst.Active {

			case "login":
				//beego.Debug(c.Ws.)
				//c.Ws.id
				go LoginWebSocket(rst, c, clientIp, key)

			case "uploadAgentInfo":
				AgentUpload(rst, clientIp)
			case "uploadAgentSafe":
				//SafetyUpload(rst, clientIp)

			case "AddSystemUser":
				go SendMsgManager(c, rst)

			case "push":
				go SendMsgManager(c, rst)

			case "BashShell":
				go SendMsgManager(c, rst)

			case "callBackPush":
				//rst.Msg = rst.Msg
				// 需加入日志记录
				rst.Active = "callBackPush"
				if c.Auth == true {
					H.Broadcast <- &rst
				}

			case "AgentVersion":
				go SendMsgManager(c, rst)

			case "callBackAgentVersion":
				if c.Auth == true {
					H.Broadcast <- &rst
				}
			case "callBackCheckHttp":
				if c.Auth == true {
					H.Broadcast <- &rst
				}

			case "callBackCheckDns":
				if c.Auth == true {
					H.Broadcast <- &rst
				}
			case "callBackCheckPing":
				if c.Auth == true {
					H.Broadcast <- &rst
				}
			case "CodeVersion":
				if c.Manager && c.Auth == true {
					H.Broadcast <- &rst
				}
			case "callBackCodeVersion":
				if c.Auth == true {
					H.Broadcast <- &rst
				}

			case "callBackBashShell":
				beego.Debug("执行shell回调数据")
				if c.Auth == true {
					H.Broadcast <- &rst
				}

			default:
				//c.Ws.Close()
			}
		}

	}
}

// 给消息，指定消息类型和荷载
// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	_ = c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	//c.Ws.Close()
	return c.Ws.WriteMessage(mt, payload)
}

func (c *Connection) writejson(mt int, payload []byte) error {
	_ = c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	//c.Ws.Close()
	return c.Ws.WriteJSON(payload)
}

// 从hub到connection写数据
// 服务器端发送消息给客户端
// writePump pumps messages from the hub to the websocket connection.
func (c *Connection) WritePump() {
	// 定时执行
	beego.Debug("测试")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		//_ = c.Ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})

				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
