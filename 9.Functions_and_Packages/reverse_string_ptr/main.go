package main

import "fmt"

func main() {
	// Pointer yordamida string ni joyida (funksiya return qimidi) qaytaruvchi funktsiyani ishlab chiqing.
	// Ex: hello => olleh
	s := "hello"
	ReverseString(&s)
	fmt.Printf("%s\n", s)
}

func ReverseString(s *string) {
	var ns string
	for i := len(*s) - 1; i >= 0; i-- {
        ns += string((*s)[i])
    }
	(*s) = ns
}
