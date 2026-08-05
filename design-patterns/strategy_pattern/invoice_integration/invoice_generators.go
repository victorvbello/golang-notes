package invoice_integration

import "fmt"

type InvoiceGenerator interface {
	GenerateInvoice(ii InvoiceIntegration) string
}

type SOAPInvoiceGenerator struct{}

func (s *SOAPInvoiceGenerator) GenerateInvoice(ii InvoiceIntegration) string {
	return "SOAP Invoice for " + ii.Client + " with amount " + fmt.Sprintf("%.2f", ii.Amount)
}

type RESTInvoiceGenerator struct{}

func (r *RESTInvoiceGenerator) GenerateInvoice(ii InvoiceIntegration) string {
	return "REST Invoice for " + ii.Client + " with amount " + fmt.Sprintf("%.2f", ii.Amount)
}

type GRPCInvoiceGenerator struct{}

func (g *GRPCInvoiceGenerator) GenerateInvoice(ii InvoiceIntegration) string {
	return "GRPC Invoice for " + ii.Client + " with amount " + fmt.Sprintf("%.2f", ii.Amount)
}
