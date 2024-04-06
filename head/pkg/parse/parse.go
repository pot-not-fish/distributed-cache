package parse

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Admin Admin
}

type Admin struct {
	Username string
	Password string
}

var (
	ConfigStructure *Config
	once            sync.Once
)

func Init(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	once.Do(func() {
		ConfigStructure = new(Config)
		if err := viper.Unmarshal(ConfigStructure); err != nil {
			panic(err)
		}
	})
}
