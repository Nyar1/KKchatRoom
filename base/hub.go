package base

// Hub 管理在线客户端，并将消息广播给所有客户端。
type Hub struct {
	// 已注册的客户端。
	clients map[*Client]bool

	// 接收客户端发来的消息的通道。
	broadcast chan []byte

	// 接收客户端注册请求的通道。
	register chan *Client

	// 接收客户端退出请求的通道。
	unregister chan *Client
}

// newHub 创建客户端管理器并初始化消息与连接管理通道。
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// run 持续处理客户端注册、退出和消息广播。
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
