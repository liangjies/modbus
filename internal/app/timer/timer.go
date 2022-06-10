package timer

import (
	"modbus-spyder/internal/app/service"
	"time"
)

// 定时器
func Timer() {
	// 使用协程运行定时任务
	//go CollectPointTimer()
}

// 定时从MySQL数据库中获取采集点信息
// 并将采集点信息放入缓存中
func CollectPointTimer(chReload chan bool) {
	// 5分钟
	timer := time.Tick(5 * time.Minute)
	for {
		select {
		case <-timer:
			//执行任务
			service.GetCollectPoint(chReload)
		}
	}
}
