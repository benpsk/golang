package main

import (
	"fmt"
)

func ceasarCipher() {
	message := "L fdph, L vdz, L frqtxhuhg."
	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
