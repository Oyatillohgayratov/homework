package main

import "fmt"

type Cars struct {
	Marka string
	Model string
	Yil   int
	Kilo  int
}

type CarRentalSystem struct {
	Cars    []Cars
	Rented  []Cars
}

func (crs *CarRentalSystem) RentCar(car Cars) {
	for i, c := range crs.Cars {
		if c == car {
			crs.Cars = append(crs.Cars[:i], crs.Cars[i+1:]...)
			crs.Rented = append(crs.Rented, car)
			fmt.Printf("Car rented: %s %s\n", car.Marka, car.Model)
			return
		}
	}

	fmt.Println("Car not found.")
}

func (crs *CarRentalSystem) ReturnCar(car Cars) {
	for i, c := range crs.Rented {
		if c == car {
			crs.Rented = append(crs.Rented[:i], crs.Rented[i+1:]...)
			crs.Cars = append(crs.Cars, car)
			fmt.Printf("Car returned: %s %s\n", car.Marka, car.Model)
			return
		}
	}

	fmt.Println("Car not rented.")
}

func (crs *CarRentalSystem) DisplayAvailableCars() {
	fmt.Println("Available cars:")
	for _, car := range crs.Cars {
		fmt.Printf("Marka: %s, Model: %s, Yil: %d, Kilo: %d\n", car.Marka, car.Model, car.Yil, car.Kilo)
	}
}

func main() {
	// Talabaning ismi, identifikatori, baholari kabi ma’lumotlarni saqlash uchun 'Student' tuzilmasini yarating.
	// O‘rtacha bahoni hisoblash, yangi baho qo‘shish va talaba ma’lumotlarini ko‘rsatish methodlarini qo‘llang.
	system := CarRentalSystem{
		Cars: []Cars{
			{Marka: "BMW", Model: "M3", Yil: 2020, Kilo: 1000},
			{Marka: "Audi", Model: "A4", Yil: 2020, Kilo: 2000},
			{Marka: "Mercedes-Benz", Model: "C-Class", Yil: 2020, Kilo: 3000},
			{Marka: "Volkswagen", Model: "Jetta", Yil: 2020, Kilo: 4000},
		},
		Rented: []Cars{},
	}

	system.DisplayAvailableCars()

	// Rent a car
	carToRent := system.Cars[0]
	system.RentCar(carToRent)

	// Return a car
	carToReturn := system.Rented[0]
	system.ReturnCar(carToReturn)

	system.DisplayAvailableCars()
}