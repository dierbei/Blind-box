package config

type ServerConfig struct {
	Addr           string      `mapstructure:"addr" json:"addr"`
	ReadTimeout    int         `mapstructure:"read_timeout" json:"read_timeout"`
	WriteTimeout   int         `mapstructure:"write_timeout" json:"write_timeout"`
	MaxHeaderBytes int         `mapstructure:"max_header_bytes" json:"max_header_bytes"`
	MySQL          MySQLConfig `mapstructure:"mysql" json:"mysql"`
}

type MySQLConfig struct {
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	DbName   string `json:"db_name" mapstructure:"dbname"`
}
