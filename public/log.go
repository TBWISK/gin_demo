package public

import (
	"context"
	"fmt"
	"strings"

	"github.com/TBWISK/goconf"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

// 通用DLTag常量定义
const (
	DLTagUndefind      = "_undef"
	DLTagMySQLFailed   = "_com_mysql_failure"
	DLTagRedisFailed   = "_com_redis_failure"
	DLTagMySQLSuccess  = "_com_mysql_success"
	DLTagRedisSuccess  = "_com_redis_success"
	DLTagThriftFailed  = "_com_thrift_failure"
	DLTagThriftSuccess = "_com_thrift_success"
	DLTagHTTPSuccess   = "_com_http_success"
	DLTagHTTPFailed    = "_com_http_failure"
	DLTagTCPFailed     = "_com_tcp_failure"
	DLTagRequestIn     = "_com_request_in"
	DLTagRequestOut    = "_com_request_out"
)

const (
	_dlTag          = "dltag"
	_traceID        = "traceid"
	_spanID         = "spanid"
	_childSpanID    = "cspanid"
	_dlTagBizPrefix = "_com_"
	_dlTagBizUndef  = "_com_undef"
)

//InitLog  日志初始化
func InitLog(path string) {
	logger := goconf.NewLoger(path)
	sugar = logger.Sugar()
}

//Info 打印
func Info(args ...interface{}) {
	sugar.Info(args)
}

//Error 打印
func Error(args ...interface{}) {
	sugar.Error(args)
}

//Warn 打印
func Warn(args ...interface{}) {
	sugar.Warn(args)
}

//Debug 打印
func Debug(args ...interface{}) {
	sugar.Debug(args)
}

//TagError tag日志打印
func TagError(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceID] = trace.TraceID
	m[_childSpanID] = trace.CSpanID
	m[_spanID] = trace.SpanID
	sugar.Error(parseParams(m))
}

//TagInfo tag日志打印
func TagInfo(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceID] = trace.TraceID
	m[_childSpanID] = trace.CSpanID
	m[_spanID] = trace.SpanID
	sugar.Info(parseParams(m))
}

//TagDebug tag日志打印
func TagDebug(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceID] = trace.TraceID
	m[_childSpanID] = trace.CSpanID
	m[_spanID] = trace.SpanID
	sugar.Debug(parseParams(m))
}

//TagWarn tag日志打印
func TagWarn(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceID] = trace.TraceID
	m[_childSpanID] = trace.CSpanID
	m[_spanID] = trace.SpanID
	sugar.Warn(parseParams(m))
}

//ContextWarning 错误日志
func ContextWarning(c context.Context, dltag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*TraceContext)
	if !ok {
		traceContext = NewTrace()
	}
	TagWarn(traceContext, dltag, m)
}

//ContextError 错误日志
func ContextError(c context.Context, dltag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*TraceContext)
	if !ok {
		traceContext = NewTrace()
	}
	TagError(traceContext, dltag, m)
}

//ContextNotice 普通日志
func ContextNotice(c context.Context, dltag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*TraceContext)
	if !ok {
		traceContext = NewTrace()
	}
	TagInfo(traceContext, dltag, m)
}

//ComLogWarning 错误日志
func ComLogWarning(c *gin.Context, dltag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	TagError(traceContext, dltag, m)
}

//ComLogNotice 普通日志
func ComLogNotice(c *gin.Context, dltag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	TagInfo(traceContext, dltag, m)
}

//GetGinTraceContext 从gin的Context中获取数据
func GetGinTraceContext(c *gin.Context) *TraceContext {
	// 防御
	if c == nil {
		return NewTrace()
	}
	traceContext, exists := c.Get("trace")
	if exists {
		if tc, ok := traceContext.(*TraceContext); ok {
			return tc
		}
	}
	return NewTrace()
}

//GetTraceContext 从Context中获取数据
func GetTraceContext(c context.Context) *TraceContext {
	if c == nil {
		return NewTrace()
	}
	traceContext := c.Value("trace")
	if tc, ok := traceContext.(*TraceContext); ok {
		return tc
	}
	return NewTrace()
}

// 校验dltag合法性
func checkDLTag(dltag string) string {
	if strings.HasPrefix(dltag, _dlTagBizPrefix) {
		return dltag
	}

	if strings.HasPrefix(dltag, "_com_") {
		return dltag
	}

	if dltag == DLTagUndefind {
		return dltag
	}
	return dltag
}

//map格式化为string
func parseParams(m map[string]interface{}) string {
	var dltag string = "_undef"
	if _dltag, _have := m["dltag"]; _have {
		if _Val, _Ok := _dltag.(string); _Ok {
			dltag = _Val
		}
	}
	for _key, _val := range m {
		if _key == "dltag" {
			continue
		}
		dltag = dltag + "||" + fmt.Sprintf("%v=%+v", _key, _val)
	}
	dltag = strings.Trim(fmt.Sprintf("%q", dltag), "\"")
	return dltag
}
