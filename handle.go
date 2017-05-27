package main

import (
	"context"
	"net/http"

	"github.com/bibaijin/apns-mock/log"
)

type handler func(context.Context, http.ResponseWriter, *http.Request)

func handle(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(CTX, log.RequestIDKey, newRequestID())
		log.Infof(ctx, "Receive a request, URL: %s, Method: %s, RemoteAddr: %s.", r.URL, r.Method, r.RemoteAddr)

		h(ctx, w, r)

		log.Infof(ctx, "Response is sent, URL: %s, Method: %s, RemoteAddr: %s.", r.URL, r.Method, r.RemoteAddr)
	}
}
