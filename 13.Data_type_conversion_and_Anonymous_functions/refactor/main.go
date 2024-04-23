// Sizning hamkasblaringizdan biri noyob main.go faylidan iborat dastur yaratdi
// Kodning barqarorligini yaxshilash uchun sizdan kodni qayta tahrirlashingiz so'raladi (refactor).
// Yangi kod tashkilotini taklif qiling:

// Qaysi paketlarni yaratishingiz kerak?

// Yangi katalog yaratish kerakmi?

package main

import (
		send "homework/sendemail"
		get "homework/getemail"
		invoice "homework/c_s_i"
    )

func main() {
	// first reservation
	customerName := "Doe"
	customerEmail := "john.doe@example.com"
	var nights uint = 12
	emailContents := get.GetEmailContents("M", customerName, nights)
	send.SendEmail(emailContents, customerEmail)
	invoice.CreateAndSaveInvoice(customerName, nights, 145.32)
	
}




