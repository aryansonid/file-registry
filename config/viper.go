package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config ...
type Config interface {
	ReadRpcUrlConfig() string
	ReadContractAddressConfig() string
	ReadPrivateKey() string
}

type ViperConfig struct{}

func (v *ViperConfig) Init(path string, configName string) {
	viper.AutomaticEnv()
	viper.AddConfigPath(path)
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func (v *ViperConfig) GetString(key string) string {
	return viper.GetString(key)
}

// NewViperConfig creates new viper for reading config.json
func NewViperConfig(path string, configName string) Config {
	v := &ViperConfig{}
	v.Init(path, configName)
	return v
}
