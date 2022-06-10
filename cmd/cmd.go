package cmd

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/initialize"
	"modbus-spyder/internal/app/modbus"
	"modbus-spyder/internal/app/timer"
)

func Execute() {
	initialize.Viper()                // 加载配置文件
	global.SYS_LOG = initialize.Zap() // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
	if global.SYS_DB != nil {
		initialize.MysqlTables(global.SYS_DB) // 初始化表-暂不用
		// 程序结束前关闭数据库链接
		db, _ := global.SYS_DB.DB()
		defer db.Close()
	}
	initialize.Global() // 初始化全局变量

	timer.Timer()                             // 加载定时器
	go modbus.RunSpyder("192.168.100.220:26") // 开启采集
	// 阻塞运行
	ch := make(chan int, 1)
	<-ch
}
