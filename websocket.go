/*
hasaki-quant websocket server
*/
package main

import "net/http"
import "log"

import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{}
type Client struct{
	hub *Hub
	conn *websocket.Conn
	send chan []byte
}
type Hub struct{
	clients map[*Client]bool // 客户端登陆
	broadcast chan []byte    // 广播的消息
	register chan *Client    // 客户端的登陆信息
	unregister chan *Client  // 客户端退出登陆
}

// 接收客户端发来的请求
func recv(w http.ResponseWriter,r *http.Request){
	c,_:=upgrader.Upgrade(w,r,nil)
	defer c.Close()
	for {  // 收到send后的处理
		mt,message,_:=c.ReadMessage()
		c.WriteMessage(mt,append([]byte("发送"),message[:]...))
	}
}

// 发送信息到客户端
func send(w http.ResponseWriter,r *http.Request){
	c,_:=upgrader.Upgrade(w,r,nil)
	defer c.Close()
	// s:=new(PriceAndVolume)
	// s.Price=1
	// s.Volume=2
	// data:=structToJson(s)
	// for {
	// 	// 获取行情
	// 	c.WriteMessage(1,append([]byte(data)))
	// }
}

// 交易所路由,收到的消息推到接收模块，接收行情模块的数据推送
func route(w http.ResponseWriter,r *http.Request){
	c,_:=upgrader.Upgrade(w,r,nil)

	defer c.Close()
	for {  // 收到send后的处理
		mt,message,_:=c.ReadMessage()   // message是byte类型
		c.WriteMessage(mt,append(message[:]))
		
	}
}

func websocketMain(){
	http.HandleFunc("/",route)    // 接收到的信息
	log.Fatal(http.ListenAndServe("0.0.0.1:8888",nil))   // 监听端口
}