package main

import "fmt"

type Student struct {
	ism       string
	Identifikator int
	Baholar     []float64
}

func (s *Student) Ortacha_bahoni_hisoblang() float64 {
	sum := 0.0
	for _, grade := range s.Baholar {
		sum += grade
	}
	average := sum / float64(len(s.Baholar))
	return average
}

func (s *Student) AddGrade(grade float64) {
	s.Baholar = append(s.Baholar, grade)
}

func (s *Student) DisplayStudentInfo() {
	fmt.Printf("ism: %s\n", s.ism)
	fmt.Printf("Identifikator: %d\n", s.Identifikator)
	fmt.Printf("Baholar: %v\n", s.Baholar)
	average := s.Ortacha_bahoni_hisoblang()
	fmt.Printf("Average Grade: %.2f\n", average)
}

func main() {
	// Talabaning ismi, identifikatori, baholari kabi ma’lumotlarni saqlash uchun 'Student' tuzilmasini yarating.
	// O‘rtacha bahoni hisoblash, yangi baho qo‘shish va talaba ma’lumotlarini ko‘rsatish methodlarini qo‘llang.
	student := Student{
		ism:       "John Doe",
		Identifikator: 12345,
		Baholar:     []float64{85.5, 92.0, 78.3},
	}

	student.DisplayStudentInfo()

	student.AddGrade(90.2)
	fmt.Println("Yangi baho qo'shildi.")

	student.DisplayStudentInfo()
}