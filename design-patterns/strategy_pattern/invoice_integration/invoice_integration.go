package invoice_integration

import "fmt"

type InvoiceIntegrationInterface interface {
	SetInvoiceGenerator(generator InvoiceGenerator)
	GetInvoiceGenerator() InvoiceGenerator
	GenerateInvoice() error
}

type InvoiceIntegration struct {
	Client           string
	Amount           float64
	invoiceGenerator InvoiceGenerator
}

func NewInvoiceIntegration(client string, amount float64, generator InvoiceGenerator) *InvoiceIntegration {
	return &InvoiceIntegration{
		Client:           client,
		Amount:           amount,
		invoiceGenerator: generator,
	}
}

func (ii *InvoiceIntegration) SetInvoiceGenerator(generator InvoiceGenerator) {
	ii.invoiceGenerator = generator
}

func (ii *InvoiceIntegration) GetInvoiceGenerator() InvoiceGenerator {
	return ii.invoiceGenerator
}

func (ii *InvoiceIntegration) GenerateInvoice() error {
	fmt.Println(ii.invoiceGenerator.GenerateInvoice(*ii))
	return nil
}
