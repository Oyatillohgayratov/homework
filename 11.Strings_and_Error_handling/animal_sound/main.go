// Hayvon tomonidan chiqarilgan tovushni qaytaruvchi MakeSound() methodo yordamida
// Animal deb nomlangan interfeys yarating.
// It, mushuk va qush kabi turli hayvonlar uchun MakeSound() methodini qo'llang.
// Turli hayvonlarning tovush chiqarishi uchun interfeysdan foydalaning.
package main

import "fmt"

type Animal interface {
	Makesaund() string
}

type Dog struct {}
func (d Dog) Makesaund() string {
	return "woof"
}

type Cat struct {}
func (c Cat) Makesaund() string {
    return "meow"
}

type bird struct {}
func (b bird) Makesaund() string {
    return "coo-coo"
}




func main() {

	var animal Animal
	var s string
	fmt.Printf("Animals (dog , cat , bird): ")
	fmt.Scan(&s)
	if s == "dog" {
		animal = Dog{}
		fmt.Println(animal.Makesaund())
	} else if s == "cat" {
		animal = Cat{}
		fmt.Println(animal.Makesaund())
	}else if s == "bird"{
		animal = bird{}
		fmt.Println(animal.Makesaund())
	}else {
		fmt.Println("Invalid input")
	}

}
