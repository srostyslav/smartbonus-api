package smartbonus

import (
	"fmt"
)

// List of order statuses
type OrderStatusType int

const (
	NEW OrderStatusType = iota
	PAYMENT_PENDING
	PAYMENT_CANCELED
	PROCESSING
	AWAITING_SHIPMENT
	AWAITING_PICKUP
	COMPLETED
	CANCELED
	REFUNDED
	AWAITING_WEB_PAYMENT
	WEB_PAYMENT_SUCCESSFUL
	_
	AWAITING_FOR_COLLECT
	COLLECTING
	TRANSFERRED_FOR_DELIVERY
	DELIVERING
	CANCEL_REQUEST
	AWAITING_POSTPAYMENT
)

func (OrderStatusType) getTitles() []string {
	return []string{"new", "payment_pending", "payment_canceled", "processing", "awaiting_shipment", "awaiting_pickup", "completed", "canceled",
		"refunded", "awaiting_web_payment", "web_payment_successful", "", "awaiting_for_collect", "collecting", "transferred_for_delivery",
		"delivering", "cancel_request", "awaiting_postpayment"}
}

func (s OrderStatusType) Get(status string) (OrderStatusType, error) {
	for i, st := range s.getTitles() {
		if st == status {
			return OrderStatusType(i), nil
		}
	}
	return 0, fmt.Errorf("status is not found '%s'", status)
}

func (a OrderStatusType) String() string {
	return a.getTitles()[a]
}

// Body of status:
type StatusBody struct {
	Store
	OrderId string          `json:"order_id"` // identifier of order in smartbonus
	Status  OrderStatusType `json:"status"`   //  one of OrderStatuses
}

// Change status of order that created in smartbonus app
// If status changed client receive push notification about it
func changeOrderStatus(storeId string, statusBody StatusBody) error {
	var result interface{}
	statusBody.StoreId = storeId

	if err := sendPostRequest(rootPath+"v2/order/status", statusBody, &result); err != nil {
		return err
	} else if result != nil {
		return fmt.Errorf("%v", result)
	}

	return nil
}
