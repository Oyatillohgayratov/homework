package main

import (
	"fmt"
	"math/rand"
)
func main() {
	// 1. 1 dan 50 gacha bo'lgan tasodifiy sonni yarating (ishora: rand.Intn(max-min+1) + min)
	// 2. Tasodifiy sonni length 20 bo'lgan slice yarating
	// 3. Har bir tasodifiy sonni slice ichiga qo'shib chiqin
	// 4. Faqat noyob elementlardan iborat slice ni qaytaring

	aSlice := []int{}
	newSlice := []int{}

	for i := 1; i <= 20; i++ {
        aSlice = append(aSlice, rand.Intn(50)+1)
    }

	// birxil sonlarni olib tashlash
	for i := 0; i<len(aSlice)-1; i++ {
		for j := i+1; j < len(aSlice); j++ {
			if aSlice[i] == aSlice[j] {
                aSlice[i]=0
            }
        }
	}
	for i := 0; i<len(aSlice); i++ {
		if aSlice[i]!= 0 {
            newSlice = append(newSlice, aSlice[i])
        }
	}
	fmt.Println(newSlice)
}
