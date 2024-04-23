package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var s string
	fmt.Print("malumot kiriting: ")
	fmt.Scan(&s)

	fmt.Fprintln(f, s)
	if err != nil {
		fmt.Println(err)
		return
	}



	date := time.Now()
	str := fmt.Sprintf("file_backup_%s.txt", date.Format("2006-01-02_150405"))
	w, err := os.Create(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	r,err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
        return
	}

	_, err = io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Faylga yozildi va muvaffaqiyatli nusxalandi!")
	defer w.Close()
}