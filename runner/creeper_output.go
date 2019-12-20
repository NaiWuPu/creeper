package runner

type CreeperOutput struct {
	Code   string      `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Custom interface{} `json:"custom"`
	Token  string      `json:"token"`
}

const (
	SuccessCode = "100000"
	ErrorCode   = "100001"
)

const (
	SuccessMsg = "操作成功"
	ErrorMsg   = "系统异常"
)

func (ot *CreeperOutput) SuccessOutput(data interface{}, msg string) *CreeperOutput {
	ot.Code = SuccessCode
	if msg == "" {
		ot.Msg = SuccessMsg
	} else {
		ot.Msg = msg
	}
	ot.Data = data
	return ot
}

func (ot *CreeperOutput) ErrorOutput(msg string) *CreeperOutput {
	ot.Code = ErrorCode
	if msg == "" {
		ot.Msg = ErrorMsg
	} else {
		ot.Msg = msg
	}
	return ot
}
