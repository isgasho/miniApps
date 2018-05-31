package main

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

type InvoiceSummary struct {
	AgencyName         string  `json:"str_agency_name"`
	ClientId           int     `json:"n_cons_focrec_id"`
	InvoiceNumber      string  `json:"str_invoice_number"`
	InvoiceDt          string  `json:"dt_invoice_date"`
	LastStatusDt       string  `json:"dt_last_status"`
	InvoiceStatusDesc  string  `json:"str_invoice_status_desc"`
	InvoiceStatus      string  `json:"str_invoice_status"`
	TotalInvoiceAmtStr string  `json:"str_total_invoice_amount"`
	TotalInvoiceAmt    float64 `json:"n_total_invoice_amount"`
	//InvoiceType         string  `json:"str_type_of_invoice"`
	//NewInvoiceNumber    string  `json:"STR_NEW_INVOICENUMBER"`
	//ChequeHandOverDt    string  `json:"dt_cheque_handover"`
	//Comments            string  `json:"str_comments"`
	//ModifyInvoiceNumber string  `json:"STR_MODIFY_INV_COMMENTS"`
	//SapConsultantCode  string `json:"str_sap_consultant_code"`
	//RegistrationNumber string `json:"str_registration_num"`
	//AmountModify        float64 `json:"n_modified_amt"`
}

type AllInvoiceSummary struct {
	Result []InvoiceSummary `json:"result"`
}

type AllCreateInvoiceEmpList struct {
	Result []CreateInvoiceEmp `json:"result"`
}

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
	InvoiceNo          string          `json:"str_invoice_number"`
	InvoiceDt          string          `json:"dt_invoice_date"`
	LastStatusDt       string          `json:"dt_last_status"`
	VendorCode         string          `json:"str_sap_consultant_code"`
	AgencyGSTN         string          `json:"agencyGSTN"`
	CGSTAmt            string          `json:"CGSTAmt"`
	SGSTAmt            string          `json:"SGSTAmt"`
	SerivceType        string          `json:"existingSAC_number"`
	ContractState      string          `json:"existingcontractorState"`
	WiproVendorDetails []VendorDetails `json:"vendorgstnList"`
}

type VendorDetails struct {
	ClassDesc   string `json:"strGSTVedorClassDescription"`
	Address     string `json:"strVendorAddress"`
	GSTCode     string `json:"strVendorGSTSAPCode"`
	Name        string `json:"strVendorName"`
	VendorState string `json:"strVendorStateSAPCode"`
}

type AllEmpInvoice struct {
	Result []InvoiceEmp `json:"result"`
}

type FinalInvoices struct {
	EmpInvoices   []InvoiceEmp
	OneInvoicePid *InvoicePid
}
