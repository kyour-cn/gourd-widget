package main

/**
 * This file is part of Gourd.
 *
 * @link     http://gourd.kyour.cn
 * @document http://gourd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/gourd/blob/master/LICENSE
 */

import (
	"github.com/kyour-cn/gourd"
	"gourd-widget/app/Http/route"
	//"gourd-widget/app/Tcp"
)

func main() {

	//创建Application
	app := gourd.NewApp()

	//指定配置文件
	app.ConfigFile("./app.conf")

	//载入Http路由
	app.HttpRoute(route.LoadRouter)

	//注册Tcp事件
	//app.RegistTcp(Tcp.)

	//启动
	app.Serve()

}
