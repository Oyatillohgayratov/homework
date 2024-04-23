package main

import "fmt"

func main() {
	// Pointer yordamida ikkita butun sonning qiymatlarini almashtiruvchi funksiya yozing
	// Ex: a = 1, b = 2 => a = 2, b = 1
	a:=1
	b:=2
	fmt.Printf("a = %d, b = %d\n", a, b)
	Changenum(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}


func Changenum(a *int, b *int){
	*a, *b = *b, *a
}
