package web

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Addr            string        `bson:"addr,omitempty" json:"addr,omitempty"`
	ShutdownTimeout time.Duration `bson:"shutdownTimeout,omitempty" json:"shutdownTimeout,omitempty"`
	ReadTimeout     time.Duration `bson:"readTimeout,omitempty" json:"readTimeout,omitempty"`
	WriteTimeout    time.Duration `bson:"writeTimeout,omitempty" json:"writeTimeout,omitempty"`
}

var keyConfig = "web"

func NewConfig(viper *viper.Viper) (*Config, error) {
	data := viper.Get(keyConfig)

	if data == nil {
		return nil, errors.New("WEB no data in config")
	}

	config := Config{}
	err := mapstructure.Decode(data, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) ConfigIsFilled() bool {
	return c.Addr != ""
}
