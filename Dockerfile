FROM golang:1.14

WORKDIR /go/src/app

COPY . .

EXPOSE 8880
EXPOSE 8800

ENTRYPOINT ["./gateway-server-linux"]