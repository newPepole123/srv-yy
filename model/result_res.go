package model

type ResultRes struct {
	Msg  string      `json:"errMsg"`
	Code int         `json:"statusCode"`
	Data interface{} `json:"data"`
}

func CreateResultRes() *ResultRes {
	return &ResultRes{}
}

func (this *ResultRes) SetErrorData(code int, msg string) *ResultRes {
	this.Code = code
	this.Msg = msg
	return this
}

func (this *ResultRes) SetSuccessData(data interface{}) *ResultRes {
	this.Code = 200
	// this.Msg = msg
	this.Data = data
	return this
}
