package main

import (
	"context"
	"net/http"
)

// Ping 用于健康检查
func Ping(_ context.Context, w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}
