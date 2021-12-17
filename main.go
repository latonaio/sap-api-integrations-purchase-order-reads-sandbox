package main

import (
	sap_api_caller "sap-api-integrations-purchase-order-reads/SAP_API_Caller"
	"sap-api-integrations-purchase-order-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Purchase_Order_Schedule_Line_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item", "Account", "PurchaseRequisition",
			"ScheduleLine",
		}
	}

	caller.AsyncGetPurchaseOrder(
		inoutSDC.PurchaseOrder.PurchaseOrder,
		inoutSDC.PurchaseOrder.PurchaseOrderItem.PurchaseOrderItem,
		inoutSDC.PurchaseOrder.PurchaseOrderItem.PurchaseRequisition,
		inoutSDC.PurchaseOrder.PurchaseOrderItem.PurchaseRequisitionItem,
		inoutSDC.PurchasingDocument,
		inoutSDC.PurchaseOrder.PurchaseOrderItem.ScheduleLine.PurchasingDocumentItem,
		accepter,
	)
}
