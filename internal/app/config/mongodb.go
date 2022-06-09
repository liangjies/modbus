package config

type MongoDB struct {
	Path        string `mapstructure:"path" json:"path" yaml:"path"`                          // 服务器地址:端口
	Dbname      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                  // 数据库名
	Username    string `mapstructure:"username" json:"username" yaml:"username"`              // 数据库用户名
	Password    string `mapstructure:"password" json:"password" yaml:"password"`              // 数据库密码
	MaxPoolSize int    `mapstructure:"max-pool-size" json:"maxPoolSize" yaml:"max-pool-size"` // 连接池中的最大连接数
}
