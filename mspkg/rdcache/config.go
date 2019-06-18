package rdcache

type Config struct {
	Host     string
	Port     int64
	Password string
}

type Configurable interface {
	GetRedisConfig() *Config
}
