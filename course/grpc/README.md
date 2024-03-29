### Using
- **protoc** `23.4-osx-x86_64`
- **protoc-gen-go** `v1.31.0`
- **protoc-gen-go-grpc** `v1.3.0`

### Compiler command
```sh
$PATH/protoc \
--proto_path=$PATH/greet/proto/ \
--plugin=protoc-gen-go=$PATH/protoc-gen-go \
--plugin=protoc-gen-go-grpc=$PATH/protoc-gen-go-grpc \
--go_opt=module=course/grpc \
--go_out=$PATH/course/grpc/ \
--go-grpc_opt=module=course/grpc \
--go-grpc_out=$PATH/course/grpc/ \
$PATH/greet/proto/*.proto
```

### gRPC types
- Unary
>  rpc Service (Request) returns (Response);
- Server Streaming
>  rpc Service (Request) returns (stream Response);
- Client Streaming
>  rpc Service (stream Request) returns (Response);
- Bi directional Streaming
>  rpc Service (stream Request) returns (stream Response);