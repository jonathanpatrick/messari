package config

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Host    string `yaml:"host"`
		Timeout int    `yaml:"timeout"`
		// ApiKey string `yaml:"api_key"`
	} `yaml:"server"`
}
