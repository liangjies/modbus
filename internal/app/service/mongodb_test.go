package service

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/initialize"
	"modbus-spyder/internal/app/model"
	"testing"
)

func TestMongoDB(t *testing.T) {
	initialize.Viper("..\\..\\..\\config\\config.yaml") // 加载配置文件
	global.SYS_LOG = initialize.Zap()                   // 初始化zap日志库
	// 初始化数据库
	global.SYS_MONGODB = initialize.MongoDB() // gorm连接数据库
	// 插入测试
	var data model.DeviceInfoEntity
	data.IMEI = "123456789012345"

	data2 := model.DeviceInfoEntity{IMEI: "592299"}

	datas := []model.DeviceInfoEntity{data, data2}

	_ = PutDataInMongoDB(datas)

}
