package config

type System struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
}
