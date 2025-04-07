package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

		linesChan := getLinesChannel(conn)
		for line := range linesChan {
			fmt.Println("read:", line)
		}
		fmt.Println("Connection has been closed")
	}
}

func getLinesChannel(conn net.Conn) <-chan string {
	lines := make(chan string)
	go func() {
		defer conn.Close()
		defer close(lines)

		currLine := ""
		for {
			b := make([]byte, 8, 8)
			n, err := conn.Read(b)
			if err != nil {
				if currLine != "" {
					lines <- currLine
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s", err.Error())
				return
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currLine, parts[i])
				currLine = ""
			}
			currLine += parts[len(parts)-1]
		}
	}()
	return lines
}
