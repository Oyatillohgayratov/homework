//  1. Oldindan belgilangan qoidalarga muvofiq foydalanuvchi
//     kiritishini ValidateUser digan method yarating.
//
// Qoidalar:
// Name empty bo'lishi kerak emas
// Name uzunligi kamida 6 belgigi bo'lishi kerak
// Age 0 dan kotta va 120 dan kichik bo'lishi
// Email empty bo'lishi kerak emas
// Email formatiga mos bolishi kerak (masalan example@domain.com)
//
// 2. Error slice yaratilgan holda barcha paydo bo'lgan errorlarni qoshib yuvorin
// 3. Foydalanuvchi ma'lumotlarni terminaldan oqib oling
// 4. Oqib oliniyatgan jarayonda errorlarni ohirida chiqarib berin
//
// Masalan:
// Name: asd
// Age: 123
// Email: ""

// Errors:
// Name: length cannot be less than a 6 characters
// Age: couldn't be more than 120
// Email: couldn't be empty

package main

import (
	"errors"
	"fmt"
	"strings"
)
type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) ValidateUser() []error{
	var err []error

	if len(u.Name) < 6 {
        err = append(err, errors.New("Name: length cannot be less than a 6 characters"))
    }
	if u.Age <= 0 || u.Age > 120 {
        err = append(err, errors.New("Age: couldn't be more than 120"))
    }

	if u.Email == "" {
		err = append(err, errors.New("Email: cannot be empty"))
	} else {
		if !strings.Contains(u.Email, "@") || !strings.Contains(u.Email, ".") {
			err = append(err, errors.New("Email: invalid format"))
		}
	}

	return err
}

func main() {
	var user User

	fmt.Print("Name: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Age: ")
	fmt.Scanln(&user.Age)
	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)

	err := user.ValidateUser()
	if len(err) > 0 {
		fmt.Println("\nErorrs: ")
        for _, e := range err {
            fmt.Println(e.Error())
        }
    }else {
		fmt.Println("success")
	}
}
