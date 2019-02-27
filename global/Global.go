package global

import (
	"context"
	"github.com/Deansquirrel/goAgentWin/config"
)

const (
	Version = "0.0.0 Build20190101"
)

var SysConfig *config.SysConfig
var Ctx context.Context
var Cancel func()
