package modbus

import (
	"log"
	"modbus-spyder/internal/app/model"
	"strconv"
	"time"
)

// 数据解析模块
// 这里主要有两种电表类型
// 电表有三相，三相数据都要

// 先判断是哪种电表
// 再进行解析
func MsgParsing(msg []byte, mb *TCPClientHandler) (datas []model.DeviceInfoEntity) {
	// 判断是哪种电表
	// 1.可以根据数据长度来判断
	// 2.查数据库
	if len(msg) == 70+5 {
		// 华立电表
		datas = HuaLiParsing(msg, mb)
	} else if len(msg) == 106+5 {
		// 北丰电表
		datas = BeiFengParsing(msg, mb)
	} else {
		log.Println("数据长度：", len(msg))
		log.Println("数据长度错误，舍弃")
	}

	return
}

// 北丰电表
func BeiFengParsing(msg []byte, mb *TCPClientHandler) (datas []model.DeviceInfoEntity) {
	var data model.DeviceInfoEntity
	for i := 0; i < 7; i++ {
		data.IMEI = "i + 1"
		data.U = float64(DataJointFour(msg, 3+14*i, 2)) / 100    // 电压
		data.I = float64(DataJointFour(msg, 5+14*i, 2)) / 100    // 电流
		data.E = float64(DataJointFour(msg, 9+14*i, 4)) / 100    // 电能
		data.PF = float64(DataJointFour(msg, 13+14*i, 2)) / 1000 // 功率因数
		data.P = float64(DataJointFour(msg, 7+14*i, 2))          // 有功功率
		data.Pts = mb.lastSuccess
		data.Ts = time.Now()
		datas = append(datas, data)
	}
	return
}

// 华立电表
func HuaLiParsing(msg []byte, mb *TCPClientHandler) (datas []model.DeviceInfoEntity) {
	var data model.DeviceInfoEntity
	for i := 0; i < 3; i++ {
		data.IMEI = mb.Address + "|" + strconv.Itoa(int(mb.SlaveID)) + "|" + strconv.Itoa(i+1)
		data.U = float64(DataJointFour(msg, 3+2*i, 2)) / 10     // 电压
		data.I = float64(DataJointFour(msg, 9+2*i, 2)) / 1000   // 电流
		data.E = float64(DataJointFour(msg, 15+4*i, 4)) / 1000  // 电能
		data.PF = float64(DataJointFour(msg, 31+2*i, 2)) / 1000 // 功率因数
		data.P = float64(DataJointFour(msg, 39+4*i, 4))         // 有功功率
		data.Pts = mb.lastSuccess
		data.Ts = time.Now()
		datas = append(datas, data)
	}
	return
}

func DataJointFour(msg []byte, start int, num int) int64 {
	var aData string
	for i := 0; i < num; i++ {
		crcRight := strconv.FormatInt(int64(msg[start+i]), 16)
		if len(crcRight) == 1 {
			aData = aData + "0" + crcRight
		} else {
			aData = aData + crcRight
		}
	}

	intd, err := strconv.ParseInt("0"+aData, 16, 0)
	if err != nil {
		return intd
	}
	return intd
}
