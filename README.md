# sap-api-integrations-supplier-invoice-reads
sap-api-integrations-supplier-invoice-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 仕入先請求書データを取得するマイクロサービスです。    
sap-api-integrations-supplier-invoice-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-supplier-invoice-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SUPPLIERINVOICE_PROCESS_SRV/overview  

## 動作環境  
sap-api-integrations-supplier-invoice-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-supplier-invoice-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-supplier-invoice-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SUPPLIERINVOICE_PROCESS_SRV/overview    
* APIサービス名(=baseURL): API_SUPPLIERINVOICE_PROCESS_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-supplier-invoice-reads には、次の API をコールするためのリソースが含まれています。  

* A_SupplierInvoice（仕入先請求書 - ヘッダ）
* A_SupplierInvoice(SupplierInvoice='{SupplierInvoice}',FiscalYear='{FiscalYear}')/to_SupplierInvoiceItemGLAcct（仕入先請求書 - 明細会計情報）
* A_SuplrInvcItemPurOrdRef（仕入先請求書 - 購買発注参照）
* A_SuplrInvcItemAcctAssgmt（仕入先請求書 - 明細勘定設定）

## API への 値入力条件 の 初期値
sap-api-integrations-supplier-invoice-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inputSDC.SupplierInvoice.SupplierInvoice（仕入先請求書）
* inputSDC.SupplierInvoice.FiscalYear（会計年度）
* inputSDC.SupplierInvoice.PurchaseOrderReference.SupplierInvoiceItem.PurchaseOrder（購買発注）
* inputSDC.SupplierInvoice.PurchaseOrderReference.SupplierInvoiceItem.PurchaseOrderItem（購買発注明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "sap.s4.beh.product.v1.Product.Created.v1",
  	"accepter": ["Header"],
  	"supplier_invoice": "21",
  	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
  	"api_schema": "sap.s4.beh.supplierinvoice.v1.	SupplierInvoice.Created.v1",
  	"accepter": ["All"],
  	"supplier_invoice": "5100000001",
  	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```
## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 仕入先請求書 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"SupplierInvoice" ～ "ReverseDocumentFiscalYear" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-supplier-invoice-reads/SAP_API_Caller/caller.go#L50",
	"function": "sap-api-integrations-supplier-invoice-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": "[{XXXXXXXXXXXXXXXXXXXXXXXXXXXXX}]",
	"time": "2021-12-11T15:33:00.054455+09:00"
}
```
