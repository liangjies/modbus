package initialize

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/service"
	"os"

	"go.uber.org/zap"
)

func Global() {
	// 从MySQL数据库中获取采集点信息
	if err, points := service.GetCollectPoint(); err == nil {
		// 将采集点信息放入缓存中
		temp := make(map[string]string)
		for _, v := range points {
			temp[v.CollectCode] = v.MeterType
		}
		global.CollectPoint = temp
	} else {
		global.SYS_LOG.Error("从MySQL数据库中获取采集点信息执行失败", zap.Any("err", err))
		os.Exit(1)
	}
}
