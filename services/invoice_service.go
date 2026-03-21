func (s *InvoiceService) GetInvoiceStatus(id uint) (string, error) {
	var invoice models.Invoice

	err := s.DB.First(&invoice, id).Error
	if err != nil {
		return "", err
	}
	return invoice.Status, nil
} 