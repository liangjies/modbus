package cmd

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/initialize"
	"modbus-spyder/internal/app/timer"
)

func Execute() {
	initialize.Viper()                // 加载配置文件
	global.SYS_LOG = initialize.Zap() // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
	if global.SYS_DB != nil {
		//initialize.MysqlTables(global.SYS_DB) // 初始化表-暂不用
		// 程序结束前关闭数据库链接
		db, _ := global.SYS_DB.DB()
		defer db.Close()
	}
	timer.Timer() // 加载定时器
}
