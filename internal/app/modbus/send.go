package modbus

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 数据解析模块
// 这里主要有两种电表类型
// 电表有三相，三相数据都要
// 统一读取全部电流电压数据

// 先判断是哪种电表
// 再进行解析
func MsgSendGroup(addr string, meterType string) (aduRequest []byte) {
	return
}

// 北丰电表
func BeiFengMsgSend(addr string) {
	startReg := 64   // 起始寄存器
	readLength := 53 // 读取长度
	_ = startReg
	_ = readLength
	return
}

// 华立电表
func HuaLiMsgSend(addr string) (aduRequest []byte) {
	// startReg := 0    // 起始寄存器
	// readLength := 35 // 读取长度
	// operate := 0x03  // 读取操作

	// addrInt, _ := strconv.ParseUint(addr, 10, 64)
	// var address = make([]byte, 8)
	// binary.BigEndian.PutUint64(address, addrInt)
	// binary.LittleEndian.PutUint64(address, addrInt)
	// binary.LittleEndian.PutUint64(address, addrInt)
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, uint8(35))
	binary.Write(buf, binary.BigEndian, uint8(87))
	binary.Write(buf, binary.BigEndian, uint8(14))
	binary.Write(buf, binary.BigEndian, uint8(254))
	aduRequest = buf.Bytes()
	fmt.Println(aduRequest)
	return
}
