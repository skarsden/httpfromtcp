package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const filepath = "messages.txt"

func main() {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	for {
		b := make([]byte, 8, 8)
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s", err)
			break
		}
		str := string(b[:n])
		fmt.Printf("read: %s\n", str)
	}
}
