package sap_api_caller

type PurchaseOrder struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	PurchaseOrder string `json:"purchase_order"`
	Deleted       string `json:"deleted"`
}    
    
type PurchaseOrderHeader struct {
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
    ExchangeRate                float64     `json:"ExchangeRate"`
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
}

type PurchaseOrderItem           struct {
	PurchaseOrder                  string `json:"PurchaseOrder"`
    PurchaseOrderItem              int    `json:"PurchaseOrderItem"`
    PurchaseOrderItemText          string `json:"PurchaseOrderItemText"`
    Plant                          string `json:"Plant"`
    StorageLocation                string `json:"StorageLocation"`
    MaterialGroup                  string `json:"MaterialGroup"`
    PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
    SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
    OrderQuantity                  float64 `json:"OrderQuantity"`
    PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
    OrderPriceUnit                 string `json:"OrderPriceUnit"`
    DocumentCurrency               string `json:"DocumentCurrency"`
    NetPriceAmount                 float64 `json:"NetPriceAmount"`
    NetPriceQuantity               float64 `json:"NetPriceQuantity"`
    ScheduleLineDeliveryDate       string `json:"ScheduleLineDeliveryDate"`
    TaxCode                        string `json:"TaxCode"`
    OverdelivTolrtdLmtRatioInPct   float64 `json:"OverdelivTolrtdLmtRatioInPct"`
    UnlimitedOverdeliveryIsAllowed string `json:"UnlimitedOverdeliveryIsAllowed"`
    UnderdelivTolrtdLmtRatioInPct  float64 `json:"UnderdelivTolrtdLmtRatioInPct"`
    IsCompletelyDelivered          string `json:"IsCompletelyDelivered"`
    IsFinallyInvoiced              string `json:"IsFinallyInvoiced"`
    PurchaseOrderItemCategory      string `json:"PurchaseOrderItemCategory"`
    AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
    GoodsReceiptIsExpected         string `json:"GoodsReceiptIsExpected"`
    GoodsReceiptIsNonValuated      string `json:"GoodsReceiptIsNonValuated"`
    InvoiceIsExpected              string `json:"InvoiceIsExpected"`
    InvoiceIsGoodsReceiptBased     string `json:"InvoiceIsGoodsReceiptBased"`
    Customer                       string `json:"Customer"`
    SupplierIsSubcontractor        string `json:"SupplierIsSubcontractor"`
    ItemNetWeight                  float64 `json:"ItemNetWeight"`
    ItemWeightUnit                 string `json:"ItemWeightUnit"`
    IncotermsClassification        string `json:"IncotermsClassification"`
    PurchaseRequisition            string `json:"PurchaseRequisition"`
    PurchaseRequisitionItem        int    `json:"PurchaseRequisitionItem"`
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
    PurchasingDocumentDeletionCode string `json:"PurchasingDocumentDeletionCode"`
}

type AccountAssignmentNumber        struct {
	PurchaseOrder                  string `json:"PurchaseOrder"`
    PurchaseOrderItem              int    `json:"PurchaseOrderItem"`
    AccountAssignmentNumber        int    `json:"AccountAssignmentNumber"`
    GLAccount                      string `json:"GLAccount"`
    BusinessArea                   string `json:"BusinessArea"`
    CostCenter                     string `json:"CostCenter"`
    SalesOrder                     string `json:"SalesOrder"`
    SalesOrderItem                 int    `json:"SalesOrderItem"`
    SalesOrderScheduleLine         int    `json:"SalesOrderScheduleLine"`
    MasterFixedAsset               string `json:"MasterFixedAsset"`
    FixedAsset                     string `json:"FixedAsset"`
    GoodsRecipientName             string `json:"GoodsRecipientName"`
    UnloadingPointName             string `json:"UnloadingPointName"`
    ControllingArea                string `json:"ControllingArea"`
    CostObject                     string `json:"CostObject"`
    OrderID                        string `json:"OrderID"`
    ProfitCenter                   string `json:"ProfitCenter"`
    WBSElement                     int         `json:"WBSElement"`
    ProjectNetwork                 string `json:"ProjectNetwork"`
    FunctionalArea                 string `json:"FunctionalArea"`
    TaxCode                        string `json:"TaxCode"`
    CostCtrActivityType            string `json:"CostCtrActivityType"`
    IsDeleted                      string`json:"IsDeleted"`
}