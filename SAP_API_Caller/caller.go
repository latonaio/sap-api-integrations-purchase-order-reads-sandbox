package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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

		
func (c *SAPAPICaller) AsyncGetPurchaseOrder(PurchaseOrder, PurchaseOrderItem, PurchaseRequisition, PurchaseRequisitionItem string) {
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go func() {
		c.Header(PurchaseOrder)
		wg.Done()
	}()
	go func() {
		c.Item(PurchaseOrder, PurchaseOrderItem)
		wg.Done()
	}()
	go func() {
		c.PurchaseRequisition(PurchaseRequisition, PurchaseRequisitionItem)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(PurchaseOrder string) {
	res, err := c.callPurchaseOrderSrvAPIRequirementHeader("A_PurchaseOrder", PurchaseOrder)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Item(PurchaseOrder, PurchaseOrderItem string) {
	res, err := c.callPurchaseOrderSrvAPIRequirementItem("A_PurchaseOrder('{PurchaseOrder}')/to_PurchaseOrderItem", PurchaseOrder, PurchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) PurchaseRequisition(PurchaseRequisition, PurchaseRequisitionItem string) {
	res, err := c.callPurchaseOrderSrvAPIRequirementPurchaseRequisition("A_PurchaseOrder('{PurchaseOrder}')/to_PurchaseOrderItem", PurchaseRequisition, PurchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementHeader(api, PurchaseOrder string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseOrder")
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s'", PurchaseOrder))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItem(api, PurchaseOrder, PurchaseOrderItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseOrder, PurchaseOrderItem")
	params.Add("$filter", fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", PurchaseOrder, PurchaseOrderItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementPurchaseRequisition(api, PurchaseRequisition, PurchaseRequisitionItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseRequisition, PurchaseRequisitionItem")
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", PurchaseRequisition, PurchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}