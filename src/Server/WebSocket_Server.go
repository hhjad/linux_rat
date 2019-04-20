package Server

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
	"net/http"
	"g"
	"strings"
)

//=====
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
// 服务器配置信息
//var Upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024000,
//	WriteBufferSize: 1024000,
//}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	ip	string  //IP地址
	uuid string //远程ID
	//u_bh string //编号
	u_os string //操作系统
	did string  //代理Id
	bz_name string //备注名称
	id     string  //本地ID
	update_time     string  //更新时间
	add_time     string     //添加时间
	socket *websocket.Conn
	send   chan []byte  //发送数据
	receive_data	[]string //接收数据

}

type Message struct {
	Action    string `json:"action,omitempty"`
	Data string `json:"data,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			//jsonMessage, _ := json.Marshal(&Message{action: "msg",data:"link ok"})   ///A new socket has connected.已连接。
			//manager.broadcast <- jsonMessage
			//manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				//s := fmt.Sprintf("del socket ip:%v \r\n", conn.ip)
				//jsonMessage, _ := json.Marshal(&Message{action: "msg",data:s})   ///删除链接
				//manager.broadcast <- jsonMessage
				//manager.send(jsonMessage, conn)
				//fmt.Printf(s)
				close(conn.send)
				delete(manager.clients, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					//fmt.Printf("del socket 2222222","\r\n")
					s := fmt.Sprintf("del socket ip:%v \r\n", conn.ip)
					fmt.Printf(s)
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
		//time.Sleep(2 * time.Second)
	}
}


//发送消息
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

type WebSocket_Server struct {
	beego.Controller
}
//var upgrader = websocket.Upgrader{}
func (this *WebSocket_Server) ServeWs() {
	defer g.Panic_Err() //异常处理
//func wsPage(res http.ResponseWriter, req *http.Request) {
//开启websocket
	//conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	//conn, error := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if error != nil {
		this.Ctx.WriteString("")   //必须返回一个数据
		return
	}
	ss,_:=uuid.NewV4()
	//xxxxxxxxx
	beego.Debug("open messageid:",ss)
	client := &Client{id: ss.String(), socket: conn, send: make(chan []byte)}
	//beego.Debug("连接websocket ip:", this.Ctx.Request.RemoteAddr)
	//client.ip= this.Ctx.Request.RemoteAddr
	client.ip= strings.Split(conn.RemoteAddr().String(), ":")[0]
	client.update_time=""  //更新时间
	client.add_time=fmt.Sprintf("%v",time.Now().Unix()) //string(time.Now().Unix()) //获取时间戳     //添加时间
	manager.register <- client

	go client.read()
	go client.write()
	go client.ping()	//ping 定时检测
	this.Ctx.WriteString("")   //必须返回一个数据
}

//======================================================
func (c *Client)ping() {     //心跳包
	defer g.Panic_Err() //异常处理
	// 对方读取下次消息等待时间
	pongWait := 60 * time.Second
	// 对方ping周期
	pingPeriod := (pongWait * 9) / 10
	// 定时执行
	//beego.Debug("测试")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		//_ = c.Ws.Close()
	}()
	for {
		defer g.Panic_Err() //异常处理
		select {
			case _, ok := <-c.send:
				if !ok {
					c.socket.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
			case <-ticker.C:
				if err := c.socket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					return
				}
		}
	}
}

func (c *Client) write() {   //发送数据
	defer g.Panic_Err() //异常处理
	defer func() {
		c.socket.Close()
	}()
	for {
		select {
			case message, ok := <-c.send:
				if !ok {
					c.socket.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				if err := c.socket.WriteMessage(websocket.TextMessage, message); err != nil {
					return
				}
		}
	}
}

func open_map_key(Data_list map[string]interface{}, key string)(bool,string){  //读取MAP 数据
	defer g.Panic_Err() //异常处理
	if _, ok := Data_list[key]; ok {
		return true,fmt.Sprintf("%v", Data_list[key])
	}
	return false,""
}

func (c *Client) read() {   //读取
	defer g.Panic_Err() //异常处理
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()
	c.socket.SetReadLimit(maxMessageSize)
	_ = c.socket.SetReadDeadline(time.Now().Add(pongWait))
	c.socket.SetPongHandler(func(string) error { c.socket.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			break
		}
		//msg_data:=string(message)
		//fmt.Printf(string(message))
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(string(message)), &dat); err == nil { //json str 转map
			//s := fmt.Sprintf("=== sss:%+v ===\r\n", c.id)
			//fmt.Printf(s)
			c.update_time=fmt.Sprintf("%v",time.Now().Unix()) // //获取时间戳  //更新时间
			if action, ok := dat["action"]; ok {  //判断键值是否存在
				switch action {
					case "ping":  //ping
						s := fmt.Sprintf("===ping time:%v ip:%s id:%s ===\r\n", time.Now().Format("2006-01-02 15:04:05"),c.ip,c.id)
						fmt.Printf(s)
						//jsonMessage, _ := json.Marshal(&Message{Action: "ping",Data:"ping"})   ///
						//fmt.Printf(string(jsonMessage))
						//manager.broadcast <- jsonMessage

						//shell:="netstat -ano"
						//jsonMessage, _ := json.Marshal(&Message{Action: "Shell",Data:g.Base64_Encode(shell)})   ///
						//fmt.Printf(string(jsonMessage))
						//manager.broadcast <- jsonMessage
						//manager.broadcast <- []byte("xxxxxxxxxxxxxxx")
					case "uuid":  //服务端唯一ID
						bool,data:=open_map_key(dat,"data")
						if bool{
							c.uuid=data
						}
					case "u_os":  //机器配置信息
						bool,data:=open_map_key(dat,"data")
						if bool{
							c.u_os=data
						}
					case "did":  //代理ID
						bool,data:=open_map_key(dat,"data")
						if bool{
							c.did=data
						}
					case "name":  //备注名称
						bool,data:=open_map_key(dat,"data")
						if bool{
							c.bz_name=g.Base64_Decode(data)
						}
					case "Shell":  //shell执行结果
						bool,data:=open_map_key(dat,"data")
						if bool{
							dd:=g.Base64_Decode(data)
							c.receive_data=append(c.receive_data,dd) //添加数据
							if len(c.receive_data)>=100 {
								c.receive_data=[]string{}
							}
							s := fmt.Sprintf("=receive_data==%v===", len(c.receive_data))
							fmt.Printf(s,"\r\n")
						}
					default:
						s := fmt.Sprintf("===err:%v===\r\n", string(message))
						fmt.Printf(s)	//解析在业务逻辑处理
						//jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
						//manager.broadcast <- jsonMessage
				}
			}
		}

		//if err != nil {
		//	manager.unregister <- c
		//	c.socket.Close()
		//	break
		//}
		//jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		//s := fmt.Sprintf("===%v===", string(message))
		////解析在业务逻辑处理
		//fmt.Printf(s,"\r\n")
		//manager.broadcast <- jsonMessage
	}
}




