package service

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/model"

	"go.uber.org/zap"
)

// 数据存入数据库
func PutData(datas []model.DeviceInfoEntity) {
	for _, data := range datas {
		db := global.SYS_DB.Model(&model.DeviceInfoEntity{})
		err := db.Create(&data).Error
		if err != nil {
			global.SYS_LOG.Error("插入数据失败", zap.Any("err", err))
		}
	}

}

// 判断数据是否是需要的数据
func IsNeedData(data model.DeviceInfoEntity) bool {

}
