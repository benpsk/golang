package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePointer := flag.String("filePath", "test.txt", "file path to read from")
	file, err := os.Open(*filePointer)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(file)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			fmt.Println("finished reading file")
			break
		}
		if err != nil {
			fmt.Printf("Error %s reading file", err)
		}
		fmt.Println(string(b[0:n]))
	}
}
