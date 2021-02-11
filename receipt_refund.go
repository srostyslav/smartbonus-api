package smartbonus

import (
	"errors"
)

// Item response of refund receipt request
// rest = price * quantity - withdrawn - immediate
type RefundItemResult struct {
	Id 		string 			`json:"id"`			// your product identifier
	Immediate 	float64			`json:"immediate"`		// amount of immediate discount
	Withdrawn 	float64			`json:"withdrawn"`		// amount of withdrawn bonuses	
	Accrued 	float64			`json:"accrued"`		// amount of accrued bonuses
}

// Refund response: list of refund items
type RefundResult []RefundItemResult

// Item of receipt refund request
type RefundItem struct {
	Id 		string 			`json:"nomenclature_id"` 	// your product identifier
	Quantity 	float64			`json:"amount"`			// quantity of product
}

// Receipt refund request body
type ReceiptRefund struct {
	Store
	Id		string 			`json:"refund_id"`		 // identifier of your refund receipt	
	ReceiptId	string 			`json:"remote_id"`		 // identifier of your sell receipt	
	Items		[]RefundItem		`json:"list"`			 // list of refund products
}

// Refund products of receipt
func refundReceipt(storeId string, receipt ReceiptRefund) (*RefundResult, error) {
	var result RefundResult
	if len(receipt.Items) == 0 {
		return &result, errors.New("No item found")
	}

	receipt.StoreId = storeId
	return &result, sendPostRequest(rootPath + "refund/receipt", receipt, &result)
}
