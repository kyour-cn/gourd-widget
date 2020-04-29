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
	"github.com/gorilla/mux"
	"net/http"
)

var prefix = "/index"

//注册方法名
func Register(r *mux.Router) {

	//定义方法 url= /index/json
	r.PathPrefix(prefix + "/json").HandlerFunc(jsonTest)

	//默认方法，必须放在最后
	r.PathPrefix(prefix).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("index：" + r.URL.Path))
		})

}

//定义控制器
func jsonTest(w http.ResponseWriter, r *http.Request) {

	j, _ := json.Marshal(r.URL)

	_, _ = w.Write([]byte(string(j)))

}
