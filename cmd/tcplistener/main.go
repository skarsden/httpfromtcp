package main

import (
	"fmt"
	"httpfromtcp/internal/request"
	"log"
	"net"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println("Connection has been accepted from", port)

		request, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalf("unable to read from connection: %s", err)
		}
		fmt.Printf("Request line:\n- Method: %s\n- Target: %s\n- Version: %s\n", request.RequestLine.Method, request.RequestLine.RequestTarget, request.RequestLine.HttpVersion)
		fmt.Println("Headers:")
		for key, value := range request.Headers {
			fmt.Printf("- %s: %s\n", key, value)
		}
		fmt.Printf("Body:\n%s\n", string(request.Body))
		fmt.Println("Connection has been closed")
	}
}
