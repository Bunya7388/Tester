package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/quic-go/quic-go"
)

func main() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}

	conn, err := quic.DialAddr(context.Background(), "localhost:4242", tlsConf, nil)
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Fatalf("write failed: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		log.Fatalf("read failed: %v", err)
	}
	fmt.Printf("server replied: %s\n", string(buf[:n]))
}
