package main

//Reimbursement Details
type RInvoiceEmp struct {
	InvoiceNo     string  `json:"str_invoice_number"`
	EmployeeID    string  `json:"str_employee_id"`
	EmployeeName  string  `json:"str_employee_name"`
	InvoiceMonth  string  `json:"str_invoice_month"`
	InvoiceYear   string  `json:"str_invoice_year"`
	InvoiceAmount float64 `json:"n_payment_amt"`
	FromDt        string  `json:"DT_EXPENSE_DATE_FROM"`
	ToDt          string  `json:"DT_EXPENSE_DATE_TO"`
}

type RFinalInvoices struct {
	Result []RInvoiceEmp `json:"result"`
}

type RResult struct {
	Invoice []RInvoiceEmp
	Err     error
}

//Invoices

type InvoiceEmp struct {
	EmployeeId      string  `json:"str_employee_id"`
	EmployeeName    string  `json:"str_employee_name"`
	ProjectCode     string  `json:"str_project_code"`
	InoviceMonth    string  `json:"str_invoice_month"`
	InvoiceYear     string  `json:"str_invoice_year"`
	TimeSheetStDt   string  `json:"dt_timesheet_start_date"`
	TimeSheetEndDt  string  `json:"dt_timesheet_end_date"`
	RmHours         string  `json:"str_rm_hrs"`
	RmDays          string  `json:"str_rm_days"`
	CurrenySymbol   string  `json:"str_contractor_currency_desc"`
	RateDesc        string  `json:"str_contractor_rate_desc"`
	RateStr         string  `json:"str_rate"`
	InvoiceAmount   float64 `json:"n_payment_amt"`
	TaxAmount       float64 `json:"N_TAX_AMOUNT"`
	ContractorState string  `json:"contractorState"`
	WiproGSTNo      string  `json:"wiproStateGSTN"`
	SezIndicator    string  `json:"sezIndicator"`
}

type InvoicePid struct {
	InvoiceNo   string `json:"str_invoice_number"`
	InvoiceDt   string `json:"dt_invoice_date"`
	SerivceType string `json:"existingSAC_number"`
}

type AllEmpInvoice struct {
	Result []InvoiceEmp `json:"result"`
}

type FinalInvoices struct {
	EmpInvoices   []InvoiceEmp
	OneInvoicePid *InvoicePid
}

type Result struct {
	Invoice *FinalInvoices
	Err     error
}

//Beena
type CreateInvoiceEmp struct {
	EmployeeName    string `json:"employeeName"`
	ContractorId    string `json:"contractorId"`
	ResumeNumber    string `json:"resumeNumber"`
	EmpCompanyCode  string `json:"empCompanyCode"`
	TimesheetYear   string `json:"timesheetYear"`
	TimeSheetMonth  string `json:"timeSheetMonth"`
	Rate            string `json:"Rate"`
	RateTypeDesc    string `json:"RateTypeDesc"`
	ContractorState string `json:"contractorState"`
	Rmdayshours     string `json:"Rmdayshours"`
	ProjectCode     string `json:"projectCode"`
	InvoiceAmount   string `json:"invoiceAmount"`
}

type AllCreateInvoiceEmpList struct {
	Result []CreateInvoiceEmp `json:"result"`
}
