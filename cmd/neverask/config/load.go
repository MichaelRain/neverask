package config

import (
	"github.com/MichaelRain/neverask/internal/platform/web"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	// Web server
	Web *web.Config
}

var v *viper.Viper

func Load() (*Config, error) {
	var err error

	file := pflag.StringP("config", "c", "", "path to config file")
	pflag.Parse()

	if file != nil && len(*file) > 0 {
		v, err = loadFromFile(*file)
	} else {
		v, err = loadFromBasePath()
	}

	if err != nil {
		return nil, err
	}

	c := Config{}

	if c.Web, err = web.NewConfig(v); err != nil {
		return nil, err
	}

	if c.Redis, err = redis.NewConfig(v); err != nil {
		return nil, err
	}

	return &c, nil
}

func loadFromFile(file string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(file)

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return v, nil

}

func loadFromBasePath() (*viper.Viper, error) {
	return loadFromFile("/etc/rw3/config.yaml")
}
