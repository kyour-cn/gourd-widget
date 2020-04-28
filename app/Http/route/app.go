package route

/**
 * This is GuerdApp Route profile.
 *
 * @link     http://guerd.kyour.cn
 * @document http://guerd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/guerd/blob/master/LICENSE
 */

import (
	"github.com/gorilla/mux"
	app_http "github.com/kyour-cn/guerd/application/app-http"
	"net/http"
)

type Routes struct {
	name string
}

//载入路由定义
func LoadRouter() (route *mux.Router) {

	route = mux.NewRouter()

	//文字响应输出
	route.HandleFunc("/guerd", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello Guerd!"))
	})

	//api路由
	route.PathPrefix("/api/").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Api Controller"))
		})

	//静态资源访问路由
	route.PathPrefix("/").
		HandlerFunc(app_http.FileMiddleware)

	return
}
