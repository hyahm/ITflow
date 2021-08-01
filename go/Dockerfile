FROM golang
WORKDIR /data
COPY src bug/src
COPY pri.key entrypoint.sh bug/
ENV GOPATH=/data/bug
WORKDIR /data/bug
RUN go build main.go && \
	chmod +x entrypoint.sh
EXPOSE 10001
VOLUME "/data/log" "/data/bugimg" "/share"
ENTRYPOINT ["./entrypoint.sh"]
CMD ["./main"]
