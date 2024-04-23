package main

import (
	"fmt"
	"os"
)

func main() {
	// fileLen funktsiya fayl nomini oladi va fayldagi baytlar sonini qaytaradi.
	// Agar faylni o'qishda xatolik bo'lsa, xatoni qaytaring.
	// Fayl to'g'ri yopilganligiga ishonch hosil qilish uchun defer dan foydalaning.
	// file oqish uchun os.Open method iborat
	len,err := fileLen("/home/azamat/Projects/8.Formatted_output_Arrays_and_slices/fileLen/test.txt")
	if err != nil {
        fmt.Println(err)
		return
    }
	fmt.Println(len)
}

func fileLen(fileName string) (int, error) {
	f,err := os.Open(fileName)

	if err != nil {
		return 0,err
    }

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		return 0,err
    }

	f.Close()
	return n , err
}
