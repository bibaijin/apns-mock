package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
	"gitlab.yxapp.in/common-services/apns-mock/log"
)

// PostDeviceToken 向设备发送信息
func PostDeviceToken(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	deviceToken := bone.GetValue(r, "deviceToken")

	if deviceToken == "" {
		log.Errorf(ctx, "deviceToken is empty.")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "deviceToken is empty")
	}

	w.WriteHeader(http.StatusOK)
}
