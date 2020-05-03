package Crontab

/**
 * This file is part of Gourd.
 *
 * @link     http://gourd.kyour.cn
 * @document http://gourd.kyour.cn/doc
 * @contact  kyour@vip.qq.com
 * @license  https://https://github.com/kyour-cn/gourd/blob/master/LICENSE
 */

import (
	app_cron "github.com/kyour-cn/gourd/application/app-cron"
	"log"
)

//注册任务
func Register() {

	//注册任务-根据名称获取 conf/cron.conf 配置对应的时间规则
	app_cron.Register("test", func() {

		log.Println("Crontab 定时任务 ：test")

	})

	//直接定义 任务
	//app_cron.AddFunc("0/10 * * * * *", func() {
	//	log.Println("test 0/10")
	//})

}
