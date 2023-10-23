package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig :nodoc:
func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil && Env() != "test" {
		logrus.Warningf("%v", err)
	}
}

// Env :nodoc:
func Env() string {
	return viper.GetString("env")
}

// HTTPPort :nodoc:
func HTTPPort() int {
	return viper.GetInt("http_port")
}

// LogLevel :nodoc:
func LogLevel() string {
	return viper.GetString("log_level")
}
