package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToPurchaseOrder(raw []byte, l *logger.Logger) *PurchaseOrder {
	pm := &responses.PurchaseOrder{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PurchaseOrder{
		PurchaseOrder                  data.PurchaseOrder,
		CompanyCode                    data.CompanyCode,
		PurchaseOrderType              data.PurchaseOrderType,
		PurchasingProcessingStatus     data.PurchasingProcessingStatus,
		CreationDate                   data.CreationDate,
		LastChangeDateTime             data.LastChangeDateTime,
		Supplier                       data.Supplier,
		Language                       data.Language,
		PaymentTerms                   data.PaymentTerms,
		PurchasingOrganization         data.PurchasingOrganization,
		PurchasingGroup                data.PurchasingGroup,
		PurchaseOrderDate              data.PurchaseOrderDate,
		DocumentCurrency               data.DocumentCurrency,
		ExchangeRate                   data.ExchangeRate,
		ValidityStartDate              data.ValidityStartDate,
		ValidityEndDate                data.ValidityEndDate,
		SupplierRespSalesPersonName    data.SupplierRespSalesPersonName,
		SupplierPhoneNumber            data.SupplierPhoneNumber,
		SupplyingPlant                 data.SupplyingPlant,
		IncotermsClassification        data.IncotermsClassification,
		ManualSupplierAddressID        data.ManualSupplierAddressID,
		AddressName                    data.AddressName,
		AddressCityName                data.AddressCityName,
		AddressFaxNumber               data.AddressFaxNumber,
		AddressPostalCode              data.AddressPostalCode,
		AddressStreetName              data.AddressStreetName,
		AddressPhoneNumber             data.AddressPhoneNumber,
		AddressRegion                  data.AddressRegion,
		AddressCountry                 data.AddressCountry,
		PurchaseOrderItem              data.PurchaseOrderItem,
		PurchaseOrderItemText          data.PurchaseOrderItemText,
		Plant                          data.Plant,
		StorageLocation                data.StorageLocation,
		MaterialGroup                  data.MaterialGroup,
		PurchasingInfoRecord           data.PurchasingInfoRecord,
		SupplierMaterialNumber         data.SupplierMaterialNumber,
		OrderQuantity                  data.OrderQuantity,
		PurchaseOrderQuantityUnit      data.PurchaseOrderQuantityUnit,
		OrderPriceUnit                 data.OrderPriceUnit,
		DocumentCurrency               data.DocumentCurrency,
		NetPriceAmount                 data.NetPriceAmount,
		NetPriceQuantity               data.NetPriceQuantity,
		ScheduleLineDeliveryDate       data.ScheduleLineDeliveryDate,
		TaxCode                        data.TaxCode,
		OverdelivTolrtdLmtRatioInPct   data.OverdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed data.UnlimitedOverdeliveryIsAllowed,
		UnderdelivTolrtdLmtRatioInPct  data.UnderdelivTolrtdLmtRatioInPct,
		IsCompletelyDelivered          data.IsCompletelyDelivered,
		IsFinallyInvoiced              data.IsFinallyInvoiced,
		PurchaseOrderItemCategory      data.PurchaseOrderItemCategory,
		AccountAssignmentCategory      data.AccountAssignmentCategory,
		GoodsReceiptIsExpected         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated      data.GoodsReceiptIsNonValuated,
		InvoiceIsExpected              data.InvoiceIsExpected,
		InvoiceIsGoodsReceiptBased     data.InvoiceIsGoodsReceiptBased,
		Customer                       data.Customer,
		SupplierIsSubcontractor        data.SupplierIsSubcontractor,
		ItemNetWeight                  data.ItemNetWeight,
		ItemWeightUnit                 data.ItemWeightUnit,
		IncotermsClassification        data.IncotermsClassification,
		PurchaseRequisition            data.PurchaseRequisition,
		PurchaseRequisitionItem        data.PurchaseRequisitionItem,
		RequisitionerName              data.RequisitionerName,
		Material                       data.Material,
		InternationalArticleNumber     data.InternationalArticleNumber,
		DeliveryAddressID              data.DeliveryAddressID,
		DeliveryAddressName            data.DeliveryAddressName,
		DeliveryAddressPostalCode      data.DeliveryAddressPostalCode,
		DeliveryAddressStreetName      data.DeliveryAddressStreetName,
		DeliveryAddressCityName        data.DeliveryAddressCityName,
		DeliveryAddressRegion          data.DeliveryAddressRegion,
		DeliveryAddressCountry         data.DeliveryAddressCountry,
		PurchasingDocumentDeletionCode data.PurchasingDocumentDeletionCode,
		AccountAssignmentNumber        data.AccountAssignmentNumber,
		GLAccount                      data.GLAccount,
		BusinessArea                   data.BusinessArea,
		CostCenter                     data.CostCenter,
		SalesOrder                     data.SalesOrder,
		SalesOrderItem                 data.SalesOrderItem,
		SalesOrderScheduleLine         data.SalesOrderScheduleLine,
		MasterFixedAsset               data.MasterFixedAsset,
		FixedAsset                     data.FixedAsset,
		GoodsRecipientName             data.GoodsRecipientName,
		UnloadingPointName             data.UnloadingPointName,
		ControllingArea                data.ControllingArea,
		CostObject                     data.CostObject,
		OrderID                        data.OrderID,
		ProfitCenter                   data.ProfitCenter,
		WBSElement                     data.WBSElement,
		ProjectNetwork                 data.ProjectNetwork,
		FunctionalArea                 data.FunctionalArea,
		TaxCode                        data.TaxCode,
		CostCtrActivityType            data.CostCtrActivityType,
		IsDeleted                      data.IsDeleted,
	}
}
