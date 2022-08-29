package config

// 配置MySQL连接参数
const (
	Username = "root"      // 账号
	Password = ""          // 密码 Todo
	Host     = "127.0.0.1" // 数据库地址，可以是Ip或者域名
	Port     = 3306        // 数据库端口
	Dbname   = "users"     // 数据库名
	Timeout  = "10s"       // 连接超时，10秒
)
