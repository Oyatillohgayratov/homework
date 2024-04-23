package main

import (
	"fmt"
	"strings"
)

var str = `
	Alice was beginning to get very tired of sitting by her sister on the bank, 
	and of having nothing to do: once or twice she had peeped into the book her sister was reading, 
	but it had no pictures or conversations in it, 
	'and what is the use of a book,' thought Alice 'without pictures or conversation?'
	So she was considering in her own mind (as well as she could, 
	for the hot day made her feel very sleepy and stupid), 
	whether the pleasure of making a daisy-chain would be worth the trouble of getting up 
	and picking the daisies, when suddenly a White Rabbit with pink eyes ran close by her.
`
func main() {
	// Unli va undoshlarni sanash:
	// Str o'zgaruvchi ni ichida unli va undoshlar sonini hisoblaydigan dastur tuzing.
	// Bo'shliqlarga e'tibor bermang va katta va kichik harflarni ekvivalent deb hisoblang.
	
	// Unli harflar Ingliz alifbosi
	vowel_letters := "aeiouyAEIOUY"
	vowel := 0
	consonant := 0
	for i:= 0; i < len(str); i++ {
		if str[i] >= 65 && str[i] <= 90 || str[i] >= 97 && str[i] <= 122{
			if strings.Count(vowel_letters, string(str[i])) == 0{
                consonant++
            } else {
                vowel++
            }
		}
	}
	fmt.Printf("number of vowel:%d\nnumber of consonants:%d",vowel,consonant)


}
