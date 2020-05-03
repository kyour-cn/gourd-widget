package controller

/**
 * This file is part of Gourd.
 *
 * @link     http://gourd.kyour.cn
 * @document http://gourd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/gourd/blob/master/LICENSE
 */

import (
	"github.com/kyour-cn/gourd/server/router"
	"gourd-widget/app/Http/controller/index"
	"net/http"
)

//注册控制器
func Register(route *router.Router) *router.Router {

	//index控制器路由
	index.Register(route)

	//api路由
	route.PathPrefix("/api").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Api Controller"))
		})

	return route
}
