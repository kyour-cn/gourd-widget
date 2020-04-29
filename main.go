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
	crontab "gourd-widget/app/Crontab"
	"gourd-widget/app/Http/route"
	tcp "gourd-widget/app/Tcp"
)

func main() {

	//创建Application
	app := gourd.NewApp()

	//指定配置文件
	app.ConfigFile("./conf/app.conf")

	//载入Http路由
	app.HttpRoute(route.LoadRouter)

	//注册Tcp事件
	app.RegistTcp(tcp.OnEvent())

	//注册CronTab任务
	crontab.Register()

	//启动
	app.Serve()

}
