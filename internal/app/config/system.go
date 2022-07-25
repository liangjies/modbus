package config

type System struct {
	DbType   string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`       // 数据库类型
	ReadOnly bool   `mapstructure:"read-only" json:"read-only" yaml:"read-only"` // 是否只读
	Interval int    `mapstructure:"interval" json:"interval" yaml:"interval"`    // 采集间隔
}
