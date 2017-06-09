package log

import (
	"context"
	"fmt"
	"log"
	"os"
)

const (
	// LogFlag 控制日志的前缀
	LogFlag = log.LstdFlags | log.Lmicroseconds | log.Lshortfile
)

type requestIDKeyType string

var (
	errLogger  *log.Logger
	infoLogger *log.Logger
	// RequestIDKey 表示请求 ID 的键名
	RequestIDKey = requestIDKeyType("RequestIDKey")
)

func init() {
	errLogger = log.New(os.Stderr, "", LogFlag)
	infoLogger = log.New(os.Stdout, "", LogFlag)
}

// Fatalf 打印错误日志并退出
func Fatalf(ctx context.Context, format string, v ...interface{}) {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		format = fmt.Sprintf("FATAL RequestID: %s %s", requestID, format)
	} else {
		format = fmt.Sprintf("FATAL %s", format)
	}

	errLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Errorf 打印错误日志并退出
func Errorf(ctx context.Context, format string, v ...interface{}) {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		format = fmt.Sprintf("ERROR RequestID: %s %s", requestID, format)
	} else {
		format = fmt.Sprintf("ERROR %s", format)
	}

	errLogger.Output(2, fmt.Sprintf(format, v...))
}

// Infof 打印错误日志并退出
func Infof(ctx context.Context, format string, v ...interface{}) {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		format = fmt.Sprintf("INFO RequestID: %s %s", requestID, format)
	} else {
		format = fmt.Sprintf("INFO %s", format)
	}

	infoLogger.Output(2, fmt.Sprintf(format, v...))
}
