package modbus

import (
	"bytes"
	"encoding/binary"
	"modbus-spyder/internal/app/utils"
)

// 数据解析模块
// 这里主要有两种电表类型
// 电表有三相，三相数据都要
// 统一读取全部电流电压数据

// 先判断是哪种电表
// 再进行解析
func MsgSendGroup(SlaveID byte, meterType string) (aduRequest []byte) {
	switch meterType {
	case "北丰":
		aduRequest = BeiFengMsgSend(SlaveID)
	case "华立":
		aduRequest = HuaLiMsgSend(SlaveID)
	}
	return
}

// 北丰电表
func BeiFengMsgSend(SlaveID byte) (aduRequest []byte) {
	startReg := 64           // 起始寄存器
	readLength := 53         // 读取长度
	var operate uint8 = 0x03 // 读取操作

	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, SlaveID)
	binary.Write(buf, binary.BigEndian, operate)
	binary.Write(buf, binary.BigEndian, startReg)
	binary.Write(buf, binary.BigEndian, readLength)
	// CRC16校验
	var crc utils.CRC
	crc.Reset()
	crc.PushBytes(buf.Bytes())
	binary.Write(buf, binary.BigEndian, crc.Value())
	aduRequest = buf.Bytes()
	return
}

// 华立电表
func HuaLiMsgSend(SlaveID byte) (aduRequest []byte) {
	var startReg uint16 = 0x00 // 起始寄存器
	var readLength uint16 = 35 // 读取长度
	var operate uint8 = 0x03   // 读取操作

	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, SlaveID)
	binary.Write(buf, binary.BigEndian, operate)
	binary.Write(buf, binary.BigEndian, startReg)
	binary.Write(buf, binary.BigEndian, readLength)
	// CRC16校验
	var crc utils.CRC
	crc.Reset()
	crc.PushBytes(buf.Bytes())
	binary.Write(buf, binary.BigEndian, crc.Value())
	aduRequest = buf.Bytes()
	// fmt.Println(aduRequest)
	return
}
