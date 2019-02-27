package router

import (
	"fmt"
	"github.com/Deansquirrel/goAgentWin/object"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/kataras/iris"
)

const (
	FormatErr = "{\"errcode\":%d,\"errmsg\":\"%s\"}\""
)

func getTranErrReturn(msg string) string {
	return fmt.Sprintf(FormatErr, object.CommonError, msg)
}

func writeResponse(ctx iris.Context, v interface{}) {
	rs, err := goToolCommon.GetJsonStr(v)
	if err != nil {
		_, _ = ctx.WriteString(getTranErrReturn(err.Error()))
	} else {
		_, _ = ctx.WriteString(string(rs))
	}
}
