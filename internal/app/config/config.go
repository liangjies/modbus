package config

type Server struct {
	// 日志
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// 数据库
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// Redis

}
