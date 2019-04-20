/*************************************************************
     FileName: src->UploadAgentInfo->AgentUpload.go
         Date: 2018/4/6 上午11:25
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http://blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package wss

import (
	"encoding/json"
	//"gcmdb/models/hardware"
	"github.com/astaxie/beego"
)

func AgentUpload(rst DialData, clientIp string) {
	var st LinuxData
	agentinfo, agerr := json.Marshal(rst.Data)
	if agerr == nil {
		er := json.Unmarshal([]byte(string(agentinfo)), &st)
		if er != nil {
			return
		}
		beego.Debug("xxxxxxx",st.Eth1, st.Serialnumber)
		if st.Serialnumber != "" {
			//host, hoster := hardware.QueryHostSn(strings.TrimSpace(st.Serialnumber))
			//if hoster == nil {
			//	//beego.Debug("sn存在，开始写入新数据")
			//	host.Memory = st.Memory
			//	host.Disk = st.Disk
			//	host.Osrelease = strings.Replace(st.Osrelease, "\n", "", -1)
			//	r := regexp.MustCompile(`(localhost+)`)
			//	if r.MatchString(st.Fqdn) {
			//		host.Fqdn = strings.Replace(clientIp, ".", "-", -1)
			//	} else {
			//		host.Fqdn = st.Fqdn
			//	}
			//	host.Cpuarch = strings.Replace(st.Osarch, "\n", "", -1)
			//	host.NumCpus = strconv.Itoa(int(st.NumCpus))
			//	host.Eth2 = st.Eth2
			//	host.Eth3 = st.Eth3
			//	host.Eth4 = st.Eth4
			//
			//	host.Mac = st.Mac
			//
			//	host.Biosversion = st.Biosversion
			//	host.Biosreleasedate = st.Biosreleasedate
			//	host.CpuModel = st.CPUModel
			//	host.Manufacturer = st.Productname
			//	host.AgentVersion = st.AgentVersion
			//	host.Os = st.Os
			//	host.Serialnumber = st.Serialnumber
			//	host.AgentSurvival = true
			//
			//
			//
			//	hard, harder := hardware.HardwareQuerytitle(st.Manufacturer)
			//	if harder == nil {
			//		host.HardwareVendorId = hard.ID
			//	} else {
			//		hard.Title = st.Manufacturer
			//		ventor, _ := hardware.CreateVendor(&hard)
			//		host.HardwareVendorId = ventor.ID
			//	}
			//
			//	ventortype, err := hardware.HardwareVentorQuerytitle("服务器")
			//	if err == nil {
			//		host.EquipmentTypeId = ventortype.ID
			//	} else {
			//		ventortype.Title = "服务器"
			//		addventortype, _ := hardware.CreateType(&ventortype)
			//		host.EquipmentTypeId = addventortype.ID
			//	}
			//	_, deterr := host.UpdateHost()
			//
			//	if deterr != nil {
			//		beego.Error("更新主机失效")
			//	}
			//
			//}

		} else {

			beego.Debug("clientIpxxxxxxx",st.Eth1, st.Serialnumber,clientIp)
			//hostInfo, hoster := hardware.QueryHostIP(clientIp)
			//if hoster == nil {
			//	//hostDetail,_ := hardware.QueryHostSampleDetailById(uint(hostInfo.ID))
			//	hostInfo.Memory = st.Memory
			//	hostInfo.Disk = st.Disk
			//	hostInfo.Osrelease = strings.Replace(st.Osrelease, "\n", "", -1)
			//
			//	r := regexp.MustCompile(`(localhost+)`)
			//	if r.MatchString(st.Fqdn) {
			//		hostInfo.Fqdn = strings.Replace(clientIp, ".", "-", -1)
			//	} else {
			//		hostInfo.Fqdn = st.Fqdn
			//	}
			//
			//	hostInfo.Cpuarch = strings.Replace(st.Osarch, "\n", "", -1)
			//	hostInfo.NumCpus = strconv.Itoa(int(st.NumCpus))
			//	hostInfo.Eth2 = st.Eth2
			//	hostInfo.Eth3 = st.Eth3
			//	hostInfo.Eth4 = st.Eth4
			//	hostInfo.Mac = st.Mac
			//
			//	hostInfo.Biosversion = st.Biosversion
			//	hostInfo.Biosreleasedate = st.Biosreleasedate
			//	hostInfo.CpuModel = st.CPUModel
			//	hostInfo.Manufacturer = st.Productname
			//	hostInfo.AgentVersion = st.AgentVersion
			//	hostInfo.Os = st.Os
			//	hostInfo.Serialnumber = st.Serialnumber
			//
			//
			//	hard, harder := hardware.HardwareQuerytitle(st.Manufacturer)
			//	if harder == nil {
			//		hostInfo.HardwareVendorId = hard.ID
			//
			//	} else {
			//		hard.Title = st.Manufacturer
			//		ventor, _ := hardware.CreateVendor(&hard)
			//		hostInfo.HardwareVendorId = ventor.ID
			//	}
			//
			//	ventortype, err := hardware.HardwareVentorQuerytitle("服务器")
			//	if err == nil {
			//		hostInfo.EquipmentTypeId = ventortype.ID
			//	} else {
			//		ventortype.Title = "服务器"
			//		addventortype, _ := hardware.CreateType(&ventortype)
			//		hostInfo.HardwareVendorId = addventortype.ID
			//	}
			//	_, updateErr := hostInfo.UpdateHost()
			//	if updateErr != nil {
			//		beego.Error("更新主机数据失败", hostInfo.Eth1, err)
			//	}
			//
			//}
		}
	}
	//beego.Debug(clientIp, "上报硬件信息")
	return
}
