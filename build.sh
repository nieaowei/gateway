CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags v1.0 -o bin/gateway-drawin ./
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags v1.0 -o bin/gateway-linux ./
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags v1.0 -o bin/gateway-windows ./
cp -r conf ./bin