## Protobuf code generation

Requirements:

- protoc
- protoc-gen-go
- protoc-gen-go-grpc

Generate Go stubs from repo root:

```bash
protoc -I proto \
  --go_out=. \
  --go-grpc_out=. \
  proto/items.proto
```
