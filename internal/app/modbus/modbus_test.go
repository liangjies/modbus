package modbus

import (
	"fmt"
	"testing"
)

func TestModbus(t *testing.T) {
	// 建立 TCP 连接
	mb := NewTCPClientHandler("192.168.100.220:26")
	err := mb.Connect()
	mb.SlaveID = 47
	if err != nil {
		fmt.Println(err)
	}
	// 接收数据
	if err = mb.Receive(); err != nil {
		fmt.Println(err)
	}
}
