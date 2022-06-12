package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

func BuildConfig() Config {
	vpr, err := ConfigureViper("config/config.yaml", "app", "yaml")
	if err != nil {
		panic("failed to create viper instance")
	}

	return ParseAll(vpr)
}

func ParseAll(vpr *viper.Viper) Config {
	value := *new(Config)
	if err := vpr.Unmarshal(&value); err != nil {
		panic("failed to unmarshal full config")
	}
	return value
}

func ConfigureViper(path, envPrefix, configType string) (*viper.Viper, error) {
	vpr := viper.New()
	defaultConfig := bytes.NewReader([]byte{})
	vpr.SetConfigType("yaml")
	if err := vpr.MergeConfig(defaultConfig); err != nil {
		return nil, err
	}

	// Override config
	vpr.SetConfigFile(path)
	vpr.SetConfigType(configType)
	if err := vpr.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return nil, err
		}
		// dont return error if file is missing. overwrite file is optional
	}

	// Overrwire env variables
	vpr.AutomaticEnv()
	vpr.SetEnvPrefix(envPrefix)
	vpr.AddConfigPath(".")
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Manually set all config values
	for _, key := range vpr.AllKeys() {
		val := vpr.Get(key)
		vpr.Set(key, val)
	}

	return vpr, nil
}
