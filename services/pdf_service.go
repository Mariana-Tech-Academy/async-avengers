package services

import (
	"fmt"
	"os"
)

func GenerateInvoicePDF(invoiceID uint) (string, error) {

	filePath := fmt.Sprint("invoice_%d.pdf", invoiceID)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	file.WriteString("Please insert Invoice details")

	return filePath, nil

}
