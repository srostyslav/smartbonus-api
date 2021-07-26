package smartbonus

import (
	"testing"
	"time"
)

func TestDiscountReceipt(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "")

	receipt := ReceiptDiscount{
		UserId: testUserId,
		Items: []NomenclatureItem{
			{
				Id:       "2",
				Quantity: 10,
				Price:    89.65,
			},
			{
				Id:       "3",
				Quantity: 0.245,
				Price:    23.9,
			},
		},
		Date: time.Now().Unix(),
	}

	if _, err := smartbonus.DiscountReceipt(receipt); err != nil {
		t.Error(err.Error())
	}
}
