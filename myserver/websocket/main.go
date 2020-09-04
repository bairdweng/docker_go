package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"com.miaoyou.server/helper"
	"github.com/gorilla/websocket"
)

func main() {
	println("websocket已经启动端口号：5322")
	wsConnAll = make(map[int64]*wsConnection)
	http.HandleFunc("/websocket", wsHandler)
	http.ListenAndServe(":5322", nil)
}

const (
	// 允许等待的写入时间
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// 最大的连接ID，每次连接都加1 处理
var maxConnID int64

// 客户端读写消息
type wsMessage struct {
	// websocket.TextMessage 消息类型
	messageType int
	data        []byte
}

// ws 的所有连接
// 用于广播
var wsConnAll map[int64]*wsConnection
var mt sync.Mutex

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道,加锁处理，这个不起作用，应该采用全局锁
	isClosed  bool
	closeChan chan byte // 关闭通知
	id        int64
	userID    string // 用户id
	toUserID  string // 发送的用户id
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有的CORS 跨域请求，正式环境可以关闭
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// ProtectRun 保护方式允许一个函数 相当于try crach
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()
	entry()
}

/// 查询url后面的参数
func queryParams(key string, req *http.Request) string {
	params := req.URL.Query()
	var value = ""
	values := params[key]
	if len(values) > 0 {
		value = values[0]
	}
	return value
}
func wsHandler(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := upgrader.Upgrade(resp, req, nil)
	// 当前用户id
	uid := queryParams("id", req)
	// 需要发送的用户id
	toID := queryParams("to_id", req)
	if err != nil {
		log.Println("升级为websocket失败", err.Error())
		return
	}

	// TODO 如果要控制连接数可以计算，wsConnAll长度
	// 连接数保持一定数量，超过的部分不提供服务
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		id:        maxConnID,
		userID:    uid,
		toUserID:  toID,
	}
	mt.Lock()
	maxConnID++
	wsConnAll[maxConnID] = wsConn
	mt.Unlock()
	log.Println("当前在线人数", len(wsConnAll))

	// 处理器,发送定时信息，避免意外关闭
	go wsConn.processLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

// 处理队列中的消息
func (wsConn *wsConnection) processLoop() {
	// 处理消息队列中的消息
	// 获取到消息队列中的消息，处理完成后，发送消息给客户端
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			log.Println("获取消息出现错误", err.Error())
			break
		}
		log.Println("接收到消息", string(msg.data))
		// 修改以下内容把客户端传递的消息传递给处理程序
		err = wsConn.wsWrite(msg.messageType, msg.data)
		if err != nil {
			log.Println("发送消息给客户端出现错误", err.Error())
			break
		}
	}
}

// 处理消息队列中的消息
func (wsConn *wsConnection) wsReadLoop() {
	// 设置消息的最大长度
	wsConn.wsSocket.SetReadLimit(maxMessageSize)
	wsConn.wsSocket.SetReadDeadline(time.Now().Add(pongWait))
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			log.Println("消息读取出现错误", err.Error())
			wsConn.close()
			return
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列,消息入栈
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			return
		}
	}
}

// 发送消息给客户端
func (wsConn *wsConnection) wsWriteLoop() {
	// 这是一个打点器
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	// 这里取消for似乎没影响
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			ProtectRun(func() {
				messageMap, err := helper.JSONToMap(string(msg.data))
				if err != nil {
					return
				}
				toID := messageMap["toID"].(string)
				message := messageMap["message"].(string)
				toConn := getConnect(toID)
				if toConn == nil {
					print("用户未找到")
					return
				}
				mData := []byte(message)
				if err := toConn.wsSocket.WriteMessage(msg.messageType, mData); err != nil {
					log.Println("发送消息给客户端发生错误", err.Error())
					// 切断服务
					wsConn.close()
					return
				}
			})
		case <-wsConn.closeChan:
			// 获取到关闭通知
			return
		case <-ticker.C:
			// 出现超时情况
			wsConn.wsSocket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := wsConn.wsSocket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

/// 根据发送id获取请求的链接
func getConnect(id string) *wsConnection {
	var connect *wsConnection
	for _, con := range wsConnAll {
		if con.userID == id {
			connect = con
		}
	}
	return connect
}

// 写入消息到队列中
func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("连接已经关闭")
	}
	return nil
}

// 读取消息队列中的消息
func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		// 获取到消息队列中的消息
		return msg, nil
	case <-wsConn.closeChan:

	}
	return nil, errors.New("连接已经关闭")
}

// 关闭连接
func (wsConn *wsConnection) close() {
	if wsConn.isClosed == false {
		wsConn.isClosed = true
		wsConn.wsSocket.Close()
		// 一定
		// wsConn.mutex.Lock()
		mt.Lock()
		delete(wsConnAll, wsConn.id)
		mt.Unlock()
		// wsConn.mutex.Unlock()
		close(wsConn.closeChan)
	}
}
