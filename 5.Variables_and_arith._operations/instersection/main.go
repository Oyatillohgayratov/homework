package main

import "fmt"

func main() {
	// Ikki slice butun sonlarning kesishishini topuvchi dasturini tuzing.
	// Ex: [1,2,3,4] [3,4,7,9] -> [3,4]
	
	s1 := []int{1, 2, 3, 4, 3}
	s2 := []int{3, 4, 7, 9, 4, 6,3}
	s := []int{}
	news := []int{}
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j]{
				s = append(s, s1[i])
				break
			}
		}
	}
	
	// birxil sonlarni olib tashlash
	for i := 0; i<len(s)-1; i++ {
		for j := i+1; j < len(s); j++ {
			if s[i] == s[j] {
                s[i]=0
            }
        }
	}
	for i := 0; i<len(s); i++ {
		if s[i]!= 0 {
            news = append(news, s[i])
        }
	}
	fmt.Print(news)
}
