package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-supplier-invoice-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			SupplierInvoice:               data.SupplierInvoice,
			FiscalYear:                    data.FiscalYear,
			CompanyCode:                   data.CompanyCode,
			DocumentDate:                  data.DocumentDate,
			PostingDate:                   data.PostingDate,
			SupplierInvoiceIDByInvcgParty: data.SupplierInvoiceIDByInvcgParty,
			InvoicingParty:                data.InvoicingParty,
			DocumentCurrency:              data.DocumentCurrency,
			InvoiceGrossAmount:            data.InvoiceGrossAmount,
			DocumentHeaderText:            data.DocumentHeaderText,
			PaymentTerms:                  data.PaymentTerms,
			DueCalculationBaseDate:        data.DueCalculationBaseDate,
			NetPaymentDays:                data.NetPaymentDays,
			PaymentBlockingReason:         data.PaymentBlockingReason,
			AccountingDocumentType:        data.AccountingDocumentType,
			BPBankAccountInternalID:       data.BPBankAccountInternalID,
			SupplierInvoiceStatus:         data.SupplierInvoiceStatus,
			DirectQuotedExchangeRate:      data.DirectQuotedExchangeRate,
			SupplyingCountry:              data.SupplyingCountry,
			PaymentMethod:                 data.PaymentMethod,
			InvoiceReference:              data.InvoiceReference,
			SupplierPostingLineItemText:   data.SupplierPostingLineItemText,
			TaxIsCalculatedAutomatically:  data.TaxIsCalculatedAutomatically,
			BusinessArea:                  data.BusinessArea,
			SupplierInvoiceIsCreditMemo:   data.SupplierInvoiceIsCreditMemo,
			ReverseDocument:               data.ReverseDocument,
			ReverseDocumentFiscalYear:     data.ReverseDocumentFiscalYear,
			ToPurchaseOrder:               data.ToPurchaseOrder.Deferred.URI,
			ToGLAccount:                   data.ToGLAccount.Deferred.URI,
			ToTax:                         data.ToTax.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToTax(raw []byte, l *logger.Logger) ([]Tax, error) {
	pm := &responses.Tax{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Tax. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	tax := make([]Tax, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		tax = append(tax, Tax{
			SupplierInvoice:          data.SupplierInvoice,
			FiscalYear:               data.FiscalYear,
			TaxCode:                  data.TaxCode,
			DocumentCurrency:         data.DocumentCurrency,
			TaxAmount:                data.TaxAmount,
			TaxBaseAmountInTransCrcy: data.TaxBaseAmountInTransCrcy,
		})
	}

	return tax, nil
}

func ConvertToAccount(raw []byte, l *logger.Logger) ([]Account, error) {
	pm := &responses.Account{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Account. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	account := make([]Account, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		account = append(account, Account{
			SupplierInvoice:               data.SupplierInvoice,
			FiscalYear:                    data.FiscalYear,
			SupplierInvoiceItem:           data.SupplierInvoiceItem,
			CompanyCode:                   data.CompanyCode,
			CostCenter:                    data.CostCenter,
			ControllingArea:               data.ControllingArea,
			BusinessArea:                  data.BusinessArea,
			ProfitCenter:                  data.ProfitCenter,
			FunctionalArea:                data.FunctionalArea,
			GLAccount:                     data.GLAccount,
			SalesOrder:                    data.SalesOrder,
			SalesOrderItem:                data.SalesOrderItem,
			CostObject:                    data.CostObject,
			WBSElement:                    data.WBSElement,
			DocumentCurrency:              data.DocumentCurrency,
			SuplrInvcAcctAssignmentAmount: data.SuplrInvcAcctAssignmentAmount,
			TaxCode:                       data.TaxCode,
			WorkItem:                      data.WorkItem,
			MasterFixedAsset:              data.MasterFixedAsset,
			FixedAsset:                    data.FixedAsset,
			DebitCreditCode:               data.DebitCreditCode,
			InternalOrder:                 data.InternalOrder,
			ProjectNetwork:                data.ProjectNetwork,
			ProfitabilitySegment:          data.ProfitabilitySegment,
		})
	}

	return account, nil
}

func ConvertToPurchaseOrder(raw []byte, l *logger.Logger) ([]PurchaseOrder, error) {
	pm := &responses.PurchaseOrder{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchaseOrder. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	purchaseOrder := make([]PurchaseOrder, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		purchaseOrder = append(purchaseOrder, PurchaseOrder{
			SupplierInvoice:                data.SupplierInvoice,
			FiscalYear:                     data.FiscalYear,
			SupplierInvoiceItem:            data.SupplierInvoiceItem,
			PurchaseOrder:                  data.PurchaseOrder,
			PurchaseOrderItem:              data.PurchaseOrderItem,
			Plant:                          data.Plant,
			TaxCode:                        data.TaxCode,
			DocumentCurrency:               data.DocumentCurrency,
			SupplierInvoiceItemAmount:      data.SupplierInvoiceItemAmount,
			PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
			QuantityInPurchaseOrderUnit:    data.QuantityInPurchaseOrderUnit,
			PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
			QtyInPurchaseOrderPriceUnit:    data.QtyInPurchaseOrderPriceUnit,
			SupplierInvoiceItemText:        data.SupplierInvoiceItemText,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
			ToAccount:                      data.ToAccount.Deferred.URI,
		})
	}

	return purchaseOrder, nil
}

func ConvertToToAccount(raw []byte, l *logger.Logger) ([]ToAccount, error) {
	pm := &responses.ToAccount{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toAccount := make([]ToAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toAccount = append(toAccount, ToAccount{
			SupplierInvoice:               data.SupplierInvoice,
			FiscalYear:                    data.FiscalYear,
			SupplierInvoiceItem:           data.SupplierInvoiceItem,
			CompanyCode:                   data.CompanyCode,
			CostCenter:                    data.CostCenter,
			ControllingArea:               data.ControllingArea,
			BusinessArea:                  data.BusinessArea,
			ProfitCenter:                  data.ProfitCenter,
			FunctionalArea:                data.FunctionalArea,
			GLAccount:                     data.GLAccount,
			SalesOrder:                    data.SalesOrder,
			SalesOrderItem:                data.SalesOrderItem,
			CostObject:                    data.CostObject,
			WBSElement:                    data.WBSElement,
			DocumentCurrency:              data.DocumentCurrency,
			SuplrInvcAcctAssignmentAmount: data.SuplrInvcAcctAssignmentAmount,
			TaxCode:                       data.TaxCode,
			WorkItem:                      data.WorkItem,
			MasterFixedAsset:              data.MasterFixedAsset,
			FixedAsset:                    data.FixedAsset,
			DebitCreditCode:               data.DebitCreditCode,
			InternalOrder:                 data.InternalOrder,
			ProjectNetwork:                data.ProjectNetwork,
			ProfitabilitySegment:          data.ProfitabilitySegment,
		})
	}

	return toAccount, nil
}

func ConvertToToGLAccount(raw []byte, l *logger.Logger) ([]ToGLAccount, error) {
	pm := &responses.ToGLAccount{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToGLAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toGLAccount := make([]ToGLAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toGLAccount = append(toGLAccount, ToGLAccount{
			SupplierInvoice:            data.SupplierInvoice,
			FiscalYear:                 data.FiscalYear,
			SupplierInvoiceItem:        data.SupplierInvoiceItem,
			CompanyCode:                data.CompanyCode,
			CostCenter:                 data.CostCenter,
			ControllingArea:            data.ControllingArea,
			BusinessArea:               data.BusinessArea,
			ProfitCenter:               data.ProfitCenter,
			FunctionalArea:             data.FunctionalArea,
			GLAccount:                  data.GLAccount,
			SalesOrder:                 data.SalesOrder,
			SalesOrderItem:             data.SalesOrderItem,
			CostObject:                 data.CostObject,
			CostCtrActivityType:        data.CostCtrActivityType,
			BusinessProcess:            data.BusinessProcess,
			WBSElement:                 data.WBSElement,
			DocumentCurrency:           data.DocumentCurrency,
			SupplierInvoiceItemAmount:  data.SupplierInvoiceItemAmount,
			TaxCode:                    data.TaxCode,
			PersonnelNumber:            data.PersonnelNumber,
			WorkItem:                   data.WorkItem,
			DebitCreditCode:            data.DebitCreditCode,
			TaxJurisdiction:            data.TaxJurisdiction,
			SupplierInvoiceItemText:    data.SupplierInvoiceItemText,
			AssignmentReference:        data.AssignmentReference,
			IsNotCashDiscountLiable:    data.IsNotCashDiscountLiable,
			InternalOrder:              data.InternalOrder,
			ProjectNetwork:             data.ProjectNetwork,
			NetworkActivity:            data.NetworkActivity,
			CommitmentItem:             data.CommitmentItem,
			FundsCenter:                data.FundsCenter,
			TaxBaseAmountInTransCrcy:   data.TaxBaseAmountInTransCrcy,
			Fund:                       data.Fund,
			GrantID:                    data.GrantID,
			QuantityUnit:               data.QuantityUnit,
			SuplrInvcItmQtyUnitSAPCode: data.SuplrInvcItmQtyUnitSAPCode,
			SuplrInvcItmQtyUnitISOCode: data.SuplrInvcItmQtyUnitISOCode,
			Quantity:                   data.Quantity,
			PartnerBusinessArea:        data.PartnerBusinessArea,
			FinancialTransactionType:   data.FinancialTransactionType,
			TaxCountry:                 data.TaxCountry,
			BudgetPeriod:               data.BudgetPeriod})
	}

	return toGLAccount, nil
}

func ConvertToToPurchaseOrder(raw []byte, l *logger.Logger) ([]ToPurchaseOrder, error) {
	pm := &responses.ToPurchaseOrder{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPurchaseOrder. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toPurchaseOrder := make([]ToPurchaseOrder, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPurchaseOrder = append(toPurchaseOrder, ToPurchaseOrder{
			SupplierInvoice:                data.SupplierInvoice,
			FiscalYear:                     data.FiscalYear,
			SupplierInvoiceItem:            data.SupplierInvoiceItem,
			PurchaseOrder:                  data.PurchaseOrder,
			PurchaseOrderItem:              data.PurchaseOrderItem,
			Plant:                          data.Plant,
			TaxCode:                        data.TaxCode,
			DocumentCurrency:               data.DocumentCurrency,
			SupplierInvoiceItemAmount:      data.SupplierInvoiceItemAmount,
			PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
			QuantityInPurchaseOrderUnit:    data.QuantityInPurchaseOrderUnit,
			PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
			QtyInPurchaseOrderPriceUnit:    data.QtyInPurchaseOrderPriceUnit,
			SupplierInvoiceItemText:        data.SupplierInvoiceItemText,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
		})
	}

	return toPurchaseOrder, nil
}

func ConvertToToTax(raw []byte, l *logger.Logger) ([]ToTax, error) {
	pm := &responses.ToTax{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToTax. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toTax := make([]ToTax, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toTax = append(toTax, ToTax{
			SupplierInvoice:          data.SupplierInvoice,
			FiscalYear:               data.FiscalYear,
			TaxCode:                  data.TaxCode,
			DocumentCurrency:         data.DocumentCurrency,
			TaxAmount:                data.TaxAmount,
			TaxBaseAmountInTransCrcy: data.TaxBaseAmountInTransCrcy,
		})
	}

	return toTax, nil
}
