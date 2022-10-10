package spyder

import (
	"context"
	"fmt"
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/modbus"
	"modbus-spyder/internal/app/service"
	"modbus-spyder/internal/app/timer"
	"sync"
)

var PointServer []string

// 采集数据
func Spyder() {
	var wg sync.WaitGroup
	var wgSend sync.WaitGroup
	chReload := make(chan bool)                             // 重新加载采集点
	ctx, cancel := context.WithCancel(context.Background()) // 创建一个context
	// 定时检测采集点是否发生变化
	go timer.CollectPointTimer(chReload)
	// 第一次先加载采集点
	go service.GetCollectPoint(chReload)
	<-chReload
	// 开始采集
	for {
		// 获取串口服务器地址
		GetPointServer()
		global.SYS_LOG.Info("开始采集")
		fmt.Println("PointServer", PointServer)
		// 采集协程
		for _, pointServer := range PointServer {
			wg.Add(1)
			go modbus.RunSpyder(pointServer, ctx, &wg)
		}
		// 发送协程
		if !global.SYS_CONFIG.System.ReadOnly {
			for _, pointServer := range PointServer {
				wgSend.Add(1)
				go modbus.SendSpyder(pointServer, ctx, &wgSend)
			}
		}
		<-chReload // 采集点不发生变化这里一直阻塞
		global.SYS_LOG.Info("采集点发生变化，等待所有协程结束")
		cancel()
		wg.Wait()     // 等待所有协程结束
		wgSend.Wait() // 等待所有协程结束
		global.SYS_LOG.Info("采集点发生变化，协程全部结束")
	}
}

// 获取串口服务器地址
func GetPointServer() {
	temp := make(map[string]bool)
	for k, _ := range global.CollectPoint {
		temp[k] = true
	}
	for k, _ := range temp {
		PointServer = append(PointServer, k)
	}
}
