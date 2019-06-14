package conf

type MsConfig struct {
	LogLevel string
}

func (mc MsConfig) GetLogLevel() string {
	return mc.LogLevel
}

type MsConfigurable interface {
	GetLogLevel() string
}
