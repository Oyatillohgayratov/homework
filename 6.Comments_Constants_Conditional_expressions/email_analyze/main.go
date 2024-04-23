package main

import (
	"fmt"
)

func main() {
	// Elektron pochta manzillarini tahlil qiling:
	// emailList o'zgaruvchidan domen nomlarini
	// (masalan, "example.com") chiqaradigan dasturni ishlab chiqing.
	// Keyin, har bir domen nomining takrorlanishini hisoblang va natijalarni chop eting.

	emailList := []string{
			"john@example.com",
			"alice@example.com",
			"bob@example.com",
			"sam@example.net",
			"lisa@example.org",
			"test@example.com",
			"example@example.com",
			"user1@gmail.com",
			"user2@gmail.com",
			"user3@gmail.com",
			"user4@yahoo.com",
			"user5@yahoo.com",
			"user6@outlook.com",
			"user7@outlook.com",
			"admin@example.com",
			"info@example.com",
			"contact@example.com",
			"support@example.com",
			"sales@example.com",
			"feedback@example.com",
			"webmaster@example.com",
		}

	m := make(map[string]int)
	for i:=0 ; i < len(emailList);i++{
		for j:= 0; j < len(emailList[i]); j++{
			if emailList[i][j] == '@'{
				m[emailList[i][j:]]++
			}
		}
	}
	for v,k := range m{
		fmt.Printf("%d ta %s\n",k,v)
	}
}
