package config

type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  MySQL  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
}
