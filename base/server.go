package base

import (
	"KKchatRoom/base/tools"
	"fmt"
	"log/slog"
	"net"
	"net/http"
)

// Start 启动客户端管理器，注册路由并监听 HTTP 服务。
func Start() {
	hub := newHub() /**/
	go hub.run()
	//配置路由
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/wjsm", wjsm)
	mux.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		OnWs(hub, writer, request)
	})
	config, err := tools.LoadConfig("?")
	if err != nil {
		slog.Error("Start err:", "tools.LoadConfig(", err.Error())
		return
	}
	listener, err := net.Listen("tcp", config.IP)
	if err != nil {
		slog.Error("监听失败", "err", err)
		return
	}
	slog.Info("服务监听成功", "addr", listener.Addr().String())
	//启动服务
	err = http.Serve(listener, mux)
	if err != nil {
		slog.Error("Start err", "http.Serve err", err.Error())
		return
	}
}

// hello 返回固定的问候文本，供接口连通性测试使用。
func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		slog.Error(fmt.Sprintf("write hello err:%v", err))
	}
}

// wjsm 根据请求中的 qqid 参数返回对应的姓名文本。
func wjsm(w http.ResponseWriter, r *http.Request) {
	qqid := r.URL.Query().Get("qqid")
	if qqid == "173930040" {
		_, err := w.Write([]byte("我叫金宇彤"))
		if err != nil {
			slog.Error(fmt.Sprintf("write hello err:%v", err))
		}
	}
}

// OnWs 将 HTTP 请求升级为 WebSocket，注册客户端并启动消息收发。
func OnWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error(err.Error())
	}
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		slog.Error("解析客户端地址失败", "err", err)
		return
	}
	slog.Info("客户端连接", "ip", ip, "port", port)
	client := NewClient(conn, hub)
	client.hub.register <- client
	go client.writePump()
	go client.readPump()
}
