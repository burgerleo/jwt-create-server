package conf

import "github.com/go-ini/ini"

var Conf *Config

type Config struct {
	BaseConf
	LogConf
}

// BaseConf inlclude deatils server components
type BaseConf struct {
	HttpPort string `ini:"HttpPort"` // http port
	Env      string `ini:"Env"`      // 運行環境
}

// LogConf record log to specific folder
type LogConf struct {
	LogPath  string `ini:"LogPath"`
	LogLevel string `ini:"LogLevel"`
}

type JWTConf struct {
	JWTSecret        string `ini:"JWTSecret"`
	JWTExpiredMinute int    `ini:"JWTExpiredMinute"`
}

func InitConfig(confPath *string) (*Config, error) {
	Conf = new(Config)
	if err := ini.MapTo(Conf, *confPath); err != nil {
		return nil, err
	}
	return Conf, nil
}
