package router

import (
	"github.com/Deansquirrel/goAgentWin/agent"
	"github.com/Deansquirrel/goAgentWin/object"
	"github.com/kataras/iris"
)

func AddWebPartService(app *iris.Application) {
	service := app.Party("service", serviceHandler)
	service.Post("/restart", serviceRestartHandler)
}

func serviceHandler(ctx iris.Context) {
	ctx.Next()
}

type serviceRestartInfo struct {
	ServiceName string `json:"servicename"`
}

func serviceRestartHandler(ctx iris.Context) {
	var info serviceRestartInfo
	err := ctx.ReadJSON(&info)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ws := agent.NewWinService(info.ServiceName)
	err = ws.Restart()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ro := object.NewOKReturn()
	writeResponse(ctx, ro)
	return
}
