package Tcp

import (
	"fmt"
	tcp "github.com/kyour-cn/gourd/server/tcp"
)

//注册事件
func OnEvent() tcp.Event {

	return tcp.Event{
		onConnect,
		onReceive,
		onClose,
	}

}

//新客户端连接事件
func onConnect(conn tcp.Connection) {

	fmt.Printf("新的连接：(%v) %s\n", conn.Fd, conn.Addr)

}

//接收客户端数据事件
func onReceive(conn tcp.Connection, buffer []byte) {

	fmt.Println("来自" + conn.Addr + "新的消息：" + string(buffer))

}

//客户端断开事件
func onClose(conn tcp.Connection) {

	fmt.Printf("连接断开：(%v) %s\n", conn.Fd, conn.Addr)

}
