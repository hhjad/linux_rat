package wss

type LinuxData struct {
	Manufacturer    string `json:"manufacturer"`  // 厂商
	Fqdn            string `json:"fqdn"`          // 主机名
	Kernel          string `json:"kernel"`        // 所属内核
	Kernelrelease   string `json:"kernelrelease"` // 内核版本
	NumCpus         int64  `json:"num_cpus"`      // CPU核数
	Os              string `json:"os"`            // 系统类型(centos ubuntu debian windows)
	Osarch          string `json:"osarch"`        // 系统平台 X86_64 i386
	Osrelease       string `json:"osrelease"`     // 系统版本
	Memory          string `json:"memory"`
	Disk            string `json:"disk"`
	Descr           string `json:"description"`
	Product         string `json:"product"`
	Serial          string `json:"serial"`
	Vendor          string `json:"vendor"`
	CPUModel        string `json:"cpu_model"`    // cpu型号	"DMI type 4"
	Cpuarch         string `json:"cpuarch"`      // cpu平台 x86_64 i386 "DMI type 4"
	Productname     string `json:"productname"`  // 服务器型号(代号) "DMI type 1"
	Serialnumber    string `json:"serialnumber"` // 主板编号	"DMI type 1"
	Virtual         string `json:"virtual"`      // 是否为虚拟机（physical为物理机，其它全为虚拟机）
	Mac             string `json:"mac"`
	Eth1            string `json:"eth_1"`
	Eth2            string `json:"eth_2"`
	Eth3            string `json:"eth_3"`
	Eth4            string `json:"eth_4"`
	Biosreleasedate string `json:"biosreleasedate"`
	Biosversion     string `json:"biosversion"`
	AgentVersion    string `json:"agent_version"`
}

type DialData struct {
	Code            uint        `json:"code"`              // 状态码
	Active          string      `json:"active"`            // 动作，根据动作调用相关方法
	ToUser          string      `json:"to_user"`           // 发送给谁
	Msg             string      `json:"msg"`               // 说明
	SecWebsocketKey string      `json:"sec_websocket_key"` // websocket通信key
	CallBackActive  string      `json:"call_back_active"`  // 回调方法
	CallBackUser    string      `json:"call_back_user"`    // 回调用户
	Auth            string      `json:"auth"`              // 权限
	Agent           bool        `json:"agent"`             // 是否为agent
	Data            interface{} `json:"data"`              // 返回数据
	Token           string      `json:"token"`             // token
	Secret          bool        `json:"secret"`            // 是否私聊
	Manager         bool        `json:"manager"`           // 是否为管理员 -> gcmdb 分配任务认证
	ToUserList      []string    `json:"to_user_list"`      // 同时发送多个用户
	//AgentVersion    string `json:"agent_version"`
}

type WebSocketLogin struct {
	Token string
}

type WebSocketConnect struct {
	Data interface{}
}

type DialDataRst struct {
	Data   interface{}                     // 返回数据
	Code   uint   `json:"code"`            // 状态码
	Msg    string `json:"msg"`             // 说明
	Time   string `json:"time" xml:"time"` // 当前时间
	Active string `json:"active" xml:"active"`
}

type AgentLogin struct {
	Eth1               string `json:"eth_1"`
	Eth2               string `json:"eth_2"`
	Mac                string `json:"mac"`
	AgentVersion       string `json:"agent_version" xml:"agent_version"`
	Serialnumber       string `json:"serialnumber"`
	SerialnumberStatus bool   `json:"serialnumber_status"`
	Token              string `json:"token"` // token
}

type UserInfo struct {
	User    string `json:"user"`
	AuthKey string `json:"auth_key"`
	Uid     int    `json:"uid"`
	Gid     int    `json:"gid"`
	Admin   bool   `json:"admin"`
}

type Safety struct {
	Tcp          SafetyTcps `json:"tcp_port"`
	TcpCount     int64      `json:"tcp_count"`
	TcpLog       []string   `json:"tcp_log"`
	Udp          Safetyudps `json:"udp_port"`
	UdpLog       []string   `json:"udp_log"`
	UdpCount     int64      `json:"udp_count"`
	Process      []string   `json:"process"`
	ProcessCount int64      `json:"process_count"`
}

type SafetyTcp struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type SafetyTcps []SafetyTcp

type SafetyUdp struct {
	Name string `json:"name"`
	Port string `json:"port"`
}
type Safetyudps []SafetyUdp

type ShellStatus struct {
	Code int         `json:"status"`
	Msg  string      `json:"Msg"`
	Bash string      `json:"bash"`
	Data interface{} `json:"data"`
}

type ShellData struct {
	Data  string `json:"Data"`
	Token string `json:"Token"`
	Salt  string `json:"Salt"`
}

type RstMsg struct {
	RstCode int         `json:"rst_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type PushData struct {
	Title           string `json:"title"`                             // 版本名称
	ReleaseVersion  string `json:"release_version"`                   // 版本号
	Remarks         string `json:"remarks"`                           // 备注
	Auto            bool   `json:"auto"`                              // 是否自动更新
	Reset           bool   `json:"reset"`                             // 是否回退
	Global          bool   `json:"global"`                            // 是否全部更新
	Release         string `json:"release"`                           // 发布指令
	ReleaseArgs     string `json:"release_args"`                      // 发布参数
	Rollback        string `json:"rollback"`                          // 回滚指令
	RollbackArgs    string `json:"rollback_args"`                     // 回滚参数
	ReleaseTimeout  int    `json:"release_timeout"`                   // 发布超时时间
	RollbackTimeout int    `json:"rollback_timeout"`                  // 发布超时时间
	ID              uint   `json:"id"`                                // 项目版本ID
	ReleaseStatus   bool   `json:"release_status"`                    // 发布状态
	RollbackStatus  bool   `json:"rollback_status"`                   // 回滚状态
	Ip              string `json:"ip"`                                // 发布主机
	CheckVersion    bool   `json:"check_version" xml:"check_version"` // 是否检测版本
}

type GetCodeVersion struct {
	Title      string `json:"title"`       // 版本名称
	Release    string `json:"release"`     // 版本号
	BashShell  string `json:"bash_shell"`  // 发布后执行脚本
	ID         uint   `json:"id"`          // 项目版本ID
	BashStatus bool   `json:"bash_status"` // 发布后执行脚本状态，成功、失败
	Env        string `json:"env"`         // 环境
	AllProject bool   `json:"all_project"` // 全部项目
}

/**
 * @Description: http探测
 * @author: 苦咖啡
 * @Date:   2018/9/19 上午10:51
*/
type CheckHttp struct {
	Url        string `json:"url" xml:"url"`                 // 检测url
	TimeOut    int16  `json:"time_out" xml:"time_out"`       // 检测超时时间
	Elapsed    string `json:"elapsed" xml:"elapsed"`         // 耗时
	StatusCode int    `json:"status_code" xml:"status_code"` // 状态
}

type CheckDns struct {
	Domain string   `json:"domain" xml:"domain"` // 域名
	Ip     []string `json:"ip" xml:"ip"`         // ip
}

type CheckPIngTimeout struct {
	Args       string  `json:"args" xml:"args"`               // 域名或ip
	AvgRtt     string  `json:"avg_rtt" xml:"avg_rtt"`         // 平均值
	PacketLoss float64 `json:"packet_loss" xml:"packet_loss"` // 丢失率
}
type TaskMsg struct {
	Ip        string `json:"ip" xml:"ip"`
	Title     string `json:"title" xml:"title"`
	Status    int    `json:"status" xml:"status"`
	Msg       string `json:"msg" xml:"msg"`
	Version   string `json:"version" xml:"version"`
	ReleaseId uint   `json:"release_id" xml:"release_id"`
	Reset     bool   `json:"reset" xml:"reset"`
	Time      string `json:"time" xml:"time"` // 当前时间
	Active    string `json:"active" xml:"active"`
	ID        uint   `json:"id" xml:"id"`
}
