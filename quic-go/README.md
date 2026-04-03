# quic-go example

This directory holds a minimal QUIC server/client example using `github.com/lucas-clemente/quic-go`.

## Requirements
- Go 1.20+

## Install

```bash
cd /workspaces/Tester/quic-go
go mod init github.com/Bunya7388/Tester/quic-go
go get github.com/lucas-clemente/quic-go
```

## Run server

```bash
go run ./server.go
```

## Run client (in another shell)

```bash
go run ./client.go
```
