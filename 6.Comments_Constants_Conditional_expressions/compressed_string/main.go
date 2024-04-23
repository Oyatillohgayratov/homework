package main

import (
	"fmt"
)

func main() {
	// Ketma-ket takrorlanuvchi belgilarni takrorlash sonidan keyin belgi bilan almashtirib,
	// siqib chiqaradigan dastur yarating.
	// Masalan, "abbbccc" "a3b3c3" ga aylanishi kerak.

	s := "aaabbbccc"
	s =  " " + s + " "
	ns := ""
	a := 1
	for i := 1; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			a++
		} else {
			ns += fmt.Sprint(string(s[i]),a)
			a = 1
		}
	}
	fmt.Println(ns)
}
