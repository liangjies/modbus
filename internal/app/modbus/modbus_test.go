package modbus

import (
	"fmt"
	"testing"
	"time"
)

func TestModbus(t *testing.T) {
	// 建立 TCP 连接
	mb := NewTCPClientHandler("192.168.10.168:8234")
	err := mb.Connect()
	mb.SlaveID = 47
	if err != nil {
		fmt.Println(err)
	}
	// ctx, cancel := context.WithCancel(context.Background()) // 创建一个context
	// 接收数据
	// if err = mb.Receive(ctx); err != nil {
	// 	fmt.Println(err)
	// }

	// 发送数据
	if err = mb.Send([]byte{0x15, 0x12}); err != nil {
		fmt.Println(err)
	}
	time.Sleep(20 * time.Second)
	// cancel()
}
