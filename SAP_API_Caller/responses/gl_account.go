package responses

type GLAccount struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SupplierInvoice            string `json:"SupplierInvoice"`
			FiscalYear                 string `json:"FiscalYear"`
			SupplierInvoiceItem        string `json:"SupplierInvoiceItem"`
			CompanyCode                string `json:"CompanyCode"`
			CostCenter                 string `json:"CostCenter"`
			ControllingArea            string `json:"ControllingArea"`
			BusinessArea               string `json:"BusinessArea"`
			ProfitCenter               string `json:"ProfitCenter"`
			FunctionalArea             string `json:"FunctionalArea"`
			GLAccount                  string `json:"GLAccount"`
			SalesOrder                 string `json:"SalesOrder"`
			SalesOrderItem             string `json:"SalesOrderItem"`
			CostObject                 string `json:"CostObject"`
			CostCtrActivityType        string `json:"CostCtrActivityType"`
			BusinessProcess            string `json:"BusinessProcess"`
			WBSElement                 string `json:"WBSElement"`
			DocumentCurrency           string `json:"DocumentCurrency"`
			SupplierInvoiceItemAmount  string `json:"SupplierInvoiceItemAmount"`
			TaxCode                    string `json:"TaxCode"`
			PersonnelNumber            string `json:"PersonnelNumber"`
			WorkItem                   string `json:"WorkItem"`
			DebitCreditCode            string `json:"DebitCreditCode"`
			TaxJurisdiction            string `json:"TaxJurisdiction"`
			SupplierInvoiceItemText    string `json:"SupplierInvoiceItemText"`
			AssignmentReference        string `json:"AssignmentReference"`
			IsNotCashDiscountLiable    bool   `json:"IsNotCashDiscountLiable"`
			InternalOrder              string `json:"InternalOrder"`
			ProjectNetwork             string `json:"ProjectNetwork"`
			NetworkActivity            string `json:"NetworkActivity"`
			CommitmentItem             string `json:"CommitmentItem"`
			FundsCenter                string `json:"FundsCenter"`
			TaxBaseAmountInTransCrcy   string `json:"TaxBaseAmountInTransCrcy"`
			Fund                       string `json:"Fund"`
			GrantID                    string `json:"GrantID"`
			QuantityUnit               string `json:"QuantityUnit"`
			SuplrInvcItmQtyUnitSAPCode string `json:"SuplrInvcItmQtyUnitSAPCode"`
			SuplrInvcItmQtyUnitISOCode string `json:"SuplrInvcItmQtyUnitISOCode"`
			Quantity                   string `json:"Quantity"`
			PartnerBusinessArea        string `json:"PartnerBusinessArea"`
			FinancialTransactionType   string `json:"FinancialTransactionType"`
			TaxCountry                 string `json:"TaxCountry"`
			BudgetPeriod               string `json:"BudgetPeriod"`
		} `json:"results"`
	} `json:"d"`
}
