package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-supplier-invoice-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetSupplierInvoice(supplierInvoice, fiscalYear, purchaseOrder, purchaseOrderItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(supplierInvoice, fiscalYear)
				wg.Done()
			}()
		case "Tax":
			func() {
				c.Tax(supplierInvoice, fiscalYear)
				wg.Done()
			}()
		case "Account":
			func() {
				c.Account(supplierInvoice, fiscalYear)
				wg.Done()
			}()
		case "PurchaseOrder":
			func() {
				c.PurchaseOrder(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(supplierInvoice, fiscalYear string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementHeader("A_SupplierInvoice", supplierInvoice, fiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementHeader(api, supplierInvoice, fiscalYear string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, supplierInvoice, fiscalYear)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Tax(supplierInvoice, fiscalYear string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementTax("A_SupplierInvoiceTax", supplierInvoice, fiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementTax(api, supplierInvoice, fiscalYear string) ([]sap_api_output_formatter.Tax, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")

	param := c.getQueryWithTax(map[string]string{}, supplierInvoice, fiscalYear)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToTax(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Account(supplierInvoice, fiscalYear string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementAccount("A_SuplrInvcItemAcctAssgmt", supplierInvoice, fiscalYear)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementAccount(api, supplierInvoice, fiscalYear string) ([]sap_api_output_formatter.Account, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")

	param := c.getQueryWithAccount(map[string]string{}, supplierInvoice, fiscalYear)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAccount(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchaseOrder(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callSupplierInvoiceSrvAPIRequirementPurchaseOrder("A_SuplrInvcItemPurOrdRef", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSupplierInvoiceSrvAPIRequirementPurchaseOrder(api, purchaseOrder, purchaseOrderItem string) ([]sap_api_output_formatter.PurchaseOrder, error) {
	url := strings.Join([]string{c.baseURL, "API_SUPPLIERINVOICE_PROCESS_SRV", api}, "/")

	param := c.getQueryWithPurchaseOrder(map[string]string{}, purchaseOrder, purchaseOrderItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchaseOrder(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, supplierInvoice, fiscalYear string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq '%s'", supplierInvoice, fiscalYear)
	return params
}

func (c *SAPAPICaller) getQueryWithTax(params map[string]string, supplierInvoice, fiscalYear string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq '%s'", supplierInvoice, fiscalYear)
	return params
}

func (c *SAPAPICaller) getQueryWithAccount(params map[string]string, supplierInvoice, fiscalYear string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SupplierInvoice eq '%s' and FiscalYear eq '%s'", supplierInvoice, fiscalYear)
	return params
}

func (c *SAPAPICaller) getQueryWithPurchaseOrder(params map[string]string, purchaseOrder, purchaseOrderItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem)
	return params
}
