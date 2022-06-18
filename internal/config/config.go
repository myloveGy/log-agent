package config

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
	"log-agent/internal/tail"
)

type Config struct {
	AppName string         `ini:"app_name"`
	Mongo   *MongoConfig   `ini:"mongodb"`
	Redis   *RedisConfig   `ini:"redis"`
	Tail    string         `ini:"tail"`
	Handles []*tail.Config `json:"handle"`
}

func Load(filename string) (*Config, error) {
	var (
		file *ini.File
		err  error
		conf = &Config{}
	)

	file, err = ini.Load(filename)
	if err != nil {
		return nil, fmt.Errorf("ini load error: %s", err.Error())
	}

	if err := file.MapTo(conf); err != nil {
		return nil, fmt.Errorf("ini map to config error: %s", err)
	}

	handles := strings.Split(conf.Tail, ",")
	for _, v := range handles {
		tailConfig := &tail.Config{}
		if err := file.Section(v).StrictMapTo(tailConfig); err != nil {
			return nil, fmt.Errorf("ini tail handle error: %s", err)
		}

		conf.Handles = append(conf.Handles, tailConfig)
	}

	return conf, nil
}
