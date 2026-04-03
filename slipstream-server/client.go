package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:9090")
    if err != nil {
        log.Fatalf("dial error: %v", err)
    }
    defer conn.Close()

    msg := "hello slipstream\n"
    _, err = conn.Write([]byte(msg))
    if err != nil {
        log.Fatalf("write error: %v", err)
    }

    reply, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        log.Fatalf("read error: %v", err)
    }

    fmt.Printf("server replied: %s", reply)
}
