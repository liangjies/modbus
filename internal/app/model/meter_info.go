package model

// MySQL数据库表-电房电表
type MeterInfo struct {
	ID          uint   `gorm:"column:id"`
	CollectCode string `gorm:"column:collect_point_code"` // 采集点编号
	DelFlag     uint   `gorm:"column:del_flag"`           // 删除标志
}
