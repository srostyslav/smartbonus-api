package smartbonus

import (
	"errors"
)

type NomenclatureItem struct {
	Id       string  `json:"nomenclature_id"` // your product identifier
	Quantity float64 `json:"amount"`          // quantity or product
	Price    float64 `json:"unit_price"`      // price of product
	PosterID string  `json:"poster_id"`       // the nomenlautre is related to the store, not to the customer
}

// Body for receipt discount method
type ReceiptDiscount struct {
	Store
	UserId     string                 `json:"user_id"`             // Phone or sanned key from smartbonus app
	Date       int64                  `json:"date,omitempty"`      // Date of receipt (optional)
	Withdrawn  float64                `json:"withdrawn,omitempty"` // Amount of money that cashier want to withdraw from client account (optional)
	Items      []NomenclatureItem     `json:"receipt"`             // products of receipt
	Additional map[string]interface{} `json:"additional"`
}

// Get discount of receipt
func discountReceipt(storeId string, receipt ReceiptDiscount) (*ReceiptResult, error) {
	var result ReceiptResult
	if len(receipt.Items) == 0 {
		return &result, errors.New("no item found")
	}

	receipt.StoreId = storeId
	return &result, sendPostRequest(rootPath+"v2/receipt/discount", receipt, &result)
}
