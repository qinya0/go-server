FROM golang:1.12.13

WORKDIR $GOPATH/bin

COPY ./go-server .
RUN chmod 777 ./go-server

# environment variables
CMD ["./go-server"]
