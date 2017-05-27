# apns-mock

Apple Push Notification Service 的 mock。

## 开发

### 本地运行

```
git clone https://github.com/bibaijin/apns-mock.git
cd apns-mock/
go install
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host apns-mock.lain.local  # 生成自签名证书（`key.pem` 和 `cert.pem`）
```

### 部署到 LAIN 集群

apns-mock 是一个 [LAIN](https://github.com/laincloud/lain) 应用，可以部署到 LAIN 集群：

```
lain build
lain tag ${LAIN-cluster}
lain push ${LAIN-cluster}
lain deploy ${LAIN-cluster}
```
