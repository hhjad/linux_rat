///*************************************************************
//     FileName: src->wss->safetyUpload.go
//         Date: 2018/4/13 下午9:58
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
//	"errors"
//	"gcmdb/common/utils"
//	"gcmdb/models/safety"
//	"github.com/astaxie/beego"
//	"github.com/gomodule/redigo/redis"
//	"reflect"
//	"strings"
//)
//
//func Contain(obj interface{}, target interface{}) (bool, error) {
//	targetValue := reflect.ValueOf(target)
//	switch reflect.TypeOf(target).Kind() {
//	case reflect.Slice, reflect.Array:
//		for i := 0; i < targetValue.Len(); i++ {
//			if targetValue.Index(i).Interface() == obj {
//				return true, nil
//			}
//		}
//	case reflect.Map:
//		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
//			return true, nil
//		}
//	}
//
//	return false, errors.New("not in array")
//}
//
//func SafetyUpload(rst DialData, clientIp string) {
//	var st Safety
//	c := utils.RedisCli.Get()
//
//	/*
//	if err := mapstructure.Decode(rst.Data, &st); err != nil {
//		beego.Error("数据结构异常", clientIp)
//		return
//	}
//	*/
//
//	agentinfo, agerr := json.Marshal(rst.Data)
//	if agerr != nil {
//		beego.Error("格式化数据失败", clientIp)
//		return
//	}
//	er := json.Unmarshal([]byte(string(agentinfo)), &st)
//
//	if er != nil {
//		beego.Error("分解安全上报数据失败", clientIp)
//		return
//	}
//
//	_, hoster := safety.QueryIp(clientIp)
//
//	if hoster != nil {
//
//		safelog := new(safety.Safety)
//		tcpinfo, tcper := json.Marshal(st.Tcp)
//		if tcper == nil {
//			safelog.TcpPort = string(tcpinfo)
//		}
//
//		udpinfo, udper := json.Marshal(st.Udp)
//		if udper == nil {
//			safelog.UdpPort = string(udpinfo)
//		}
//
//		safelog.Process = strings.Join(st.Process, ",")
//		safelog.ProcessCount = st.ProcessCount
//		//beego.Debug(strings.Join(st.TcpLog, ","))
//		//beego.Debug(strings.Join(st.UdpLog, ","))
//		safelog.UdpPort = ""
//		safelog.Ip = clientIp
//		safelog.TcpCount = st.TcpCount
//		safelog.UdpCount = st.UdpCount
//		safelog.TcpLog = strings.Join(st.TcpLog, ",")
//		safelog.UdpPort = strings.Join(st.UdpLog, ",")
//		/**
//		 * @Description: 增加tcp端口
//		 * @author: 苦咖啡
//		 * @Date:   2018/11/6 1:56 PM
//		*/
//		_, delTcpEr := redis.Bool(c.Do("del", clientIp+"_tcp"))
//		if delTcpEr != nil {
//			beego.Error("删除：", clientIp, "tcp redis数据不存在")
//		}
//		for _, v := range st.Tcp {
//			//beego.Debug(clientIp, "udp", v.Port)
//			if v.Name != "local" {
//				_, v2Err := redis.Bool(c.Do("hset", clientIp+"_tcp", v.Port, v.Port))
//				if v2Err != nil {
//					beego.Debug(v2Err)
//				}
//			}
//
//		}
//		/**
//		 * @Description: 增加udp端口
//		 * @author: 苦咖啡
//		 * @Date:   2018/11/6 2:09 PM
//		*/
//		_, delUdpEr := redis.Bool(c.Do("del", clientIp+"_udp"))
//		if delUdpEr != nil {
//			beego.Error("删除：", clientIp, "udp redis数据不存在")
//		}
//		for _, u := range st.Udp {
//			//beego.Debug(clientIp, "udp", u.Port)
//			if u.Name != "local" {
//				_, v2Err := redis.Bool(c.Do("hset", clientIp+"_udp", u.Port, u.Port))
//				if v2Err != nil {
//					beego.Debug(v2Err)
//				}
//			}
//		}
//		/**
//		 * @Description: 增加进程
//		 * @author: 苦咖啡
//		 * @Date:   2018/11/6 2:10 PM
//		*/
//		_, delPidEr := redis.Bool(c.Do("del", clientIp+"_pid"))
//		if delPidEr != nil {
//			beego.Error("删除：", clientIp, "pid redis数据不存在")
//		}
//		for _, p := range st.Process {
//
//			_, v2Err := redis.Bool(c.Do("hset", clientIp+"_pid", p, p))
//			if v2Err != nil {
//				beego.Debug(v2Err)
//			}
//		}
//
//		_, err := safety.Create(safelog)
//		if err != nil {
//			beego.Error("保存服务器服务端口白名单失败 -> ", clientIp)
//		}
//		return
//	}
//	/**
//	 * @Description: 检测ip是否在过滤组
//	 * @author: 苦咖啡
//	 * @Date:   2018/11/7 12:06 PM
//	*/
//	_, safeBypassErr := redis.String(c.Do("hget", "safe_by_pass", clientIp))
//	if safeBypassErr != nil {
//		/**
//	 * @Description: 检测tcp端口
//	 * @author: 苦咖啡
//	 * @Date:   2018/11/5 3:35 PM
//	*/
//
//		oldtcpLen, tcpEr := redis.Int(c.Do("hlen", clientIp+"_tcp"))
//		if tcpEr != nil {
//			beego.Debug("tcp len: ", tcpEr)
//		}
//		newTcpLen := len(st.Tcp)
//		/*beego.Debug(clientIp, "tcp端口数量: ", oldtcpLen)
//		beego.Debug(clientIp, "上报端口数量: ", newTcpLen)*/
//
//		if newTcpLen < oldtcpLen {
//			oldTcp, oldEr := redis.Strings(c.Do("hkeys", clientIp+"_tcp"))
//			if oldEr != nil {
//				beego.Debug(oldEr)
//			}
//			for _, v := range oldTcp {
//				for _, val := range st.Tcp {
//					_, v2Er := redis.Bool(c.Do("hset", clientIp+"_tcp_check", val.Port, val.Port))
//					if v2Er != nil {
//						beego.Debug(clientIp, "检测Tcp端口异常: ", val.Port, val.Name, v2Er.Error())
//					}
//				}
//				_, v2Er := redis.String(c.Do("hget", clientIp+"_tcp_check", v))
//				if v2Er != nil {
//					beego.Debug(clientIp, "tcp异常端口: ", v)
//				}
//			}
//			_, delPidEr := redis.Bool(c.Do("del", clientIp+"_tcp_check"))
//			if delPidEr != nil {
//				beego.Error("删除：", clientIp, "对比数据")
//			}
//		} else {
//			for _, v := range st.Tcp {
//				if v.Name != "local" {
//					_, v2Er := redis.String(c.Do("hget", clientIp+"_tcp", v.Port))
//					if v2Er != nil {
//						beego.Debug(clientIp, "新增Tcp端口: ", v.Port, v.Name)
//					}
//				}
//			}
//		}
//
//		/**
//		 * @Description: 检测udp端
//		 * @author: 苦咖啡
//		 * @Date:   2018/11/5 3:35 PM
//		*/
//		oldUdpLen, udpEr := redis.Int(c.Do("hlen", clientIp+"_udp"))
//		if udpEr != nil {
//			beego.Debug("tcp len: ", tcpEr)
//		}
//		newUdpLen := len(st.Udp)
//		/*beego.Debug(clientIp, "udp端口数量: ", oldUdpLen)
//		beego.Debug(clientIp, "udp上报端口数量: ", newTcpLen)*/
//
//		if newUdpLen < oldUdpLen {
//			oldUdp, oldEr := redis.Strings(c.Do("hkeys", clientIp+"_udp"))
//			if oldEr != nil {
//				beego.Debug(oldEr)
//			}
//			//abnormal := ""
//			for _, v := range oldUdp {
//				for _, val := range st.Udp {
//					_, v2Er := redis.Bool(c.Do("hset", clientIp+"_udp_check", val.Port, val.Port))
//					if v2Er != nil {
//						beego.Debug(clientIp, "检测Udp端口异常: ", val.Port, val.Name, v2Er)
//					}
//				}
//				_, v2Er := redis.String(c.Do("hget", clientIp+"_udp_check", v))
//				if v2Er != nil {
//					beego.Debug(clientIp, "Udp异常端口: ", v)
//				}
//			}
//			_, delPidEr := redis.Bool(c.Do("del", clientIp+"_udp_check"))
//			if delPidEr != nil {
//				beego.Error("删除：", clientIp, "udp_对比数据")
//			}
//			//beego.Debug(abnormal)
//		} else {
//			for _, v := range st.Udp {
//				if v.Name != "local" {
//					_, v2Er := redis.String(c.Do("hget", clientIp+"_udp", v.Port))
//					if v2Er != nil {
//						beego.Debug(clientIp, "新增Udp端口: ", v.Port, v.Name)
//					}
//				}
//			}
//		}
//
//		/**
//		 * @Description: 检测进程
//		 * @author: 苦咖啡
//		 * @Date:   2018/11/6 2:03 PM
//		*/
//		for _, v := range st.Process {
//			_, v2Er := redis.String(c.Do("hget", clientIp+"_pid", v))
//			if v2Er != nil {
//				//beego.Debug(clientIp, "新增进程: ", v)
//			}
//		}
//	}
//
//	defer c.Close()
//
//	return
//
//}
