package services

import (
	"os"
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendInvoiceEmail(ToEmail string, PDFPath string) error { //Define the function


//1. Create a new email object
	m := gomail.NewMessage()                     
	m.SetHeader("From", os.Getenv("EMAIL_USER")) //Set the senders email address
	m.SetHeader("To", ToEmail)                   //Set the recipients email address
	m.SetHeader("Subject", "Invoice")            //Set Email Subject

	m.SetBody("text/plain", "Hi Thanks for using the Cake Complany please find attached your invoice.") //Set Email Body
	
//2. Attach PDF
	m.Attach(PDFPath)                                                                                 

						

//3. Create a Connection to the email server
	d := gomail.NewDialer(
		"smtp.gmail.com",        //Simple Mail Transfer Protocool - usiing Gmail
		587,                     //Standard secure SMTP Port
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASS"),
	)


//4. Send email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("Failed to send email:%v", err)

	
		} // Send the email

		return nil

}
