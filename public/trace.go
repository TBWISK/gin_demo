package public

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

//Trace 跟踪
type Trace struct {
	TraceID     string
	SpanID      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}

//TraceContext context
type TraceContext struct {
	Trace
	CSpanID string
}

//NewTrace 跟踪
func NewTrace() *TraceContext {
	trace := &TraceContext{}
	trace.TraceID = GetTraceID()
	trace.SpanID = NewSpanID()
	return trace
}

// NewSpanID 新跟踪id
func NewSpanID() string {
	timestamp := uint32(time.Now().Unix())
	ipToLong := binary.BigEndian.Uint32(LocalIP.To4())
	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf("%08x", ipToLong^timestamp))
	b.WriteString(fmt.Sprintf("%08x", rand.Int31()))
	return b.String()
}

//GetTraceID 获取跟踪id
func GetTraceID() (traceID string) {
	return calcTraceID(LocalIP.String())
}

func calcTraceID(ip string) (traceID string) {
	now := time.Now()
	timestamp := uint32(now.Unix())
	timeNano := now.UnixNano()
	pid := os.Getpid()

	b := bytes.Buffer{}
	netIP := net.ParseIP(ip)
	if netIP == nil {
		b.WriteString("00000000")
	} else {
		b.WriteString(hex.EncodeToString(netIP.To4()))
	}
	b.WriteString(fmt.Sprintf("%08x", timestamp&0xffffffff))
	b.WriteString(fmt.Sprintf("%04x", timeNano&0xffff))
	b.WriteString(fmt.Sprintf("%04x", pid&0xffff))
	b.WriteString(fmt.Sprintf("%06x", rand.Int31n(1<<24)))
	b.WriteString("b0") // 末两位标记来源,b0为go

	return b.String()
}
