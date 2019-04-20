// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wss

import (
	//"gcmdb/common/utils"
	"gcmdb/controllers/websocket/agentInfo"
	//"gcmdb/models/hardware"
	"github.com/astaxie/beego"
)

// hub maintains the set of active connections and Broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	// 注册连接
	Connections map[*Connection]bool

	// Inbound messages from the connections.
	// 连接中的绑定消息
	Broadcast chan *DialData
	// Broadcast chan *RstData

	// Register requests from the connections.
	// 添加新连接
	Register chan *Connection

	// UnRegister requests from connections.
	// 删除连接
	UnRegister chan *Connection
}

var H = Hub{
	// 广播slice
	Broadcast: make(chan *DialData),
	// Broadcast: make(chan *SocketMessage),
	// 注册者slice
	Register: make(chan *Connection),
	// 未注册者sclie
	UnRegister: make(chan *Connection),
	// 连接map
	Connections: make(map[*Connection]bool),
}

func (h *Hub) Run() {
	for {
		select {
		// 注册者有数据，则插入连接map

		case c := <-h.Register:
			h.Connections[c] = true
			//c.WritePump()
			// 非注册者有数据，则删除连接map
		case c := <-h.UnRegister:

			if _, ok := h.Connections[c]; ok {

				_, oker := agentInfo.AgentMap.ReadMap(c.Username)

				if oker != nil {
					if c.Username != "GcmdbWebSocket" {
						beego.Info(c.Username, "用户下线了...")
						delete(h.Connections, c)
						//_ = c.Ws.Close()
					}

					//return //如果return,后续上线服务器则不会记录
				} else {
					//agent, _ := hardware.QueryHostIP(c.Username)
					//
					//if agent.Eth1 == c.Username {
					//	agent.AgentSurvival = false
					//	_, hoerr := agent.UpdateAgentSurvival()
					//	if hoerr != nil {
					//		beego.Warn("更新状态失败")
					//	}
					//}

					beego.Debug(ok)
					beego.Error(c.Username, "主机下线...")
					agentInfo.AgentMap.DeleteMap(c.Username, c.Username)

					//ca := utils.RedisCli.Get()
					//_, hmsetErr := ca.Do("hdel", "websocket_auth", c.Username)
					//if hmsetErr != nil {
					//	beego.Error(hmsetErr)
					//}
					//defer func() {
					//	_ = ca.Close()
					//}()
					delete(h.Connections, c)
				}

			}
			// 广播有数据
		case m := <-h.Broadcast:
			// 递归所有广播连接
			for c := range h.Connections {
				var send_msg []byte
				//remip := strings.Split(c.Ws.RemoteAddr().String(), ":")[0]
				beego.Debug(m.Active)
				if c.Auth {
					switch m.Active {
					case "login":

					case "update":
						send_msg = []byte("upload: " + m.Msg)
						select {
						// 发送数据给连接
						case c.Send <- send_msg:
							// 关闭连接
						default:

						}

					case "push":
						// var rst PushData
						beego.Debug("websocket发布代码")
						//go PushAgent(m, c, send_msg)

					//case "callBackPush":
					//	go CallBackPush(m, c, send_msg)
					//
					//case "CallBackTask":
					//	go CallBackTask(m, c, send_msg)

					case "callBackMsg":
						go CallBackMsg(m, c, send_msg)

					case "callBackAgentVersion":
						go CallBackMsg(m, c, send_msg)

					case "callBackCodeVersion":
						go CallBackMsg(m, c, send_msg)

					case "BashShell":
						beego.Warning("shell执行")
					//	go PushAgentShell(m, c, send_msg)
					//
					//case "AgentVersion":
					//	go PushAgent(m, c, send_msg)
					//
					//case "CodeVersion":
					//	go PushAgent(m, c, send_msg)

					case "callBackBashShell":
						go CallBackShellMsg(m, c, send_msg)
					case "CheckHttp":
						beego.Info("开始进行http检测")
						go CheckHttpUrl(m, c, send_msg)
					case "callBackCheckHttp":
						beego.Info("回调检测数据")
						go CallBackCheckHttp(m, c, send_msg)
					case "CheckDns":
						beego.Info("开始进行http检测")
						go CheckTools(m, c, send_msg)
					case "CheckPing":
						beego.Info("开始进行ping检测")
						go CheckTools(m, c, send_msg)
					case "callBackCheckDns":
						beego.Info("dns回调检测数据")
						beego.Debug(m.Data)
						go CallBackCheckActive(m, c, send_msg)
					case "callBackCheckPing":
						beego.Info("ping回调检测数据")
						beego.Debug(m.Data)
						go CallBackCheckActive(m, c, send_msg)

					default:
						send_msg = []byte("默认: " + "接收所有操作")
						select {
						// 发送数据给连接
						case c.Send <- send_msg:
							// 关闭连接
						default:
							close(c.Send)
							delete(h.Connections, c)
						}

					}

				}
			}
		}
	}
}
