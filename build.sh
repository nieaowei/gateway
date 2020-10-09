rm -rf ./bin/*
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags v$1 -o bin/gateway-server-drawin ./
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags v$1 -o bin/gateway-server-linux ./
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags v$1 -o bin/gateway-server-windows ./
cp -r conf ./bin
cp run.sh ./bin
cp -r docs ./bin
cp -r templates ./bin
cp Dockerfile ./bin