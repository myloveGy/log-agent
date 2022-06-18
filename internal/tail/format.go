package tail

import (
	"encoding/json"
)

type FormatHandler func(text string, conf *Config) (map[string]interface{}, error)

var formatHandles = map[string]FormatHandler{}

func jsonFormat(text string, _ *Config) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &data)
	return data, err
}

func RegisterFormat(format string, fn FormatHandler) {
	formatHandles[format] = fn
}

func init() {
	RegisterFormat("json", jsonFormat)
}
