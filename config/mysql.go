package config

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // mysql主机
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // 端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         //数据库配置
	DbName       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                         // 数据库名称
	User         string `mapstructure:"user" json:"user" yaml:"user"`                               // 数据库用户
	Passwd       string `mapstructure:"passwd" json:"passwd" yaml:"passwd"`                         // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mod" json:"log-mod" yaml:"log-mod"`                      // 日志级别
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Passwd + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
}
func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
