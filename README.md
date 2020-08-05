Generating protos

```
protoc -I proto proto/services/helloworld/helloworld.proto --go_out=plugins=grpc:./gen
```

Build docker

```
docker build --tag go-envoy:latest .
```

Run docker

```
docker run -p 55051:55051 go-envoy:latest
```