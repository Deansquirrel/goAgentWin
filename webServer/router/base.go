package router

import (
	"github.com/Deansquirrel/goAgentWin/global"
	"github.com/Deansquirrel/goAgentWin/object"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
)

func AddWebPartBase(app *iris.Application) {
	app.Get("/version", versionHandler)
	app.Get("/test", testOKHandler)
	app.Get("/err", testErrHandler)
}

type versionInfo struct {
	Version string `json:"version"`
}

func versionHandler(ctx iris.Context) {
	v := versionInfo{
		Version: global.Version,
	}
	ro := object.NewDataReturn(v)
	rs, err := goToolCommon.GetJsonStr(ro)
	if err != nil {

	} else {
		_, _ = ctx.WriteString(string(rs))
	}
}

func testOKHandler(ctx iris.Context) {
	ro := object.NewOKReturn()
	writeResponse(ctx, ro)
}

func testErrHandler(ctx iris.Context) {
	ro := object.NewErrReturn(errors.New("Test Error"))
	writeResponse(ctx, ro)
}

//func clientInfoInfoHandler(ctx iris.Context) {
//	var info object.ClientInfo
//	err := ctx.ReadJSON(&info)
//	if err != nil {
//		ctx.StatusCode(iris.StatusBadRequest)
//		_, _ = ctx.WriteString(GetErrReturn(err.Error()))
//		log.Warn(err.Error())
//		return
//	}
//
//	b, err := json.Marshal(info)
//	if err != nil {
//		log.Error(err.Error())
//	} else {
//		log.Debug(string(b))
//	}
//	_, _ = ctx.WriteString(GetMsgReturn("OK"))
//}
