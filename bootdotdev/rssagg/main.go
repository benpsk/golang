package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}
	fmt.Println("Port: ", portString)
}
