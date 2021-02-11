package smartbonus

import (
	"errors"
	"fmt"
)

// List of order statuses
var OrderStatuses map[uint]string = map[uint]string{
	0: "new",
	1: "payment_pending",
	2: "payment_canceled",
	3: "processing",
	4: "awaiting_shipment",
	5: "awaiting_pickup",
	6: "completed",
	7: "canceled",
	8: "refunded",
}

// Body of status:
type StatusBody struct {
	Store
	OrderId		string 			`json:"order_id"`	// identifier of order in smartbonus
	Status		uint 			`json:"status"`		//  one of OrderStatuses
}

// Change status of order that created in smartbonus app
// If status changed client receive push notification about it
func changeOrderStatus(storeId string, statusBody StatusBody) error {
	if _, ok := OrderStatuses[statusBody.Status]; !ok {
		return errors.New(fmt.Sprintf("Status %d does not exist", statusBody.Status))
	}

	var result interface{}
	statusBody.StoreId = storeId

	if err := sendPostRequest(rootPath + "order/status", statusBody, &result); err != nil {
		return err
	} else if result != nil {
		return errors.New(fmt.Sprintf("%v", result))
	}

	return nil
}
