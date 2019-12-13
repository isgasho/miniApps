package main

import (
	"fmt"
)

type ASPDocFields struct {
	DocumentAuthor                string
	DocumentAuthorDesignation     string
	DocumentAuthorContactNo       string
	ProposalNumber                string
	ProposalDate                  string
	ProposalValidityDays          string
	ProjectLeadTimeInWeeks        string
	ProjectCompletionDays         string
	ApplicationCurrentRelease     string
	ApplicationCurrentReleaseDate string

	BankName       string
	BranchCount    string
	BranchCountNum int
	Branches       []string

	IncludeBCPDR                  bool
	RPOInMinutes                  string
	RTOInMinutes                  string
	ReplicationLinkBandwithInMbps string
	RentPaymentFrequency          string

	ContractRenewalHikePercentage string
	ContractPeriodInYears         string
	OrganizationAgeInYears        string
	OrganizationStrength          string
}

func initData() *ASPDocFields {
	fields := ASPDocFields{
		BankName:                      "The Demo Bank Ltd",
		ProposalNumber:                "AIPL/19-20/20042019.0",
		ProposalDate:                  "20/04/2019",
		ApplicationCurrentReleaseDate: "11/02/2019",
		ApplicationCurrentRelease:     "11.02.01.1",
		ProjectCompletionDays:         "100",
		BranchCount:                   "10",
		DocumentAuthor:                "Muktesh S Shah",
		DocumentAuthorDesignation:     "Director - Business Development",
		DocumentAuthorContactNo:       "+919925204916",
		PropsalValidityDays:           "30",
		ProjectLeadTimeInWeeks:        "4",
		ReplicationLinkBandwithInMbps: "4",
		IncludeBCPDR:                  true,
		RPOInMinutes:                  "45",
		RTOInMinutes:                  "90",
		ContractPeriodYears:           "3",
		RentPaymentFrequency:          "Half-Yearly",
		ContractRenewalHikePercentage: "10",
		OrganizationStrength:          "250",
		OrganizationAgeInYears:        "13",
		Branches:                      []string{"Head Office", "Main Bajar", "Gunj Nagar", "Home Delta"},
	}
	fields.BranchCount = fmt.Sprintf("%d", len(fields.Branches))
	fields.BranchCountNum = len(fields.Branches)
	return &fields
}

func getTemplateFns() map[string]interface{} {
	funcMap := make(map[string]interface{}, 0)
	funcMap["iter"] = N
	funcMap["inc"] = Incr
	return funcMap
}

func N(n int) []struct{} {
	return make([]struct{}, n)
}

func Incr(val int) string {
	return fmt.Sprintf("%d", val+1)
}
