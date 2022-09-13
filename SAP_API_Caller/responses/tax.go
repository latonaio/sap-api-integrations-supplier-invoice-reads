package responses

type Tax struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SupplierInvoice          string `json:"SupplierInvoice"`
			FiscalYear               string `json:"FiscalYear"`
			TaxCode                  string `json:"TaxCode"`
			DocumentCurrency         string `json:"DocumentCurrency"`
			TaxAmount                string `json:"TaxAmount"`
			TaxBaseAmountInTransCrcy string `json:"TaxBaseAmountInTransCrcy"`
		} `json:"results"`
	} `json:"d"`
}
