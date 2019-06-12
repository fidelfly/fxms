package conf

type MsConfig struct {
	Version  string
	LogLevel string
}

func (mc MsConfig) GetVersion() string {
	return mc.Version
}

func (mc MsConfig) GetLogLevel() string {
	return mc.LogLevel
}

type MsConfigurable interface {
	GetVersion() string
	GetLogLevel() string
}
