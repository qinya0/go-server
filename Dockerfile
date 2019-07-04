FROM hub.iot.chinamobile.com/library/golang:1.9.2

WORKDIR $GOPATH/bin

COPY ./go-server .
RUN chmod 777 ./go-server

# environment variables
CMD ["./go-server"]
