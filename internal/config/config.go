package config

import (
	"flag"
	"github.com/spf13/viper"
)

type Config struct {
	ServerIp   string `mapstructure:"SERVER_IP"`
	ServerPort uint   `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (Config, error) {
	var ret Config
	viper.AddConfigPath(path)
	viper.SetConfigName("serverConfig")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return ret, err
	}

	err = viper.Unmarshal(&ret)
	ret.ServerIp = *flag.String("server_ip", ret.ServerIp, "")
	ret.ServerPort = *flag.Uint("server_port", ret.ServerPort, "")

	return ret, nil
}
