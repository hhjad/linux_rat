package main

//=====
import (
	"fmt"

	"Server"
	"time"
	"ini"
	"g"
	//"runtime"
	//"os"
	//"encoding/json"
)

var Dl_id = ""
var Www_url = "http://ej.ifinetop.com" //"http://b2b.com"    //

func try_Err() {
	if err := recover(); err != nil {
		fmt.Println(err) //这里的err其实就是panic传入的内容，55
	}
}
//=============================================
func run() {
	defer try_Err() //异常处理
	fmt.Printf("xxxxxx")
}



func main() {

	//ff:=g.Unix_Time("1555314603")
	//fmt.Printf(ff)
	//==============================
	g.Log_msgbox( "===========zzkey.com    LINUX集群控制(LINUX反弹式远控)===========")
	g.Log_msgbox( "========================QQ:879301117============================")
	//fmt.Printf("启动websocket")
	//go wss.H.Run()
	Server_ip := string(ini.GetValue_ini("server.ini", "Server", "Server_ip"))
	admin_post := string(ini.GetValue_ini("server.ini", "Server", "admin_port"))

	go func() { //启动WWW
		time.Sleep(time.Second * 1)
		Server.Web_run(Server_ip, admin_post) //启动服务
	}()
	//==============================
	for { //死循环
		time.Sleep(10 * time.Second)
	}
}

/*
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	fmt.Printf("启动websocket")
	for {
		select {
		case conn := <-manager.register:
			fmt.Printf("xxx1111111111")
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			fmt.Printf("xxx2222222222")
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			fmt.Printf("xxx3333333333")
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			fmt.Printf("xxxaaaaaaaaa")
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func main() {
	fmt.Println("Starting application...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":12345", nil)
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	ss,_:=uuid.NewV4()
	client := &Client{id: ss.String(), socket: conn, send: make(chan []byte)}

	manager.register <- client

	go client.read()
	go client.write()
}
*/
