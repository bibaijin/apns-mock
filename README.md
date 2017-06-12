# apns-mock

Apple Push Notification Service 的 mock。

## 运行

```
docker-compose pull
docker-compose up -d
```

## 其他

`cert.pem` 和 `key.pem` 是使用 golang 的库生成的：

```
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host apns-mock.test  # 生成自签名证书（`key.pem` 和 `cert.pem`）
```
