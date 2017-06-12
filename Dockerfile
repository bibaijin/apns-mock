FROM golang:1.8

LABEL maintainer "bibaijin@gmail.com"

WORKDIR $GOPATH/src/github.com/bibaijin/apns-mock
COPY . .
RUN go install
RUN ln -s $GOPATH/bin/apns-mock /apns-mock
RUN ln -s $GOPATH/src/github.com/bibaijin/apns-mock/cert.pem /cert.pem
RUN ln -s $GOPATH/src/github.com/bibaijin/apns-mock/key.pem /key.pem

EXPOSE 8443
CMD /apns-mock -port=8443 -certfile=/cert.pem -keyfile=/key.pem