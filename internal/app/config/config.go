package config

type Server struct {
	// 日志
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// 数据库
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// MongoDB
	MongoDB MongoDB `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
	// 系统配置
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
