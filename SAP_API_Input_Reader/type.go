package sap_api_input_reader

type EC_MC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	SupplierInvoice struct {
		SupplierInvoice             string `json:"document_no"`
		DeliverTo                   string `json:"deliver_to"`
		QuantityInPurchaseOrderUnit string `json:"quantity"`
		PickedQuantity              string `json:"picked_quantity"`
		SupplierInvoiceItemAmount   string `json:"price"`
		Batch                       string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema              string `json:"api_schema"`
	MaterialCode           string `json:"material_code"`
	InvoicingParty         string `json:"plant/supplier"`
	Stock                  string `json:"stock"`
	AccountingDocumentType string `json:"document_type"`
	SupplierInvoiceNo      string `json:"document_no"`
	PlannedDate            string `json:"planned_date"`
	PostingDate            string `json:"validated_date"`
	Deleted                bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	SupplierInvoice struct {
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
		Tax                           struct {
			TaxCode                  string `json:"TaxCode"`
			DocumentCurrency         string `json:"DocumentCurrency"`
			TaxAmount                string `json:"TaxAmount"`
			TaxBaseAmountInTransCrcy string `json:"TaxBaseAmountInTransCrcy"`
		} `json:"Tax"`
		PurchaseOrder struct {
			SupplierInvoiceItem            string `json:"SupplierInvoiceItem"`
			PurchaseOrder                  string `json:"PurchaseOrder"`
			PurchaseOrderItem              string `json:"PurchaseOrderItem"`
			Plant                          string `json:"Plant"`
			TaxCode                        string `json:"TaxCode"`
			DocumentCurrency               string `json:"DocumentCurrency"`
			SupplierInvoiceItemAmount      string `json:"SupplierInvoiceItemAmount"`
			PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
			QuantityInPurchaseOrderUnit    string `json:"QuantityInPurchaseOrderUnit"`
			PurchaseOrderPriceUnit         string `json:"PurchaseOrderPriceUnit"`
			QtyInPurchaseOrderPriceUnit    string `json:"QtyInPurchaseOrderPriceUnit"`
			SupplierInvoiceItemText        string `json:"SupplierInvoiceItemText"`
			PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
		} `json:"PurchaseOrder"`
		AccountAssignment struct {
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
		} `json:"AccountAssignment"`
	} `json:"SupplierInvoice"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	SupplierInvoiceNo string   `json:"supplier_invoice"`
	Deleted           bool     `json:"deleted"`
}
