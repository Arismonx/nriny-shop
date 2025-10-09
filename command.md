### command proto fiel `Auth` service

```
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/auth/authPb/authPb.proto
```

### command proto fiel `Player` service

```
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/player/playerPb/playerPb.proto
```
