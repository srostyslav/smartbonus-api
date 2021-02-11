package smartbonus

import (
	"errors"
)

// Body for receipt confirmation
type ReceiptConfirm struct {
	Store
	UserId 			string 			`json:"user_id"`		// Phone or sanned key from smartbonus app
	Date 			int64 			`json:"date,omitempty"`		// Date of receipt (optional)
	Discount		float64			`json:"discount"`		// Amount of discount that received from DiscountReceipt method.
	Items			[]NomenclatureItem	`json:"list"`			// List of products
	Id			string 			`json:"remote_id"`		// Unique receipt identifier
	Change 			float64			`json:"accrued,omitempty"`	// Rest of money that will accrue to smartbonus account
}

// Confirmation of receipt
func confirmReceipt(storeId string, receipt ReceiptConfirm) (*ReceiptResult, error) {
	var result ReceiptResult
	if len(receipt.Items) == 0 {
		return &result, errors.New("No item found")
	}

	receipt.StoreId = storeId
	return &result, sendPostRequest(rootPath + "receipt/confirm", receipt, &result)
}
