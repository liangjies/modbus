package modbus

import (
	"fmt"
	"testing"
)

func TestModbus(t *testing.T) {
	// 建立 TCP 连接
	mb := NewTCPClientHandler("192.168.1.1:80")
	err := mb.Connect()
	if err != nil {
		fmt.Println(err)
	}
}
