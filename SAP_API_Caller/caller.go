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

func (c *SAPAPICaller) AsyncGetPurchaseOrder(purchaseOrder, purchaseOrderItem, purchasingDocument, purchasingDocumentItem, purchaseRequisition, purchaseRequisitionItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(purchaseOrder)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()
		case "ItemScheduleLine":
			func() {
				c.ItemScheduleLine(purchasingDocument, purchasingDocumentItem)
				wg.Done()
			}()
		case "ItemPricingElement":
			func() {
				c.ItemPricingElement(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()
		case "ItemAccount":
			func() {
				c.ItemAccount(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()
		case "PurchaseRequisition":
			func() {
				c.PurchaseRequisition(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(purchaseOrder string) {
	headerData, err := c.callPurchaseOrderSrvAPIRequirementHeader("A_PurchaseOrder", purchaseOrder)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)
	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemScheduleLineData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)

	itemAccountData, err := c.callToItemAccount(itemData[0].ToItemAccount)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemAccountData)

}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementHeader(api, purchaseOrder string) ([]sap_api_output_formatter.Header, error) {
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

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemScheduleLine(url string) ([]sap_api_output_formatter.ToItemScheduleLine, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPricingElement(url string) ([]sap_api_output_formatter.ToItemPricingElement, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemAccount(url string) ([]sap_api_output_formatter.ToItemAccount, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(purchaseOrder, purchaseOrderItem string) {
	itemData, err := c.callPurchaseOrderSrvAPIRequirementItem("A_PurchaseOrderItem", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemScheduleLineData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)

	itemAccountData, err := c.callToItemAccount(itemData[0].ToItemAccount)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemAccountData)
	
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItem(api, purchaseOrder, purchaseOrderItem string) ([]sap_api_output_formatter.Item, error) {
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

func (c *SAPAPICaller) callToItemScheduleLine2(url string) ([]sap_api_output_formatter.ToItemScheduleLine, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPricingElement2(url string) ([]sap_api_output_formatter.ToItemPricingElement, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemAccount2(url string) ([]sap_api_output_formatter.ToItemAccount, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemScheduleLine(purchasingDocument, purchasingDocumentItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementItemScheduleLine("A_PurchaseOrderScheduleLine", purchasingDocument, purchasingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItemScheduleLine(api, purchasingDocument, purchasingDocumentItem string) ([]sap_api_output_formatter.ItemScheduleLine, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItemScheduleLine(req, purchasingDocument, purchasingDocumentItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemPricingElement(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementItemPricingElement("A_PurOrdPricingElement", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItemPricingElement(api, purchaseOrder, purchaseOrderItem string) ([]sap_api_output_formatter.ItemPricingElement, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItemPricingElement(req, purchaseOrder, purchaseOrderItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) ItemAccount(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementItemAccount("A_PurOrdAccountAssignment", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItemAccount(api, purchaseOrder, purchaseOrderItem string) ([]sap_api_output_formatter.ItemAccount, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItemAccount(req, purchaseOrder, purchaseOrderItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemAccount(byteArray, c.log)
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

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementPurchaseRequisition(api, purchaseRequisition, purchaseRequisitionItem string) ([]sap_api_output_formatter.Item, error) {
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

func (c *SAPAPICaller) getQueryWithItemScheduleLine(req *http.Request, purchasingDocument, purchasingDocumentItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchasingDocument eq '%s' and PurchasingDocumentItem eq '%s'", purchasingDocument, purchasingDocumentItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItemPricingElement(req *http.Request, purchaseOrder, purchaseOrderItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItemAccount(req *http.Request, purchaseOrder, purchaseOrderItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPurchaseRequisition(req *http.Request, purchaseRequisition, purchaseRequisitionItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()
}
