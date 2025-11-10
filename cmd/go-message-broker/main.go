package main

import (
	"io"
	"log"
	"net"
)

func intToBytes(n int32) []byte {
	return []byte{
		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var buf = make([]byte, 12)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}

	correlation_id := buf[8:12]
	response := append(intToBytes(0), correlation_id...)

	_, err = conn.Write(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		log.Fatal("Failed to bind to port 9092")
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)

	}

}
