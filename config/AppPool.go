package config

import "strings"

type AppPool struct {
	Path         string `toml:"path"`
	StartDelay   int    `toml:"restartDelay"`
	StartTimeout int    `toml:"restartTimeout"`
}

func (ap *AppPool) FormatConfig() {
	ap.Path = strings.Trim(ap.Path, " ")
	//设置默认值
	if ap.Path == "" {
		ap.Path = "C:\\Windows\\System32\\INetSrv\\AppCmd.exe"
	}
	if ap.StartDelay < 15 {
		ap.StartDelay = 15
	}
	//默认300秒超时
	if ap.StartTimeout == 0 {
		ap.StartTimeout = 300
	}
	//如果配置时间小于60，强制改为60
	if ap.StartTimeout < 60 {
		ap.StartTimeout = 60
	}
}
