package main

type EmployeeDetail struct {
	EmployeeName    string
	EmployeeID      string
	InvoiceMonth    string
	InvoiceYear     string
	PeriodFrom      string
	PeriodTo        string
	WorkingDuration string
	ContractorRate  string
	InvoiceAmount   string
}

type OneInvoiceDetail []*EmployeeDetail

type Result struct {
	InvDetail OneInvoiceDetail
	InvoiceNo string
	Error     error
}
