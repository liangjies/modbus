package timer

import (
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/service"
	"time"

	"go.uber.org/zap"
)

// 定时器
func Timer() {
	// 使用协程运行定时任务
	go CollectPointTimer()
}

// 定时从MySQL数据库中获取采集点信息
// 并将采集点信息放入缓存中
func CollectPointTimer() {
	// 5分钟
	timer := time.Tick(5 * time.Minute)
	for {
		select {
		case <-timer:
			//执行任务
			if err, points := service.GetCollectPoint(); err == nil {
				// 将采集点信息放入缓存中
				var temp map[string]string
				for _, v := range points {
					temp[v.CollectCode] = v.MeterType
				}
				global.CollectPoint = temp
			} else {
				global.SYS_LOG.Error("定时从MySQL数据库中获取采集点信息执行失败", zap.Any("err", err))
			}
		}
	}
}
