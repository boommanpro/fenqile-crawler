package config

import "go.uber.org/config"

type ServerChan struct {
	Server    string `yaml:"server"`
	SecretKey string `yaml:"secret-key"`
	Debug     bool   `yaml:"debug"`
	Proxy     bool   `yaml:"proxy"`
}

type TentcentCos struct {
	BaseUrl   string `yaml:"base-url"`
	SecretId  string `yaml:"secret-id"`
	SecretKey string `yaml:"secret-key"`
	Debug     bool   `yaml:"debug"`
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

func GetTencentCosForYml(path string) *TentcentCos {
	provider, err := config.NewYAML(config.File(path))
	if err != nil {
		panic(err) // handle error
	}

	var tencentCos TentcentCos
	if err := provider.Get("application.tencent-cos").Populate(&tencentCos); err != nil {
		panic(err) // handle error
	}
	return &tencentCos
}
