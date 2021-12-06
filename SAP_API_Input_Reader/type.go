package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	PurchaseOrder struct {
		PurchaseOrder             string      `json:"document_no"`
		Plant                     string      `json:"deliver_to"`
		OrderQuantity             string      `json:"quantity"`
		PickedQuantity            string      `json:"picked_quantity"`
		NetPriceAmount            string      `json:"price"`
	    Batch                     string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                         string `json:"api_schema"`
	MaterialCode                      string `json:"material_code"`
	Supplier                          string `json:"plant/supplier"`
	Stock                             string  `json:"stock"`
	PurchaseOrderType                 string `json:"document_type"`
	PurchaseOrder                     string `json:"document_no"`
	ScheduleLineDeliveryDate          string `json:"planned_date"`
	Validated Date                    string `json:"validated_date"`
	Deleted                           bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	PurchaseOrder struct {
		PurchaseOrder               string      `json:"PurchaseOrder"`
		CompanyCode                 string      `json:"CompanyCode"`
		PurchaseOrderType           string      `json:"PurchaseOrderType"`
		PurchasingProcessingStatus  string      `json:"PurchasingProcessingStatus"`
		CreationDate                string      `json:"CreationDate"`
		LastChangeDateTime          string      `json:"LastChangeDateTime"`
		Supplier                    string      `json:"Supplier"`
		Language                    string      `json:"Language"`
		PaymentTerms                string      `json:"PaymentTerms"`
		PurchasingOrganization      string      `json:"PurchasingOrganization"`
		PurchasingGroup             string      `json:"PurchasingGroup"`
		PurchaseOrderDate           string      `json:"PurchaseOrderDate"`
		DocumentCurrency            string      `json:"DocumentCurrency"`
		ExchangeRate                string      `json:"ExchangeRate"`
		ValidityStartDate           string      `json:"ValidityStartDate"`
		ValidityEndDate             string      `json:"ValidityEndDate"`
		SupplierRespSalesPersonName string      `json:"SupplierRespSalesPersonName"`
		SupplierPhoneNumber         string      `json:"SupplierPhoneNumber"`
		SupplyingPlant              string      `json:"SupplyingPlant"`
		IncotermsClassification     string      `json:"IncotermsClassification"`
		ManualSupplierAddressID     string      `json:"ManualSupplierAddressID"`
		AddressName                 string      `json:"AddressName"`
		AddressCityName             string      `json:"AddressCityName"`
		AddressFaxNumber            string      `json:"AddressFaxNumber"`
		AddressPostalCode           string      `json:"AddressPostalCode"`
		AddressStreetName           string      `json:"AddressStreetName"`
		AddressPhoneNumber          string      `json:"AddressPhoneNumber"`
		AddressRegion               string      `json:"AddressRegion"`
		AddressCountry              string      `json:"AddressCountry"`
		PurchaseOrderItem           struct {
			PurchaseOrderItem              string `json:"PurchaseOrderItem"`
			PurchaseOrderItemText          string `json:"PurchaseOrderItemText"`
			Plant                          string `json:"Plant"`
			StorageLocation                string `json:"StorageLocation"`
			MaterialGroup                  string `json:"MaterialGroup"`
			PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
			SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
			OrderQuantity                  string `json:"OrderQuantity"`
			PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
			OrderPriceUnit                 string `json:"OrderPriceUnit"`
			DocumentCurrency               string `json:"DocumentCurrency"`
			NetPriceAmount                 string `json:"NetPriceAmount"`
			NetPriceQuantity               string `json:"NetPriceQuantity"`
			ScheduleLineDeliveryDate       string `json:"ScheduleLineDeliveryDate"`
			TaxCode                        string `json:"TaxCode"`
			OverdelivTolrtdLmtRatioInPct   string `json:"OverdelivTolrtdLmtRatioInPct"`
			UnlimitedOverdeliveryIsAllowed string `json:"UnlimitedOverdeliveryIsAllowed"`
			UnderdelivTolrtdLmtRatioInPct  string `json:"UnderdelivTolrtdLmtRatioInPct"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			IsFinallyInvoiced              bool   `json:"IsFinallyInvoiced"`
			PurchaseOrderItemCategory      string `json:"PurchaseOrderItemCategory"`
			AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
			GoodsReceiptIsExpected         bool   `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      bool   `json:"GoodsReceiptIsNonValuated"`
			InvoiceIsExpected              bool   `json:"InvoiceIsExpected"`
			InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
			Customer                       string `json:"Customer"`
			SupplierIsSubcontractor        string `json:"SupplierIsSubcontractor"`
			ItemNetWeight                  string `json:"ItemNetWeight"`
			ItemWeightUnit                 string `json:"ItemWeightUnit"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			PurchaseRequisition            string `json:"PurchaseRequisition"`
			PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
			RequisitionerName              string `json:"RequisitionerName"`
			Material                       string `json:"Material"`
			InternationalArticleNumber     string `json:"InternationalArticleNumber"`
			DeliveryAddressID              string `json:"DeliveryAddressID"`
			DeliveryAddressName            string `json:"DeliveryAddressName"`
			DeliveryAddressPostalCode      string `json:"DeliveryAddressPostalCode"`
			DeliveryAddressStreetName      string `json:"DeliveryAddressStreetName"`
			DeliveryAddressCityName        string `json:"DeliveryAddressCityName"`
			DeliveryAddressRegion          string `json:"DeliveryAddressRegion"`
			DeliveryAddressCountry         string `json:"DeliveryAddressCountry"`
			PurchasingDocumentDeletionCode bool   `json:"PurchasingDocumentDeletionCode"`
			AccountAssignmentNumber        struct {
				AccountAssignmentNumber string `json:"AccountAssignmentNumber"`
				GLAccount               string `json:"GLAccount"`
				BusinessArea            string `json:"BusinessArea"`
				CostCenter              string `json:"CostCenter"`
				SalesOrder              string `json:"SalesOrder"`
				SalesOrderItem          string `json:"SalesOrderItem"`
				SalesOrderScheduleLine  string `json:"SalesOrderScheduleLine"`
				MasterFixedAsset        string `json:"MasterFixedAsset"`
				FixedAsset              string `json:"FixedAsset"`
				GoodsRecipientName      string `json:"GoodsRecipientName"`
				UnloadingPointName      string `json:"UnloadingPointName"`
				ControllingArea         string `json:"ControllingArea"`
				CostObject              string `json:"CostObject"`
				OrderID                 string `json:"OrderID"`
				ProfitCenter            string `json:"ProfitCenter"`
				WBSElement              string `json:"WBSElement"`
				ProjectNetwork          string `json:"ProjectNetwork"`
				FunctionalArea          string `json:"FunctionalArea"`
				TaxCode                 string `json:"TaxCode"`
				CostCtrActivityType     string `json:"CostCtrActivityType"`
				IsDeleted               bool   `json:"IsDeleted"`
			} `json:"AccountAssignmentNumber"`
		} `json:"PurchaseOrderItem"`
	} `json:"PurchaseOrder"`
	APISchema     string `json:"api_schema"`
	PONumber      string `json:"purchase_order"`
	Deleted       bool   `json:"deleted"`
}
