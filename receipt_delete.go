package smartbonus

import (
	"errors"
)

// Delete previous confirmed receipts
func deleteReceipts(storeId string, receipts []string) error {
	if len(receipts) > 100 {
		return errors.New("length of receipts must be less or equal than 100 elements")
	} else if len(receipts) == 0 {
		return errors.New("no element found")
	}

	elements := []map[string]string{}
	for _, receipt := range receipts {
		elements = append(elements, map[string]string{"remote_id": receipt})
	}

	var result string
	body := map[string]interface{}{"store": storeId, "elements": elements}

	if err := sendPostRequest(rootPath+"v2/delete/receipt", body, &result); err != nil {
		return err
	} else if result != "Delete success" {
		return errors.New(result)
	}

	return nil
}
