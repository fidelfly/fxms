package conf

const (
	StageDevelopment = "development"
	//StageProduction
	StageProduction = "production"
)

type MsConfig struct {
	Log   LogConfig
	Stage string
}

type LogConfig struct {
	Level string
	Path  string
}

func (lc LogConfig) GetLevel() string {
	return lc.Level
}

func (mc MsConfig) GetLog() LogConfig {
	return mc.Log
}

func (mc MsConfig) GetStage() string {
	return mc.Stage
}

type MsConfigurable interface {
	GetLog() *LogConfig
	GetStage() string
}
