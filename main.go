package main

import (
	sap_api_caller "sap-api-integrations-supplier-invoice-reads/SAP_API_Caller"
	"sap-api-integrations-supplier-invoice-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Supplier_Invoice_Purchase_Order_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Tax", "Account", "PurchaseOrder",
		}
	}

	caller.AsyncGetSupplierInvoice(
		inoutSDC.SupplierInvoice.SupplierInvoice,
		inoutSDC.SupplierInvoice.FiscalYear,
		inoutSDC.SupplierInvoice.PurchaseOrder.PurchaseOrder,
		inoutSDC.SupplierInvoice.PurchaseOrder.PurchaseOrderItem,
		accepter,
	)
}
