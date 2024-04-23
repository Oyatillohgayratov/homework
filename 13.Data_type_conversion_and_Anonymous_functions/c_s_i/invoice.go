package csi

import "fmt"

// create the invoice for the reservation
func CreateAndSaveInvoice(name string, nights uint, price float32) {
	fmt.Printf("Invoice for %d night(s) for %s: $%.2f\n", nights, name, price)
}
