package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig(configFile string) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../config") // path to look for the config file in

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
}

type AppConfig struct{}

type Configuration interface {
	GetServerHost() string
	GetServerPort() int
	// redis config
	GetRedisHostname() string
	GetRedisPort() int
	GetRedisDatabase() int
	GetRedisPassword() string
}

func NewConfig() Configuration {
	return &AppConfig{}
}

func (c *AppConfig) GetServerHost() string {
	return viper.GetString("server.host")
}

func (c *AppConfig) GetServerPort() int {
	return viper.GetInt("server.port")
}

func (c *AppConfig) GetRedisHostname() string {
	return viper.GetString("redis.host")
}

func (c *AppConfig) GetRedisPort() int {
	return viper.GetInt("redis.port")
}

func (c *AppConfig) GetRedisDatabase() int {
	return viper.GetInt("redis.database")
}

func (c *AppConfig) GetRedisPassword() string {
	return viper.GetString("redis.password")
}
