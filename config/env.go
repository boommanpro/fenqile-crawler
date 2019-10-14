package config

import "go.uber.org/config"

type ServerChan struct {
	Server    string `yaml:"server"`
	SecretKey string `yaml:"secret-key"`
	Debug     bool   `yaml:"debug"`
	Proxy     bool   `yaml:"proxy"`
}

func GetServerChanForYml(path string) *ServerChan {
	provider, err := config.NewYAML(config.File(path))
	if err != nil {
		panic(err) // handle error
	}

	var serverChan ServerChan
	if err := provider.Get("application.server-chan").Populate(&serverChan); err != nil {
		panic(err) // handle error
	}
	return &serverChan
}
