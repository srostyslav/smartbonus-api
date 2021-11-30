package smartbonus

import (
	"fmt"
)

// Config order hook:
// when user created order in smartbonus we send it to your api orderUrl with body Order struct.
// 1. orderUrl - your api that wait new order from smartbonus app
// 2. statusUrl - your api that wait new status of order: receive body StatusBody
// 3. token - your unique identifier of store: your receive it in orderUrl & statusUrl hooks in field "store"
// 4. syncNomenclatureFull - sign that sync nomenclature object full
// 5. syncNomenclatureByCustomer - sign that sync nomenclature by customer
func configOrder(storeId, orderUrl, statusUrl, token string, syncNomenclatureFull, syncNomenclatureByCustomer bool) error {
	var result interface{}
	body := map[string]interface{}{
		"store":                         storeId,
		"order_url":                     orderUrl,
		"status_url":                    statusUrl,
		"token":                         token,
		"sync_nomenclature_full":        syncNomenclatureFull,
		"sync_nomenclature_by_customer": syncNomenclatureByCustomer,
	}

	if err := sendPostRequest(rootPath+"v2/order/config", body, &result); err != nil {
		return err
	} else if result != nil {
		return fmt.Errorf("%v", result)
	}

	return nil
}
