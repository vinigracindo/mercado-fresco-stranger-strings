package domain

type PurchaseOrdersModel struct {
	OrderNumber     string `json:"order_number"`
	OrderDate       string `json:"order_date"`
	TrackingCode    string `json:"tracking_code"`
	BuyerId         int64  `json:"buyer_id"`
	ProductRecordId int64  `json:"product_record_id"`
	OrderStatusId   int64  `json:"order_status_id"`
}
