package main

import "fmt"

func main() {
	// Ikkita maplarni birlashtiradigan dastur tuzing
	// Masalan: {"a": 1, "b": 2} | {"a": 2, "c": 3} => {"a": 2, "b": 2, "c": 3}

	dec1 := map[string]int{"a": 1, "b": 2}
	dec2 := map[string]int{"a": 2, "c": 3}
	dec3 := make(map[string]int)

	for k, v := range dec1 {
		for k2, v2 := range dec2 {
			dec3[k] = v
			dec3[k2] = v2
		}
    }
	fmt.Println(dec3)
}

