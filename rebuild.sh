set -e

protoc -I proto proto/services/helloworld/helloworld.proto --go_out=plugins=grpc:./gen
go build -o server/main server/main.go
go build -o client/main client/main.go