package Websocket

/**
 * This is GourdApp Route profile.
 *
 * @link     http://gourd.kyour.cn
 * @document http://gourd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/gourd/blob/master/LICENSE
 */

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kyour-cn/gourd/server/ws"
	"net/http"
	"strconv"
)

//url
var prefix = "/ws"

//注册ws模块
func Regist(r *mux.Router) {

	//定义路由
	r.PathPrefix(prefix).HandlerFunc(Demo)

}

//ws消息模型示例
func Demo(writer http.ResponseWriter, request *http.Request) {

	//封装ws握手，维护连接池
	ws.Handshake(writer, request, func(conn ws.Connection) {
		//新连接
		fmt.Printf("新的连接：%v\n", conn.Fd)
		err := conn.Send(ws.TextMessage, []byte("Hello "+strconv.Itoa(int(conn.Fd))))
		if err != nil {
			fmt.Println("发送消息失败")
		}
	}, func(conn ws.Connection, msg ws.Message) {
		//收到消息

		//回复消息
		//err := conn.Send(msg.Type, []byte("收到:"+string(msg.Data)))
		//if err != nil {
		//	fmt.Println("回复消息失败")
		//}

		//向所有连接广播消息 并获得发送（成功和失败）次数
		ss, es := ws.BroadcastMessage([]byte("收到Fd(" + strconv.Itoa(int(conn.Fd)) + "):" + string(msg.Data)))

		fmt.Printf("发送成功%v，失败%v\n", ss, es)

		//向fd为参数1的连接发送数据消息
		//_ = ws.SendWsMessage(fd, "来自"+strconv.Itoa(conn.Fd)+"的消息")

	}, func(conn ws.Connection) {
		//连接断开
		fmt.Printf("连接%v关闭\n", conn.Fd)
	})
}
