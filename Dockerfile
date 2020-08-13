FROM golang:1.14

WORKDIR /go/src/app

COPY . .

RUN go get -d -v -t ./...
RUN go build -v .

EXPOSE 8880

ENTRYPOINT["./gateway","-conf=pro"]