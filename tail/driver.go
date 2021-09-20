package tail

type LogDriver func(data interface{}, conf *Config) error

var logDriver = map[string]LogDriver{}

func RegisterDriver(name string, driver LogDriver) {
	logDriver[name] = driver
}
