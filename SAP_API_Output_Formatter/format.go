package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Header{
		PurchaseOrder:                  data.PurchaseOrder,
		CompanyCode:                    data.CompanyCode,
		PurchaseOrderType:              data.PurchaseOrderType,
		PurchasingProcessingStatus:     data.PurchasingProcessingStatus,
		CreationDate:                   data.CreationDate,
		LastChangeDateTime:             data.LastChangeDateTime,
		Supplier:                       data.Supplier,
		Language:                       data.Language,
		PaymentTerms:                   data.PaymentTerms,
		PurchasingOrganization:         data.PurchasingOrganization,
		PurchasingGroup:                data.PurchasingGroup,
		PurchaseOrderDate:              data.PurchaseOrderDate,
		DocumentCurrency:               data.DocumentCurrency,
		ExchangeRate:                   data.ExchangeRate,
		ValidityStartDate:              data.ValidityStartDate,
		ValidityEndDate:                data.ValidityEndDate,
		SupplierRespSalesPersonName:    data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:            data.SupplierPhoneNumber,
		SupplyingPlant:                 data.SupplyingPlant,
		IncotermsClassification:        data.IncotermsClassification,
		ManualSupplierAddressID:        data.ManualSupplierAddressID,
		AddressName:                    data.AddressName,
		AddressCityName:                data.AddressCityName,
		AddressFaxNumber:               data.AddressFaxNumber,
		AddressPostalCode:              data.AddressPostalCode,
		AddressStreetName:              data.AddressStreetName,
		AddressPhoneNumber:             data.AddressPhoneNumber,
		AddressRegion:                  data.AddressRegion,
		AddressCountry:                 data.AddressCountry,
	}, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Item{
		PurchaseOrder:                  data.PurchaseOrder,
        PurchaseOrderItem:              data.PurchaseOrderItem,
		PurchaseOrderItemText:          data.PurchaseOrderItemText,
		Plant:                          data.Plant,
		StorageLocation:                data.StorageLocation,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		SupplierMaterialNumber:         data.SupplierMaterialNumber,
		OrderQuantity:                  data.OrderQuantity,
		PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
		OrderPriceUnit:                 data.OrderPriceUnit,
		DocumentCurrency:               data.DocumentCurrency,
		NetPriceAmount:                 data.NetPriceAmount,
		NetPriceQuantity:               data.NetPriceQuantity,
		TaxCode:                        data.TaxCode,
		OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		IsFinallyInvoiced:              data.IsFinallyInvoiced,
		PurchaseOrderItemCategory:      data.PurchaseOrderItemCategory,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		InvoiceIsExpected:              data.InvoiceIsExpected,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		Customer:                       data.Customer,
		SupplierIsSubcontractor:        data.SupplierIsSubcontractor,
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		IncotermsClassification:        data.IncotermsClassification,
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		RequisitionerName:              data.RequisitionerName,
		Material:                       data.Material,
		InternationalArticleNumber:     data.InternationalArticleNumber,
		DeliveryAddressID:              data.DeliveryAddressID,
		DeliveryAddressName:            data.DeliveryAddressName,
		DeliveryAddressPostalCode:      data.DeliveryAddressPostalCode,
		DeliveryAddressStreetName:      data.DeliveryAddressStreetName,
		DeliveryAddressCityName:        data.DeliveryAddressCityName,
		DeliveryAddressRegion:          data.DeliveryAddressRegion,
		DeliveryAddressCountry:         data.DeliveryAddressCountry,
		PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
	}, nil
}

func ConvertToAccount(raw []byte, l *logger.Logger) (*Account, error) {
	pm := &responses.Account{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Account. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Account{
		PurchaseOrder:                  data.PurchaseOrder,
        PurchaseOrderItem:              data.PurchaseOrderItem,
		AccountAssignmentNumber:        data.AccountAssignmentNumber,
		GLAccount:                      data.GLAccount,
		BusinessArea:                   data.BusinessArea,
		CostCenter:                     data.CostCenter,
		SalesOrder:                     data.SalesOrder,
		SalesOrderItem:                 data.SalesOrderItem,
		SalesOrderScheduleLine:         data.SalesOrderScheduleLine,
		MasterFixedAsset:               data.MasterFixedAsset,
		FixedAsset:                     data.FixedAsset,
		GoodsRecipientName:             data.GoodsRecipientName,
		UnloadingPointName:             data.UnloadingPointName,
		ControllingArea:                data.ControllingArea,
		CostObject:                     data.CostObject,
		OrderID:                        data.OrderID,
		ProfitCenter:                   data.ProfitCenter,
		WBSElement:                     data.WBSElement,
		ProjectNetwork:                 data.ProjectNetwork,
		FunctionalArea:                 data.FunctionalArea,
		TaxCode:                        data.TaxCode,
		CostCtrActivityType:            data.CostCtrActivityType,
		IsDeleted:                      data.IsDeleted,
	}, nil
}

func ConvertToScheduleLine(raw []byte, l *logger.Logger) (*ScheduleLine, error) {
	pm := &responses.ScheduleLine{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ScheduleLine. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ScheduleLine{
	PurchasingDocument:             data.PurchasingDocument,
    PurchasingDocumentItem:         data.PurchasingDocumentItem,
	ScheduleLine:                   data.ScheduleLine,
	DelivDateCategory:              data.DelivDateCategory,
	ScheduleLineDeliveryDate:       data.ScheduleLineDeliveryDate,
	PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
	ScheduleLineOrderQuantity:      data.ScheduleLineOrderQuantity,
	ScheduleLineDeliveryTime:       data.ScheduleLineDeliveryTime,
	PurchaseRequisition:            data.PurchaseRequisition,
	PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
	ScheduleLineCommittedQuantity:  data.ScheduleLineCommittedQuantity,
	}, nil
}
