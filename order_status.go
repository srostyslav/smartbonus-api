package smartbonus

import (
	"errors"
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
)

// Body of status:
type StatusBody struct {
	Store
	OrderId		string 			`json:"order_id"`	// identifier of order in smartbonus
	Status		OrderStatusType `json:"status"`		//  one of OrderStatuses
}

// Change status of order that created in smartbonus app
// If status changed client receive push notification about it
func changeOrderStatus(storeId string, statusBody StatusBody) error {
	var result interface{}
	statusBody.StoreId = storeId

	if err := sendPostRequest(rootPath + "order/status", statusBody, &result); err != nil {
		return err
	} else if result != nil {
		return errors.New(fmt.Sprintf("%v", result))
	}

	return nil
}
