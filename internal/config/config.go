package config

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
	"log-agent/internal/tail"
)

type Config struct {
	AppName string       `ini:"app_name"`
	Mongo   *MongoConfig `ini:"mongodb"`
	Handler map[string]*tail.Config
	Start   time.Time
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
