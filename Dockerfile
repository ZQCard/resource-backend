FROM golang:latest

MAINTAINER Card "445864742@qq.com"

WORKDIR $GOPATH/src/resource_backend
COPY . $GOPATH/src/resource_backend
RUN go build .

EXPOSE 8888

ENTRYPOINT ["./resource_backend"]
