package object

const (
	Success     = 0
	CommonError = -1
)

type responseObject struct {
	//错误编码
	ErrCode int `json:"errcode"`
	//错误描述
	ErrMsg string `json:"errmsg"`
	//数据
	Data interface{} `json:"data"`
}

type responseObjectWithNoData struct {
	//错误编码
	ErrCode int `json:"errcode"`
	//错误描述
	ErrMsg string `json:"errmsg"`
}

func NewOKReturn() *responseObjectWithNoData {
	return &responseObjectWithNoData{
		ErrCode: Success,
		ErrMsg:  "OK",
	}
}

func NewErrReturn(err error) *responseObjectWithNoData {
	return &responseObjectWithNoData{
		ErrCode: CommonError,
		ErrMsg:  err.Error(),
	}
}

func NewDataReturn(v interface{}) *responseObject {
	return &responseObject{
		ErrCode: Success,
		ErrMsg:  "OK",
		Data:    v,
	}
}
