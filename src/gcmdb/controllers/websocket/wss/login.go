/*************************************************************
     FileName: src->Auth->login.go
         Date: 2018/4/7 下午9:32
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http://blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package wss

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	//"gcmdb/common/utils"
	"gcmdb/controllers/websocket/agentInfo"
	//"gcmdb/models/hardware"
	"github.com/astaxie/beego"
)

func LoginWebSocket(rst DialData, c *Connection, clientIp, key string) {
	if rst.Manager {
		tokenMd5 := (beego.AppConfig.String("websocket::managerToken") + "voilet@qq.com")
		h := md5.New()
		h.Write([]byte(tokenMd5))
		md5str2 := fmt.Sprintf("%x", h.Sum(nil))
		if rst.Token == md5str2 {
			c.Manager = true
			c.Auth = true
			c.Username = "GcmdbWebSocket"
		}
		return
	} else {
		//timeStr := time.Now().Format("2006-01-02 15")
		tokenMd5 := (beego.AppConfig.String("websocket::authToken") + "voilet@qq.com")
		h := md5.New()
		h.Write([]byte(tokenMd5))
		md5str2 := fmt.Sprintf("%x", h.Sum(nil))
		//cl := utils.RedisCli.Get()
		//defer func() {
		//	_ = cl.Close()
		//	return
		//}()
		if rst.Token == md5str2 {
			if rst.Agent {
				var agent AgentLogin
				var ip string
				//var serialNumber string
				result, agerr := json.Marshal(rst.Data)
				if agerr != nil {
					beego.Error("格式化数据失败", clientIp)
					return
				}
				er := json.Unmarshal([]byte(string(result)), &agent)

				if er != nil {
					beego.Error("分解登录上报数据失败", clientIp)
					return
				}
				beego.Debug(agent.Eth1)
				beego.Debug(clientIp)
				if agent.Eth1 != "" {
					ip = agent.Eth1
					//serialNumber = strings.TrimSpace(agent.Serialnumber)
				} else {
					ip = clientIp
					//serialNumber = strings.TrimSpace(agent.Serialnumber)
				}
				beego.Debug(ip)
				beego.Debug(agent.AgentVersion)
				if agent.Eth1 != "" {
					c.Username = agent.Eth1
					_, agentErr := agentInfo.AgentMap.ReadMap(ip)
					if agentErr != nil {
						agentInfo.AgentMap.WriteMap(ip, ip)
					}
					//_, hmsetErr := cl.Do("hset", "websocket_auth", ip, ip)
					//if hmsetErr != nil {
					//	beego.Debug(hmsetErr)
					//}
				} else {
					c.Username = ip
					//_, hmsetErr := cl.Do("hset", "websocket_auth", ip, ip)
					//if hmsetErr != nil {
					//	beego.Debug(hmsetErr)
					//}

					_, agentErr := agentInfo.AgentMap.ReadMap(ip)
					if agentErr != nil {
						agentInfo.AgentMap.WriteMap(ip, ip)
					}
				}

				//物理机
				if agent.SerialnumberStatus {
				//	ActiveAgent, Sner := hardware.QueryHostSn(serialNumber)
				//	ActiveAgent.Eth1 = agent.Eth1
				//	ActiveAgent.Eth2 = agent.Eth2
				//
				//	if agent.Eth1 != "" {
				//		ActiveAgent.Fqdn = strings.Replace(agent.Eth1, ".", "-", -1)
				//	} else {
				//		ActiveAgent.Fqdn = strings.Replace(agent.Eth2, ".", "-", -1)
				//	}
				//
				//	//没有查询到序列号,添加主机
				//	if Sner != nil {
				//		ActiveAgent.Status = 1
				//		ActiveAgent.Serialnumber = serialNumber
				//		ActiveAgent.AgentVersion = agent.AgentVersion
				//		ActiveAgent.AgentSurvival = true
				//		_, saer := hardware.CreateHost(&ActiveAgent)
				//		if saer != nil {
				//			beego.Error("存入数据失败", clientIp, "SN号: ", rst.Msg, saer)
				//		} else {
				//			beego.Debug("数据库中无: ", clientIp, "写入新数据到数据库")
				//		}
				//		return
				//
				//	}
				//
				//	ActiveAgent.AgentSurvival = true
				//	ActiveAgent.Mac = agent.Mac
				//	_, uper := ActiveAgent.UpdateAgentSurvival()
				//
				//	if uper != nil {
				//		beego.Error("更新数据失败:", clientIp, uper)
				//	}
				//
				//	_, uperhost := ActiveAgent.UpdateHost()
				//	if uperhost != nil {
				//		beego.Error("更新数据失败:", clientIp, uperhost)
				//	}
				//
				//	ss, hmsetErr := cl.Do("hlen", "websocket_auth")
				//	if hmsetErr != nil {
				//		beego.Debug(hmsetErr)
				//	}
				//	/*beego.Debug(hmsetErr)
				//	beego.Debug(ss)*/
				//	beego.Info("新上线Agent: ", ip, "当前redis在线用户数: ", ss)
				//	beego.Info("新上线Agent: ", ip, "当前mem在线用户数: ", agentInfo.AgentMap.LenMap())
					beego.Info("新上线Agent: ", ip, "当前redis在线用户数: ")
					beego.Info("新上线Agent: ", ip, "当前mem在线用户数: ", agentInfo.AgentMap.LenMap())
					c.Auth = true
					return

				} else {
					//虚机
					//ActiveAgent, iper := hardware.QueryHostIPInfo(ip)
					//ActiveAgent.Eth1 = ip
					//ActiveAgent.Fqdn = strings.Replace(ip, ".", "-", -1)
					//ActiveAgent.Status = 1
					//ActiveAgent.AgentVersion = agent.AgentVersion
					//ActiveAgent.AgentSurvival = true
					//ActiveAgent.Serialnumber = agent.Serialnumber
					//ActiveAgent.Mac = agent.Mac
					//if iper != nil {
					//	_, saer := hardware.CreateHost(&ActiveAgent)
					//	if saer != nil {
					//		beego.Error("存入数据失败", clientIp, "SN号: ", rst.Msg, saer)
					//	} else {
					//		beego.Debug("数据库中无: ", clientIp, "写入新数据到数据库", agent.AgentVersion, agent.Serialnumber)
					//	}
					//} else {
					//	_, uper := ActiveAgent.UpdateHost()
					//	if uper != nil {
					//		beego.Error("更新数据失败:", clientIp, uper)
					//	}
					//}
					//ss, hmsetErr := cl.Do("hlen", "websocket_auth")
					//if hmsetErr != nil {
					//	beego.Debug(hmsetErr)
					//}
					//beego.Info("新上线Agent: ", ip, "当前redis在线用户数: ", ss)
					beego.Info("新上线Agent: ", ip, "当前redis在线用户数: ")
					beego.Info("新上线Agent: ", ip, "当前mem在线用户数: ", agentInfo.AgentMap.LenMap())
					c.Auth = true
					return
				}
			}
		} else {
			c.Username = key
			c.Auth = true
		}
	}
	return
}

func (login AgentLogin) AgentLoginAuth() AgentLogin {
	rst := login
	return rst
}
