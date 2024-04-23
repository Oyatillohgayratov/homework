package sendemail_test

import "fmt"



func SendEmail(emailContents string, customerEmail string){
    fmt.Printf("Sending email to %s: %s\n", customerEmail, emailContents)
}