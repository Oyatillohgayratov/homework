package main

import "fmt"

// String tipidagi kirish parametriga ega bo'lgan va string qaytaruvchi funksiyani qaytaruvchi
// prefixer deb nomlangan funktsiyani yozing.
// Qaytarilgan funktsiya o'z kiritilishini prefixer o'tgan string bilan oldindan belgilashi kerak.
// Prefixerni sinab ko'rish uchun quyidagi asosiy funktsiyadan foydalaning:

func main() {
	helloPrefix := prefixer("Salom")
	fmt.Println(helloPrefix("Bob"))    // Salom Bobni chop etish kerak
	fmt.Println(helloPrefix("Mariya")) // Salom Mariyani chop etish kerak

}


func prefixer(prefix string) func(name string) string {
	return func(name string) string {
        return prefix + " " + name
    }
}
