package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	log.Println("slipstream-server listening on :9090")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	r := bufio.NewReader(c)

	line, err := r.ReadString('\n')
	if err != nil {
		log.Printf("read error: %v", err)
		return
	}

	fmt.Printf("received: %s", line)
	_, _ = c.Write([]byte("slipstream ack: " + line))
}
