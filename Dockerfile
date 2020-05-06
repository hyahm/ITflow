FROM golang
MAINTAINER hyahm.com
WORKDIR /root/
COPY go .
ENV GOPATH /root
RUN go get github.com/go-redis/redis
