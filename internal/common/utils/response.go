package utils

type Response struct {
	// 错误码 , 描述对应OpenAPI中的ErrCode
	Code *string `json:"code,omitempty"`

	// 错误字段
	Field *string `json:"field,omitempty"`

	// 错误信息
	Msg string `json:"msg"`

	// 错误
	Err error `json:"-"`

	// 响应体内容
	Data interface{} `json:"-"`
}

type ModResponse func(resp *Response)

func WithCode(code string) ModResponse {
	return func(resp *Response) {
		resp.Code = &code
	}
}

func WithField(field string) ModResponse {
	return func(resp *Response) {
		resp.Field = &field
	}
}

func WithMsg(msg string) ModResponse {
	return func(resp *Response) {
		resp.Msg = msg
	}
}

func WithByteForMsg(body []byte) ModResponse {
	return func(resp *Response) {
		resp.Msg = string(body)
	}
}

func WithError(err error) ModResponse {
	return func(resp *Response) {
		resp.Err = err
	}
}

func WithData(data interface{}) ModResponse {
	return func(resp *Response) {
		resp.Data = data
	}
}
