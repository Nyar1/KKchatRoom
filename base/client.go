package base

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const (
	// 向客户端写入消息的超时时间。
	writeWait = 10 * time.Second

	// 等待客户端下一次 Pong 心跳响应的超时时间。
	pongWait = 60 * time.Second

	// 向客户端发送 Ping 心跳的间隔，必须小于 pongWait。
	pingPeriod = (pongWait * 9) / 10

	// 允许接收的客户端消息大小上限。
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client 表示一个在线客户端，保存连接和待发送消息。
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan Message

	kkcrId   string
	kkcrName string
}

// NewClient 创建客户端并初始化消息发送队列。
func NewClient(conn *websocket.Conn, hub *Hub, r *http.Request) *Client {
	cId := `1937`
	cName := `KK`
	return &Client{
		hub:      hub,
		conn:     conn,
		send:     make(chan Message, 256),
		kkcrId:   cId,
		kkcrName: cName,
	}
}

// readPump 持续读取客户端消息，整理后交给 Hub 广播。
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, messageBody, err := c.conn.ReadMessage()
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			slog.Error(fmt.Sprintf("read err: %v", err))
			return
		}
		if err != nil {
			slog.Error(fmt.Sprintf("客户端%v接收消息报错", c.kkcrId), "readPump err", err.Error())
			return
		}
		slog.Info(string(messageBody))
		message := Message{}
		if err = json.Unmarshal(messageBody, &message); err != nil {
			message = NewMessage()
			message.Content = string(messageBody)
			message.MessageType = MsgTypeTest
		}
		message.Check(c)
		c.hub.broadcast <- message
	}
}

// writePump 向客户端发送队列中的消息，并定期发送心跳。
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case messageBody, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Hub 已关闭消息发送通道。
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			sendMessage(messageBody, w)

			//n := len(c.send)
			//for i := 0; i < n; i++ {
			//	w.Write(newline)
			//	sendMessage(<-c.send, w)
			//}
			if err = w.Close(); err != nil {
				if websocket.IsUnexpectedCloseError(
					err,
					websocket.CloseGoingAway,
					websocket.CloseAbnormalClosure,
				) {
					slog.Error(fmt.Sprintf("w.Close() err: %v", err))
				}
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			c.conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}

func sendMessage(messageBody Message, w io.WriteCloser) {
	message, err := json.Marshal(messageBody)
	if err != nil {
		slog.Error("writePump()", "json.Marshal(messageBody)", err.Error())
		return
	}
	w.Write(message)
}
