package index

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var prefix = "/index"

//注册方法名
func Regist(r *mux.Router) {

	//定义方法
	r.PathPrefix(prefix + "/json").HandlerFunc(jsonTest)

	//默认方法，最后定义
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
