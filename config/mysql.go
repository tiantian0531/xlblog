package config

type MySQL struct {
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
}
