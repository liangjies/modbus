package model

// MySQL数据库表
type CollectPoint struct {
	ID          uint   `gorm:"column:id"`
	CollectCode string `gorm:"column:collect_code"` // 采集点编号
	MeterType   string `gorm:"column:meter_type"`   // 表类型
	DelFlag     uint   `gorm:"column:del_flag"`     // 删除标志
}
