package route

/**
 * This is GourdApp Route profile.
 *
 * @link     http://gourd.kyour.cn
 * @document http://gourd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/gourd/blob/master/LICENSE
 */

import (
	"github.com/kyour-cn/gourd/server/middleware"
	"github.com/kyour-cn/gourd/server/router"
	"gourd-widget/app/Http/controller"
	ws_module "gourd-widget/app/Websocket"
	"net/http"
)

//载入路由定义
func LoadRouter() (route *router.Router) {

	//实例化创建路由
	route = router.NewRouter()

	//=======================================================================
	// HTTP路由定义
	//=======================================================================

	//文字响应输出 -完全匹配
	route.HandleFunc("/gourd", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello Gourd!"))
	})

	//将route对象交给controller/define.go定义控制器路由
	controller.Register(route)

	//=======================================================================
	// Websocket 路由定义
	//=======================================================================

	//ws模块定义
	ws_module.Regist(route)

	//静态资源访问路由 -必须放在最后面，否则上面的路由不生效
	route.PathPrefix("/").
		HandlerFunc(middleware.FileMiddleware)

	return
}
