// Ish haqini hisoblash uchun CalculateSalary() va DisplayDetails() kabi methodlardan foydalangan holda
// Employee deb nomlangan interfeysni aniqlang va
// to'liq kunlik, yarim kunlik va shartnoma xodimlari kabi har xil turdagi xodimlarning tafsilotlarini ko'rsating.
// Har bir xodim turi uchun interfeysni amalga oshiring.
// Shartnoma xodimlari oyiga 160 soatni ishashini farase qiling
package main

import (
	"fmt"
)

type Employee interface {
	CalculateSalary() float64
	DisplayDetails()
}

type FullTimeEmployee struct {
	Name   string
	Salary float64
}

func (f FullTimeEmployee) CalculateSalary() float64 {
	return f.Salary
}

func (f FullTimeEmployee) DisplayDetails() {
	fmt.Printf("Full-Time Employee\nName: %s\nSalary: %.2f\n", f.Name, f.Salary)
}

type PartTimeEmployee struct {
	Name       string
	HourlyRate float64
	HoursWorked int
}

func (p PartTimeEmployee) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}

func (p PartTimeEmployee) DisplayDetails() {
	fmt.Printf("Part-Time Employee\nName: %s\nHourly Rate: %.2f\nHours Worked: %d\n", p.Name, p.HourlyRate, p.HoursWorked)
}

type ContractEmployee struct {
	Name       string
	MonthlyRate float64
}

func (c ContractEmployee) CalculateSalary() float64 {
	hoursWorked := 160.0 // Assume contract employees work 160 hours per month
	return c.MonthlyRate / hoursWorked
}

func (c ContractEmployee) DisplayDetails() {
	fmt.Printf("Contract Employee\nName: %s\nMonthly Rate: %.2f\n", c.Name, c.MonthlyRate)
}

func main() {
	fullTimeEmp := FullTimeEmployee{Name: "kimdir", Salary: 50 }
	partTimeEmp := PartTimeEmployee{Name: "Sarah", HourlyRate: 10.0, HoursWorked: 30}
	contractEmp := ContractEmployee{Name: "Mike", MonthlyRate: 2000.0}

	employees := []Employee{fullTimeEmp, partTimeEmp, contractEmp}

	for _, emp := range employees {
		emp.DisplayDetails()
		fmt.Printf("Salary: %.2f\n\n", emp.CalculateSalary())
	}
}