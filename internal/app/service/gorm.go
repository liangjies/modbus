package service

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/model"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

func GetCollectPoint(chReload chan bool) {
	// 将采集点信息放入缓存中
	temp := make(map[string]string)
	// 设备采集点信息
	if err, points := GetCollectPointByEquipment(); err == nil {
		for _, v := range points {
			temp[v.CollectCode] = v.MeterType
		}
	} else {
		global.SYS_LOG.Error("定时从MySQL数据库中获取设备采集点信息执行失败", zap.Any("err", err))
		return
	}
	// 电房采集点信息
	// 电房采集三相
	if err, points := GetCollectPointByRoom(); err == nil {
		for _, v := range points {
			// 三相采集点
			for i := 1; i <= 3; i++ {
				temp[v.CollectCode[:len(v.CollectCode)-1]+strconv.Itoa(i)] = "华立"
			}
		}
	} else {
		global.SYS_LOG.Error("定时从MySQL数据库中获取电房采集点信息执行失败", zap.Any("err", err))
		return
	}
	if CollectPointIsChanged(temp) {
		// 加锁
		var mutex sync.Mutex
		mutex.Lock()
		// 将采集点信息放入全局变量中
		global.CollectPoint = temp
		mutex.Unlock()
		// 发送通知
		chReload <- true
	}
}

// 检测采集点是否有变化
func CollectPointIsChanged(temp map[string]string) bool {
	// 长度是否相同
	if len(global.CollectPoint) != len(temp) {
		return true
	}
	// 内容是否相同
	for k, v := range temp {
		if v2, ok := global.CollectPoint[k]; !ok {
			return true
		} else {
			if v2 != v {
				return true
			}
		}
	}
	return false
}

// 获取设备采集点
func GetCollectPointByEquipment() (err error, points []model.CollectPoint) {
	db := global.SYS_DB.Model(&model.CollectPoint{})
	err = db.Find(&points).Where("del_flag=1").Error
	return
}

// 获取电房采集点
func GetCollectPointByRoom() (err error, points []model.MeterInfo) {
	db := global.SYS_DB.Model(&model.MeterInfo{})
	err = db.Find(&points).Where("del_flag=1").Error
	return
}
