package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bibaijin/apns-mock/log"
	"github.com/go-zoo/bone"
)

var (
	port     = flag.Int("port", 8443, "The port to listen")
	certfile = flag.String("certfile", "cert.pem", "The cert file")
	keyfile  = flag.String("keyfile", "key.pem", "The key file")
	// CTX 表示顶级环境
	CTX = context.Background()
)

func init() {
	flag.Parse()
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)

	mux := bone.New()
	mux.PostFunc("/3/device/:deviceToken", handle(PostDeviceToken))
	mux.GetFunc("/ping", handle(Ping))

	server := http.Server{
		Addr:    ":" + strconv.Itoa(*port),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServeTLS(*certfile, *keyfile); err != nil {
			log.Errorf(CTX, "server.ListenAndServeTLS() failed, error: %s.", err)
		}
	}()
	defer func() {
		if err := server.Shutdown(CTX); err != nil {
			log.Errorf(CTX, "server.Shutdown() failed, error: %s.", err)
		}
	}()
	log.Infof(CTX, "server.ListenAndServeTLS()..., port: %d.", *port)

	<-quit
	log.Infof(CTX, "Shutting down...")
}
