package config

type ServerConfig struct {
	Addr           string       `mapstructure:"addr" json:"addr"`
	ReadTimeout    int          `mapstructure:"read_timeout" json:"read_timeout"`
	WriteTimeout   int          `mapstructure:"write_timeout" json:"write_timeout"`
	MaxHeaderBytes int          `mapstructure:"max_header_bytes" json:"max_header_bytes"`
	MySQL          MySQLConfig  `mapstructure:"mysql" json:"mysql"`
	Wx             WxConfig     `mapstructure:"wx" json:"wx"`
	Redis          RedisConfig  `mapstructure:"redis" json:"redis"`
	AliOSS         AliOSSConfig `mapstructure:"alioss" json:"alioss"`
}

type MySQLConfig struct {
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	DbName   string `json:"db_name" mapstructure:"dbname"`
}

type WxConfig struct {
	AppID     string `json:"app_id" mapstructure:"appid"`
	Secret    string `json:"secret" mapstructure:"secret"`
	GrantType string `json:"grant_type" mapstructure:"grant_type"`
}

type RedisConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
}

type AliOSSConfig struct {
	AccessKey       string `json:"access_key" mapstructure:"access_key"`
	AccessKeySecret string `json:"access_key_secret" mapstructure:"access_key_secret"`
	EndPoint        string `json:"end_point" mapstructure:"end_point"`
	Bucket          string `json:"bucket" mapstructure:"bucket"`
	BaseUrl         string `json:"base_url" mapstructure:"base_url"`
}
