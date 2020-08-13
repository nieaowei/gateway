FROM golang:1.14

RUN ls -a

RUN go get -d -v -t ./...
RUN go build -v .

ENTRYPOINT ["./main"]