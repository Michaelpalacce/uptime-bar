package configuration

import (
	"fmt"

	"github.com/Michaelpalacce/uptime-bar/pkgs/status"
	"github.com/spf13/viper"
)

type Configuration struct {
	HttpStatuses []status.HttpStatus `mapstructure:"httpStatuses"`
}

// LoadConfiguration will attempt to load the configuration from the config path.
func LoadConfiguration() (*Configuration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// viper.AddConfigPath("$HOME/.config/uptime-bar")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var configuration Configuration

	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, fmt.Errorf("error while unmarshaling HttpStatuses configuration. Err was: %w", err)
	}

	return &configuration, nil
}
