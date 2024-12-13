package main

import (
	"fmt"
	"strings"
)

func decrypteVigenere(cipherText, keyword string) {

	cipherText = strings.ToUpper(cipherText)
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	keyword = strings.ToUpper(keyword)

	keyLength := len(keyword)
	/*
		decoded := make([]rune, len(cipherText))

		for i, ch := range cipherText {
			keyChar := rune(keyword[i%keyLength])
			shift := (ch - keyChar + 26) % 26
			decoded[i] = 'A' + shift
		}
		fmt.Println(string(decoded))
	*/
	decoded := ""
	for i := 0; i < len(cipherText); i++ {
		ch := cipherText[i]
		keyChar := keyword[i%keyLength]
		ch = (ch-keyChar+26)%26 + 'A'
		decoded = decoded + string(ch)
	}
	fmt.Println(decoded)
}

func vigenereCipher() {
	cipherText := "CS OIT EUI WUI ZNS ROCNKFD"
	keyword := "GOLANG"
	decrypteVigenere(cipherText, keyword)
}
