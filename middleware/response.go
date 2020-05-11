package middleware

import (
	"encoding/json"
	"tbwisk/public"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000

	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
)

type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
	TraceID   interface{}  `json:"trace_id"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*public.TraceContext)
	TraceID := ""
	if traceContext != nil {
		TraceID = traceContext.TraceID
	}

	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "", TraceID: TraceID}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*public.TraceContext)
	TraceID := ""
	if traceContext != nil {
		TraceID = traceContext.TraceID
	}

	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data, TraceID: TraceID}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
