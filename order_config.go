package smartbonus

import (
	"errors"
	"fmt"
)

// Config order hook: 
// when user created order in smartbonus we send it to your api orderUrl with body Order struct.
// 1. orderUrl - your api that wait new order from smartbonus app
// 2. statusUrl - your api that wait new status of order: receive body StatusBody
// 3. token - your unique identifier of store: your receive it in orderUrl & statusUrl hooks in field "store"
func configOrder(storeId, orderUrl, statusUrl, token string) error {
	var result interface{}
	body := map[string]interface{}{"store": storeId, "order_url": orderUrl, "status_url": statusUrl, "token": token}

	if err := sendPostRequest(rootPath + "order/config", body, &result); err != nil {
		return err
	} else if result != nil {
		return errors.New(fmt.Sprintf("%v", result))
	}

	return nil
}
