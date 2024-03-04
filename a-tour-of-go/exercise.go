package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

// slice exercise
func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)

	for y := range picture {
		picture[y] = make([]uint8, dx)

		for x := range picture[y] {
			picture[y][x] = uint8((x + y) / 2)
			//picture[y][x] = uint8(x*y)
			//picture[y][x] = uint8(x^y)
		}
	}
	return picture
}

// map exercise
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	fmt.Println("---------Slice Exercise---------")
	pic.Show(Pic)
	fmt.Println("----------Map Exercise---------")
	wc.Test(WordCount)
}
