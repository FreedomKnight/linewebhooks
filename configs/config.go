package configs

import (
	// "log"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper = viper.New()

func Get() *viper.Viper {
    config.AddConfigPath("configs")
    config.AutomaticEnv() 
    config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    return config
}
