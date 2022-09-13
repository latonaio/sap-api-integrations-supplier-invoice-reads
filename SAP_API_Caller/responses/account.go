package responses

type Account struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SupplierInvoice               string `json:"SupplierInvoice"`
			FiscalYear                    string `json:"FiscalYear"`
			SupplierInvoiceItem           string `json:"SupplierInvoiceItem"`
			CompanyCode                   string `json:"CompanyCode"`
			CostCenter                    string `json:"CostCenter"`
			ControllingArea               string `json:"ControllingArea"`
			BusinessArea                  string `json:"BusinessArea"`
			ProfitCenter                  string `json:"ProfitCenter"`
			FunctionalArea                string `json:"FunctionalArea"`
			GLAccount                     string `json:"GLAccount"`
			SalesOrder                    string `json:"SalesOrder"`
			SalesOrderItem                string `json:"SalesOrderItem"`
			CostObject                    string `json:"CostObject"`
			WBSElement                    string `json:"WBSElement"`
			DocumentCurrency              string `json:"DocumentCurrency"`
			SuplrInvcAcctAssignmentAmount string `json:"SuplrInvcAcctAssignmentAmount"`
			TaxCode                       string `json:"TaxCode"`
			WorkItem                      string `json:"WorkItem"`
			MasterFixedAsset              string `json:"MasterFixedAsset"`
			FixedAsset                    string `json:"FixedAsset"`
			DebitCreditCode               string `json:"DebitCreditCode"`
			InternalOrder                 string `json:"InternalOrder"`
			ProjectNetwork                string `json:"ProjectNetwork"`
			ProfitabilitySegment          string `json:"ProfitabilitySegment"`
		} `json:"results"`
	} `json:"d"`
}
