package main

import "fmt"

func main() {
	// String ni palindrom ekanligini tekshirish dasturini tuzing
	// Ex: Palindrome: "91019", chap bilan o'n tomonidan o'qilsa bir hil natija bo'ladi

	var s string = "91019"
	a := len(s)-1
	t_f := true
	for i:=0;i<len(s)/2;i++{
		if s[i] != s[a]{
			t_f = false
			break
		}
		a--
	}
	fmt.Print(t_f)
}
