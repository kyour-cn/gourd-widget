package main

/**
 * This file is part of Guerd.
 *
 * @link     http://guerd.kyour.cn
 * @document http://guerd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/guerd/blob/master/LICENSE
 */

import (
	"github.com/kyour-cn/guerd"
	"guerd-widget/app/Http/route"
)

func main() {

	//创建Application
	app := guerd.NewApp()

	//指定配置文件
	app.ConfigFile("./app.conf")

	//载入路由
	app.HttpRoute(route.LoadRouter)

	//启动
	app.Serve()

}
