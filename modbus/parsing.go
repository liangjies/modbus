package modbus

import (
	"modbus-spyder/model"
	"strconv"
)

// 数据解析模块
// 这里主要有两种电表类型
// 电表有三相，三相数据都要
func BeiFengParsing(msg []byte) (datas []model.DeviceInfoEntity) {
	var data model.DeviceInfoEntity
	for i := 0; i < 7; i++ {
		data.IMEI = "i + 1"
		data.U = float64(DataJointFour(msg, 3+14*i, 2)) / 100
		data.I = float64(DataJointFour(msg, 5+14*i, 2)) / 100
		data.E = float64(DataJointFour(msg, 9+14*i, 4)) / 100
		data.PF = float64(DataJointFour(msg, 13+14*i, 2)) / 1000
		datas = append(datas, data)
	}
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
