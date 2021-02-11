package smartbonus

import (
	"errors"
)

// Sync many receipts in one request: useful for receipts in offline mode
// warning: ensure that count of elements has to be less than or equal 100 in a request
func syncReceipts(storeId string, receipts []ReceiptConfirm) error {
	if len(receipts) > 100 {
		return errors.New("Length of receipts must be less or equal than 100 elements")
	} else if len(receipts) == 0 {
		return errors.New("No element found")
	}

	var result string
	body := map[string]interface{}{"store": storeId, "elements": receipts}

	if err := sendPostRequest(rootPath + "sync/receipt", body, &result); err != nil {
		return err
	} else if result != "Sync success" {
		return errors.New(result)
	}

	return nil
}
