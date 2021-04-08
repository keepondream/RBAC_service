package utils

import (
	"database/sql"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Option struct {
	Code       string      `json:"errorCode"`
	DevelopMsg string      `json:"developMsg,omitempty"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	Err        error       `json:"-"`
}

type ModOption func(option *Option)

func WithData(data interface{}) ModOption {
	return func(option *Option) {
		option.Data = data
	}
}

func WithErr(err error) ModOption {
	return func(option *Option) {
		option.Err = err
	}
}

func Success(c *gin.Context, modOptions ...ModOption) {
	data := []interface{}{}

	option := Option{
		Code:       Err_No.ErrCode,
		DevelopMsg: "",
		Msg:        Err_No.ErrMessage,
		Data:       data,
	}

	for _, fn := range modOptions {
		fn(&option)
	}

	c.AbortWithStatusJSON(http.StatusOK, option)
}

func Failed(c *gin.Context, modOptions ...ModOption) {
	c.GetBool("is_debug")
	data := []interface{}{}
	httpCode := http.StatusBadRequest

	option := Option{
		Code:       Err_Failed.ErrCode,
		DevelopMsg: "",
		Msg:        Err_Failed.ErrMessage,
		Data:       data,
	}

	for _, fn := range modOptions {
		fn(&option)
	}

	if option.Err != nil {
		if c.GetBool("debug") {
			option.DevelopMsg = option.Err.Error()
		}

		if e, ok := option.Err.(Errno); ok {
			option.Code = e.ErrCode
		} else if _, ok := option.Err.(validator.ValidationErrors); ok {
			option.Code = Err_HTTP_Param.ErrCode
		} else if e, ok := option.Err.(net.Error); ok && e.Timeout() {
			httpCode = http.StatusGatewayTimeout
		} else if e == sql.ErrNoRows {
			option.Code = Err_No_Row.ErrCode
		} else {
			httpCode = http.StatusInternalServerError
			option.Code = Err_Failed.ErrCode
		}
	}

	c.AbortWithStatusJSON(httpCode, option)
}
