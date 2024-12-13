package main

import (
	"fmt"
	"time"
)

func channel() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleppyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished slepping")
	}
}

func sleppyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id
}
