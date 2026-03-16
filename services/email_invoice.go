package services

import (
	"gopkg.in/gomail.v2"
)

func SendInvoiceEmail(ToEmail string, PDFPath string) error { //Define the function

	m := gomail.NewMessage()                     //Create a new email object
	m.SetHeader("From", "cakecompany@gmail.com") //Set the senders email address
	m.SetHeader("To", ToEmail)                   //Set the recipients email address
	m.SetHeader("Subject", "Invoice")            //Set Email Subject

	m.SetBody("text/plain", "Hi Thanks for using the Cake Complany please find attached your invoice.") //Set Email Body
	m.Attach(PDFPath)                                                                                   //Attache the PDF

	//Create a Connection to the email server
	d := gomail.NewDialer(
		"smtp.gmail.com",        //Simple Mail Transfer Protocool - usiing Gmail
		587,                     //Standard secure SMTP Port
		"cakecompany@gmail.com", //
		"Password",
	)

	return d.DialAndSend(m) // Send the email

}
