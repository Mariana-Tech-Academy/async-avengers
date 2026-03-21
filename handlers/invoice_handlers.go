func (h *InvoiceHandler) GetInvoiceStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid invoice ID"})
		return
	}
	status, err := h.service.GetInvoiceStatus(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get invoice status"})
		return
	}
	c.JSON(200, gin.H{"status": status})
}