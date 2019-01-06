FROM golang:latest

MAINTAINER Card "445864742@qq.com"

WORKDIR $GOPATH/src/resource-backend
ADD . $GOPATH/src/resource-backend

RUN go build .

EXPOSE 8888

ENTRYPOINT ["./resource-backend"]