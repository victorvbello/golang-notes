package main

import "github.com/victorvbello/gonotes/design-patterns/strategy_pattern/invoice_integration"

func main() {

	soapGenerator := &invoice_integration.SOAPInvoiceGenerator{}
	restGenerator := &invoice_integration.RESTInvoiceGenerator{}
	grpcGenerator := &invoice_integration.GRPCInvoiceGenerator{}

	GovernmentInvoice := invoice_integration.NewInvoiceIntegration("Government", 1000.00, soapGenerator)
	CorporateInvoice := invoice_integration.NewInvoiceIntegration("Corporate", 500.00, restGenerator)
	StartupInvoice := invoice_integration.NewInvoiceIntegration("Startup", 250.00, grpcGenerator)
	NaturalPersonInvoice := invoice_integration.NewInvoiceIntegration("Natural Person", 100.00, restGenerator)

	GovernmentInvoice.GenerateInvoice()
	CorporateInvoice.GenerateInvoice()
	StartupInvoice.GenerateInvoice()
	NaturalPersonInvoice.GenerateInvoice()
}
