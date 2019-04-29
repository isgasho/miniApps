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
	AttendanceDtl   []*EmployeeAttendance
}

type OneInvoiceDetail []*EmployeeDetail

type Result struct {
	InvDetail OneInvoiceDetail
	InvoiceNo string
	Error     error
}

type EmployeeAttendance struct {
	AttendaceDate    string
	SwipeInDate      string
	InTime           string
	SwipeOutDate     string
	TimeOut          string
	TimeFromAttenSys string
	LunchTime        string
	ActualTimeWorked string
	Remarks          string
}
