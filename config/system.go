package config

type System struct {
	ENV    string `mapstructure:"env" json:"env" yaml:"host"`            // mysql主机
	Addr   string `mapstructure:"addr" json:"addr" yaml:"host"`          // mysql主机
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // mysql主机
}
