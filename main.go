package main

import (
	sap_api_caller "sap-api-integrations-credit-memo-request-reads/SAP_API_Caller"
	"sap-api-integrations-credit-memo-request-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Credit_Memo_Request_Item_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}

	caller.AsyncGetCreditMemoRequest(
		inoutSDC.CreditMemoRequest.CreditMemoRequest,
		inoutSDC.CreditMemoRequest.CreditMemoRequestItem.CreditMemoRequestItem,
		accepter,
	)
}