package main

import "fmt"

type Dictionary map[string]string

func main() {
	// 1. Foydalanuvchilarga so'zni kiritish va lug'atda mavjud bo'lsa, uning ta'rifini ko'rsating.
	// 2. Lug'atti qiymatlari bo'yicha so'z qidirish (masalan: small berilganda barcha key larni qaytarish)

	dict := Dictionary{
		"apple":      "a fruit that grows on trees and is typically red, yellow, or green",
		"banana":     "a long curved fruit which grows in clusters and has soft pulpy flesh and yellow skin when ripe",
		"orange":     "a round juicy citrus fruit with a tough bright reddish-yellow rind",
		"grape":      "a small oval fruit with smooth skin and juicy flesh, typically eaten raw and used for making wine",
		"pineapple":  "a large tropical fruit with a tough orange or green skin, sweet yellow flesh, and a fibrous core",
		"watermelon": "a large melon with a hard, green rind and sweet, juicy, usually red flesh",
		"strawberry": "a sweet soft red fruit with a seed-studded surface",
		"blueberry":  "a small sweet fruit, typically dark blue, growing on low bushes",
		"peach":      "a round stone fruit with juicy yellow flesh and downy pinkish-yellow skin",
		"pear":       "a yellowish- or brownish-green edible fruit that is typically narrower at the stalk and wider toward the base",
		"kiwi":       "a brown, hairy fruit with bright green flesh and black seeds",
		"mango":      "a tropical fruit with juicy flesh, distinctive aroma, and sweet flavor",
		"avocado":    "a pear-shaped fruit with a rough, dark green or black skin, smooth flesh, and large stone",
		"lemon":      "a yellow, oval citrus fruit with thick skin and fragrant, acidic juice",
		"lime":       "a rounded citrus fruit with green skin, juicy flesh, and a sour taste",
		"coconut":    "a large oval brown seed of a tropical palm, consisting of a hard shell lined with edible white flesh and containing a clear liquid",
		"papaya":     "a tropical fruit shaped like a large melon, with sweet orange flesh and many black seeds",
		"melon":      "a large round fruit with sweet juicy flesh and a thick rind, typically eaten sliced",
		"fig":        "a soft pear-shaped fruit with sweet dark flesh and many small seeds, eaten fresh or dried",
		"apricot":    "a small, round, yellow or orange fruit with a slightly sour taste, eaten fresh or dried",
		"plum":       "a round or oval fruit with a smooth skin and a flattish pointed stone, sweet or tart in flavor",
		"cherry":     "a small, round stone fruit that is typically bright or dark red",
	}
	_ = dict
	var a string
	fmt.Scan(&a)
	for k, v := range dict {
		if a == k {
            fmt.Println( v)
        }else if a == "small"{
			fmt.Println(k)
		}
	}

}
