///*************************************************************
//     FileName: gcmdb->wss->hupPush.go
//         Date: 2018/4/8 下午11:05
//       Author: 苦咖啡
//        Email: voilet@qq.com
//         blog: http://blog.kukafei520.net
//      Version: 0.0.1
//      History:
//**************************************************************/
//
package wss
//
//import (
//	"encoding/json"
//	"gcmdb/controllers/websocket/agentInfo"
//	//"gcmdb/models/carving"
//	"github.com/astaxie/beego"
//	"time"
//)
//
//func PushAgent(m *DialData, c *Connection, send_msg []byte) {
//	beego.Info(m.Data)
//	rst, err := json.Marshal(m)
//	if err != nil {
//		beego.Error("数据格式错误")
//	} else {
//		if m.ToUser == c.Username {
//			send_msg = []byte(rst)
//			select {
//			// 发送数据给连接
//			case c.Send <- send_msg:
//				// 关闭连接
//			default:
//
//			}
//		}
//	}
//}
//
//func PushAgentShell(m *DialData, c *Connection, send_msg []byte) {
//	var rstMsg DialData
//	rstMsg.Data = m.Data
//	rstMsg.Code = m.Code
//	rstMsg.Msg = m.Msg
//	//rstMsg.Time = time.Now().Format("2006-01-02 15:04:05")
//	rstMsg.Active = m.Active
//	rest, err := json.Marshal(rstMsg)
//
//	if err != nil {
//		beego.Error("数据格式错误", c.Username)
//	} else {
//		send_msg = []byte(rest)
//		_, agentErr := agentInfo.AgentMap.ReadMap(c.Username)
//		if agentErr != nil {
//			return
//		}
//		if m.Secret {
//			beego.Debug(m.CallBackUser)
//			beego.Debug(c.Username)
//			if m.ToUser == c.Username {
//				beego.Warning(m.ToUser)
//				beego.Warning(c.Username)
//				beego.Error(m.Active)
//				beego.Error(string(send_msg))
//				select {
//				// 发送数据给连接
//				case c.Send <- send_msg:
//					// 关闭连接
//				default:
//
//				}
//			}
//
//		} else {
//			beego.Warning(rest)
//			//if m.ToUser == c.Username{
//			beego.Info(m.Secret)
//			beego.Error(m.CallBackUser)
//			beego.Warning(m.ToUser)
//			select {
//			// 发送数据给连接
//			case c.Send <- send_msg:
//				// 关闭连接
//			default:
//
//			}
//			//}
//		}
//	}
//}
//
//func CallBackPush(m *DialData, c *Connection, send_msg []byte) {
//	var rstMsg DialDataRst
//	rstMsg.Data = m.Data
//	rstMsg.Code = m.Code
//	rstMsg.Msg = m.Msg
//	rstMsg.Active = m.Active
//	rstMsg.Time = time.Now().Format("2006-01-02 15:04:05")
//	/**
//	 * @Description: 日志入库
//	 * @author: 苦咖啡
//	 * @Date:   2018/11/13 1:21 PM
//	*/
//	var pushlog TaskMsg
//	pushrst, er := json.Marshal(m.Data)
//	if er != nil {
//		beego.Error("转数据格式异常")
//	} else {
//		_ = json.Unmarshal(pushrst, &pushlog)
//		beego.Debug(pushlog.Version)
//		beego.Debug(pushlog.Reset)
//		beego.Debug(pushlog.Title)
//		beego.Debug(pushlog.Ip)
//		//s, e := carving.ReleaceByLogIpId(pushlog.ReleaseId, pushlog.Ip)
//		//if e != nil {
//		//	beego.Debug("无此ip日志入库数据", pushlog.Ip)
//		//} else {
//		//	s.Reset = pushlog.Reset
//		//	if !s.Reset {
//		//		s.PostStatus = false
//		//		if pushlog.Status != 200{
//		//			s.PreStatus = false
//		//		}else{
//		//			s.PreStatus = true
//		//		}
//		//	}else{
//		//		s.PreStatus = false
//		//		if pushlog.Status != 200{
//		//			s.PostStatus = false
//		//		}else{
//		//			s.PostStatus = true
//		//		}
//		//	}
//		//	s.LogMsg = pushlog.Msg
//		//	_, relogEr := s.UpdateReleaceByLogId()
//		//	if relogEr != nil {
//		//		beego.Debug("更新日志失败")
//		//	}
//			beego.Debug("日志入库完成")
//		}
//
//	}
//
//	//rest, err := json.Marshal(rstMsg)
//	_, err := json.Marshal(rstMsg)
//	if err != nil {
//		beego.Error("数据格式错误", c.Username)
//	} else {
//		if m.CallBackUser == c.Username {
//			send_msg = []byte(rest)
//			select {
//			// 发送数据给连接
//			case c.Send <- send_msg:
//				// 关闭连接
//			default:
//
//			}
//		}
//
//	}
//}
//
//func CallBackTask(m *DialData, c *Connection, send_msg []byte) {
//	beego.Error(m.Data)
//	var rstMsg DialDataRst
//	rstMsg.Data = m.Data
//	rstMsg.Code = m.Code
//	rstMsg.Msg = m.Msg
//	rstMsg.Active = "CallBackTask"
//	rstMsg.Time = time.Now().Format("2006-01-02 15:04:05")
//	rest, err := json.Marshal(rstMsg)
//	if err != nil {
//		beego.Error("数据格式错误", c.Username)
//	} else {
//		if m.CallBackUser == c.Username {
//			send_msg = []byte(rest)
//			select {
//			// 发送数据给连接
//			case c.Send <- send_msg:
//				// 关闭连接
//			default:
//
//			}
//		}
//
//	}
//}
