package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port *int

func init() {
	port = flag.Int("port", 8000, "port")
	flag.Parse()
}

func main() {
	ipport := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", ipport)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
