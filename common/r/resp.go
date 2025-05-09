package r

type Resp struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Data(code int, data any, msg ...string) (Resp, error) {
	var msgStr = "success"
	if len(msg) > 0 {
		msgStr = msg[0]
	}
	return Resp{
		Code: code,
		Data: data,
		Msg:  msgStr,
	}, nil
}

func Err(err error) (Resp, error) {
	return Resp{}, err
}

func Code(code int) (Resp, error) {
	return Resp{
		Code: code,
	}, nil
}
