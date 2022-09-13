package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SupplierInvoice               string `json:"SupplierInvoice"`
			FiscalYear                    string `json:"FiscalYear"`
			CompanyCode                   string `json:"CompanyCode"`
			DocumentDate                  string `json:"DocumentDate"`
			PostingDate                   string `json:"PostingDate"`
			SupplierInvoiceIDByInvcgParty string `json:"SupplierInvoiceIDByInvcgParty"`
			InvoicingParty                string `json:"InvoicingParty"`
			DocumentCurrency              string `json:"DocumentCurrency"`
			InvoiceGrossAmount            string `json:"InvoiceGrossAmount"`
			DocumentHeaderText            string `json:"DocumentHeaderText"`
			PaymentTerms                  string `json:"PaymentTerms"`
			DueCalculationBaseDate        string `json:"DueCalculationBaseDate"`
			NetPaymentDays                string `json:"NetPaymentDays"`
			PaymentBlockingReason         string `json:"PaymentBlockingReason"`
			AccountingDocumentType        string `json:"AccountingDocumentType"`
			BPBankAccountInternalID       string `json:"BPBankAccountInternalID"`
			SupplierInvoiceStatus         string `json:"SupplierInvoiceStatus"`
			DirectQuotedExchangeRate      string `json:"DirectQuotedExchangeRate"`
			SupplyingCountry              string `json:"SupplyingCountry"`
			PaymentMethod                 string `json:"PaymentMethod"`
			InvoiceReference              string `json:"InvoiceReference"`
			SupplierPostingLineItemText   string `json:"SupplierPostingLineItemText"`
			TaxIsCalculatedAutomatically  bool   `json:"TaxIsCalculatedAutomatically"`
			BusinessArea                  string `json:"BusinessArea"`
			SupplierInvoiceIsCreditMemo   string `json:"SupplierInvoiceIsCreditMemo"`
			ReverseDocument               string `json:"ReverseDocument"`
			ReverseDocumentFiscalYear     string `json:"ReverseDocumentFiscalYear"`
		} `json:"results"`
	} `json:"d"`
}
