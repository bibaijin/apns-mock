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
	errLogger    *log.Logger
	infoLogger   *log.Logger
	RequestIDKey = requestIDKeyType("RequestIDKey")
)

func init() {
	errLogger = log.New(os.Stderr, "", LogFlag)
	infoLogger = log.New(os.Stdout, "", LogFlag)
}

// Fatalf 打印错误日志并退出
func Fatalf(ctx context.Context, format string, v ...interface{}) {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		format = fmt.Sprintf("RequestID: %s %s", requestID, format)
	}

	errLogger.Output(2, fmt.Sprintf(format, v...))
}

// Errorf 打印错误日志并退出
func Errorf(ctx context.Context, format string, v ...interface{}) {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		format = fmt.Sprintf("RequestID: %s %s", requestID, format)
	}

	errLogger.Output(2, fmt.Sprintf(format, v...))
}

// Infof 打印错误日志并退出
func Infof(ctx context.Context, format string, v ...interface{}) {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		format = fmt.Sprintf("RequestID: %s %s", requestID, format)
	}

	infoLogger.Output(2, fmt.Sprintf(format, v...))
}
