package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "localhost:42069"

func main() {
	address, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("Error resolving address")
	}

	conn, err := net.DialUDP("udp", nil, address)
	if err != nil {
		fmt.Println("error settting up connection")
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Printf("error writing '%s' to connection: %s", str, err)
		}
	}
}
