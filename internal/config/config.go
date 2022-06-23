package config

import (
	"fmt"
	"time"

	"log-agent/internal/tail"

	"gopkg.in/ini.v1"
)

type Config struct {
	AppName    string       `ini:"app_name"`
	Mongo      *MongoConfig `ini:"mongodb"`
	Jwt        *Jwt         `ini:"jwt"`
	HttpConfig *HttpConfig  `ini:"http"`
	Handler    map[string]*tail.Config
	Start      time.Time
}

type HttpConfig struct {
	AllowOrigins  string `ini:"allow_origins"`
	AllowRegister bool   `ini:"allow_register"`
}

func Load(filename string) (*Config, error) {

	file, err := ini.Load(filename)
	if err != nil {
		return nil, fmt.Errorf("ini load error: %s", err.Error())
	}

	config := &Config{
		Handler: make(map[string]*tail.Config),
		Start:   time.Now(),
	}
	if err := file.MapTo(config); err != nil {
		return nil, fmt.Errorf("ini map to config error: %s", err)
	}

	for _, section := range file.Section("log").ChildSections() {
		tailConfig := &tail.Config{}
		if err := section.StrictMapTo(tailConfig); err != nil {
			return nil, fmt.Errorf("ini tail handle error: %s", err)
		}

		if tailConfig.Path == "" {
			continue
		}

		config.Handler[section.Name()] = tailConfig
	}

	return config, nil
}
