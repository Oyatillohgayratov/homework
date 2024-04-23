package main

import (
	"fmt"
	"strings"
)

func main() {
	// Berilgan gapdagi har bir so‘zning birinchi harfini bosh harf bilan yozish dastur ishlab chiqing.
	// Bo'shliqlar yoki tinish belgilari bilan ajratilgan so'zlarni ko'rib chiqing.
	// Bosh harf bilan yozilgan gapni qaytaring.

	sentence := "this is a test sentence, with punctuation! hello world 123"
	_ = sentence

	for j := 33; j <= 64; j++{
		sentence= strings.ReplaceAll(sentence,string(j),"")
	}
	a := " " + sentence
	ns := []rune(a)
	for i := 0; i < len(ns)-1; i++{
		if ns[i] == 32{
			ns[i+1] = ns[i+1]-32
		}
	}
	fmt.Print(string(ns))
}

