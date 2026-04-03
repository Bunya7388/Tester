package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"

	"github.com/quic-go/quic-go"
)

func main() {
	tlsConf := &tls.Config{
		Certificates: []tls.Certificate{generateTLSConfig()},
		NextProtos:   []string{"quic-echo-example"},
	}

	listener, err := quic.ListenAddr("0.0.0.0:4242", tlsConf, nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("quic-go server listening on 0.0.0.0:4242")

	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go func(conn *quic.Conn) {
			stream, err := conn.AcceptStream(context.Background())
			if err != nil {
				log.Printf("accept stream error: %v", err)
				return
			}
			log.Println("accepted stream")

			buf := make([]byte, 1024)
			n, err := stream.Read(buf)
			if err != nil && err != io.EOF {
				log.Printf("read error: %v", err)
				return
			}

			msg := string(buf[:n])
			fmt.Printf("received: %s\n", msg)
			_, _ = stream.Write([]byte("echo: " + msg))
			stream.Close()
		}(conn)
	}
}

func generateTLSConfig() tls.Certificate {
	cert, err := tls.X509KeyPair([]byte(serverCert), []byte(serverKey))
	if err != nil {
		log.Fatalf("invalid cert: %v", err)
	}
	return cert
}

const serverCert = `-----BEGIN CERTIFICATE-----
MIID...REPLACE_WITH_REAL_CERT...IDAQAB
-----END CERTIFICATE-----`

const serverKey = `-----BEGIN PRIVATE KEY-----
MIIE...REPLACE_WITH_REAL_KEY...AoIBAQ
-----END PRIVATE KEY-----`
