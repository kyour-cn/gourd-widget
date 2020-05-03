package index

/**
 * This file is part of Gourd.
 *
 * @link     http://gourd.kyour.cn
 * @document http://gourd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/gourd/blob/master/LICENSE
 */

import (
	"encoding/json"
	"fmt"
	"github.com/kyour-cn/gourd/server/router"
	"github.com/kyour-cn/gourd/server/session"
	"net/http"
)

var prefix = "/index"

//注册方法名
func Register(r *router.Router) {

	//定义方法 url= /index/json
	r.PathPrefix(prefix + "/json").HandlerFunc(jsonTest)

	r.PathPrefix(prefix + "/login").HandlerFunc(sessionDemoLogin)

	r.PathPrefix(prefix + "/logout").HandlerFunc(sessionDemoLogout)

	//默认方法，必须放在最后
	r.PathPrefix(prefix).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("index：" + r.URL.Path))
		})

}

//定义路由事件
func jsonTest(w http.ResponseWriter, r *http.Request) {

	j, _ := json.Marshal(r.URL)

	_, _ = w.Write([]byte(string(j)))

}

//session示例
func sessionDemoLogin(w http.ResponseWriter, r *http.Request) {

	s := session.GetSession(w, r)

	_ = s.Set("test", "demo")

	_, _ = fmt.Fprint(w, "sessionId:", s.SessionID(), ",Data:", s.GetAll())
	//_, _ = w.Write([]byte("sessionId:" + s.SessionID()))

}

//session示例
func sessionDemoLogout(w http.ResponseWriter, r *http.Request) {

	s := session.GetSession(w, r)

	_, _ = w.Write([]byte("清空session:" + s.SessionID()))

	//清空
	_ = s.Clear()

}
