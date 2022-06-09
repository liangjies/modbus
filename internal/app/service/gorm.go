package service

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/model"
)

// 获取采集点
func GetCollectPoint() (err error, points []model.CollectPoint) {
	db := global.SYS_DB.Model(&model.CollectPoint{})
	err = db.Find(&points).Where("del_flag=1").Error
	return
}
