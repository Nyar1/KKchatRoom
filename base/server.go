package base

import (
	"fmt"
	"log/slog"
	"net/http"
)

func Start() {
	hub := newHub() /**/
	go hub.run()
	//配置路由
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		OnWs(hub, writer, request)
	})
	//启动服务
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		slog.Error(fmt.Sprintf("Start::http.ListenAndServ err:%v", err))
		return
	}
	slog.Info(`服务启动成功`)
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		slog.Error(fmt.Sprintf("write hello err:%v", err))
	}
}
