package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-purchase-order-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchaseOrder(purchaseOrder, purchaseOrderItem, purchaseRequisition, purchaseRequisitionItem, purchasingDocument, purchasingDocumentItem string) {
	wg := &sync.WaitGroup{}

	wg.Add(5)
	func() {
		c.Header(purchaseOrder)
		wg.Done()
	}()
	func() {
		c.Item(purchaseOrder, purchaseOrderItem)
		wg.Done()
	}()
	func() {
		c.Account(purchaseOrder, purchaseOrderItem)
		wg.Done()
	}()
	func() {
		c.PurchaseRequisition(purchaseRequisition, purchaseRequisitionItem)
		wg.Done()
	}()
	func() {
		c.ScheduleLine(purchasingDocument, purchasingDocumentItem)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(purchaseOrder string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementHeader("A_PurchaseOrder", purchaseOrder)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementHeader(api, purchaseOrder string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, purchaseOrder)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementItem("A_PurchaseOrderItem", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItem(api, purchaseOrder, purchaseOrderItem string) (*sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, purchaseOrder, purchaseOrderItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Account(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementAccount("A_PurOrdAccountAssignment", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementAccount(api, purchaseOrder, purchaseOrderItem string) (*sap_api_output_formatter.Account, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithAccount(req, purchaseOrder, purchaseOrderItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchaseRequisition(purchaseRequisition, purchaseRequisitionItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementPurchaseRequisition("A_PurchaseOrderItem", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementPurchaseRequisition(api, purchaseRequisition, purchaseRequisitionItem string) (*sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPurchaseRequisition(req, purchaseRequisition, purchaseRequisitionItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ScheduleLine(purchasingDocument, purchasingDocumentItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementScheduleLine("A_PurchaseOrderScheduleLine", purchasingDocument, purchasingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementScheduleLine(api, purchasingDocument, purchasingDocumentItem string) (*sap_api_output_formatter.ScheduleLine, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithScheduleLine(req, purchasingDocument, purchasingDocumentItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, purchaseOrder string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s'", purchaseOrder))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, purchaseOrder, purchaseOrderItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithAccount(req *http.Request, purchaseOrder, purchaseOrderItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPurchaseRequisition(req *http.Request, purchaseRequisition, purchaseRequisitionItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithScheduleLine(req *http.Request, purchasingDocument, purchasingDocumentItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchasingDocument eq '%s' and PurchasingDocumentItem eq '%s'", purchasingDocument, purchasingDocumentItem))
	req.URL.RawQuery = params.Encode()
}
