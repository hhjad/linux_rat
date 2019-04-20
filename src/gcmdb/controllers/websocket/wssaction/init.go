/*************************************************************
     FileName: gcmdb->websocket->init.go
         Date: 2018/3/15 下午2:30
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http://blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package wssaction

import (
	"gcmdb/controllers"
)

type WebSocketController struct {
	controllers.BaseController
}


type AgentData []AgentDict

type AgentDict struct {
	ID          uint   `json:"ID" xml:"id"`
	Fqdn     string `json:"fqdn" xml:"fqdn"`
	Eth1      string   `json:"eth1" xml:"eth1"`
}

type AgentActive map[string]string

type AgentResult struct {
	Down AgentData `json:"down"`
	Up  interface{} `json:"up"`
}