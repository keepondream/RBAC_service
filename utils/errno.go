package utils

type Errno struct {
	ErrCode    string
	ErrMessage string
}

func (err Errno) Error() string {

	return err.ErrMessage
}

func (err Errno) WithErr(e error) error {
	err.ErrMessage = e.Error()

	return err
}

var (
	Err_HTTP_Param         = &Errno{"400_1", ""} // param parse
	Err_HTTP_Query         = &Errno{"400_2", ""} // query parse
	Err_HTTP_Json          = &Errno{"400_3", ""} // json parse
	Err_HTTP_Authorization = &Errno{"400_4", ""} // not found Authorization

	Err_No          = &Errno{"000000", ""} // 没有错误
	Err_Failed      = &Errno{"999999", ""} // 未定义的错误
	Err_Down_Stream = &Errno{"888888", ""} // 下游服务错误
	Err_No_Row      = &Errno{"777777", ""} // 数据不存在
	Err_Duplicate   = &Errno{"666666", ""} // 数据重复
)
