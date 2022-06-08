package model

import "time"

type DeviceInfoEntity struct {
	IMEI string
	Ts   time.Time // 采集时间
	Pts  time.Time // 上次采集时间
	U    float64   // 电压
	I    float64   // 电流
	E    float64   // 电能
	PF   float64   // 功率因数
	P    float64   // 有功功率
	Q    float64   // 无功功率
	F    float64   // 频率
}
