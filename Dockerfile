FROM golang:latest

MAINTAINER Card "445864742@qq.com"

WORKDIR $GOPATH/src/resource-backend
ADD . $GOPATH/src/resource-backend
#由于暂时无法解决net包的下载问题,只能在本地编译过后生成镜像
#RUN go get github.com/gin-gonic/gin
#RUN go get github.com/Unknwon/com
#RUN go get github.com/Unknwon/goconfig
#RUN go get github.com/dgrijalva/jwt-go
#RUN go get github.com/go-ozzo/ozzo-validation
#RUN go get github.com/go-ozzo/ozzo-validation/is
#RUN go get github.com/go-sql-driver/mysql
#RUN go get github.com/jinzhu/gorm
#RUN go get github.com/pkg/errors
#RUN go get gopkg.in/gomail.v2
#官方包特殊处理
#RUN mkdir -p $GOPATH/src/golang.org
#RUN cd $GOPATH/src/golang.org
#RUN git clone https://github.com/golang/net.git
#RUN cd $GOPATH/src/resource-backend
#RUN go get github.com/qiniu/api.v7/auth/qbox
#RUN go get github.com/qiniu/api.v7/storage
#RUN go build .

EXPOSE 8888

ENTRYPOINT ["./resource-backend"]