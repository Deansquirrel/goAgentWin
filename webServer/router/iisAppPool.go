package router

import (
	"github.com/Deansquirrel/goAgentWin/agent"
	"github.com/Deansquirrel/goAgentWin/object"
	"github.com/kataras/iris"
)

func AddWebPartIISAppPool(app *iris.Application) {
	iisAppPool := app.Party("iisapppool", iisAppPoolHandler)
	iisAppPool.Post("/restart", iisAppPoolRestartHandler)
	iisAppPool.Post("/isrunning", iisAppPoolIsRunningHandler)
	iisAppPool.Post("/stop", iisAppPoolStopHandler)
}

func iisAppPoolHandler(ctx iris.Context) {
	ctx.Next()
}

type iisAppPoolRestartInfo struct {
	AppPoolName string `json:"apppoolname"`
}

func iisAppPoolRestartHandler(ctx iris.Context) {
	var info iisAppPoolRestartInfo
	err := ctx.ReadJSON(&info)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ws := agent.NewIISAppPool(info.AppPoolName)
	err = ws.Restart()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ro := object.NewOKReturn()
	writeResponse(ctx, ro)
	return

}

type runningState struct {
	IsRunning bool `toml:"isrunning"`
}

func iisAppPoolIsRunningHandler(ctx iris.Context) {
	var info iisAppPoolRestartInfo
	err := ctx.ReadJSON(&info)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ws := agent.NewIISAppPool(info.AppPoolName)
	b, err := ws.IsRunning()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ro := object.NewDataReturn(runningState{
		IsRunning: b,
	})
	writeResponse(ctx, ro)
	return
}

func iisAppPoolStopHandler(ctx iris.Context) {
	var info iisAppPoolRestartInfo
	err := ctx.ReadJSON(&info)
	if err != nil {

		ctx.StatusCode(iris.StatusBadRequest)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ws := agent.NewIISAppPool(info.AppPoolName)
	err = ws.Stop()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ro := object.NewErrReturn(err)
		writeResponse(ctx, ro)
		return
	}
	ro := object.NewOKReturn()
	writeResponse(ctx, ro)
	return

}
