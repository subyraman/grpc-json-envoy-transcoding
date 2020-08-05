# Simple GRPC-JSON Transcoding Example with Go and Envoy

I created this project as a learning experience for understanding how Envoy can be used to host public-facing GRPC APIs. Envoy is used [to transcode JSON/HTTP1.1 traffic into GRPC messages](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/other_protocols/grpc#arch-overview-grpc) which are passed to the GRPC service, written in Go.

### To build the services:

```
docker-compose build
```

### Running the services:

```
docker-compose up
```

This will run two containers: an `envoy` container, exposed on the host at `0.0.0.0:10000`. JSON traffic is transcoded and proxied to the `grpc-api` container. 

### Calling the server

Send a JSON POST request to the Envoy proxy:

```
curl 0.0.0.0:10000/say -d '{"name": "Suby"}'
```

### Send a GRPC request to the Envoy proxy:

This will run `client/main` in the `grpc-api` container, which will send a GRPC request to the `envoy` container. 

```
docker-compose run grpc-api client/main envoy:10000 Suby
```

### Errors

Supplying an empty name will return an error:

`curl 0.0.0.0:10000/say -d '{"name": ""}'`

## Building/running locally

If you wish to build the Go server/client locally, you will need to install the [protobuf compiler](https://github.com/protocolbuffers/protobuf) and [protobuf-gen-go](https://developers.google.com/protocol-buffers/docs/reference/go-generated).

Download all dependences with `go mod download`.

After that you can run `./local-rebuild.sh` to recompile protocol buffers and rebuild the Go server/client.

Run the server at `./server/main`.

Run the client at `./client/main 0.0.0.0:50051 Suby`.
