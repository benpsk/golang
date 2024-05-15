package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var contents []byte

func main() {
	fmt.Println("Contents of file: ", string(contents))
}
