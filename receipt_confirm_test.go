package smartbonus

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestConfirmReceipt(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "", testCustomerId)

	// If client pay 900 change 2.36 you can accrued to smartbonus account
	receipt := ReceiptConfirm{
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
		Id:     uuid.NewV4().String(),
		Change: 2.36,
	}

	if _, err := smartbonus.ConfirmReceipt(receipt); err != nil {
		t.Error(err.Error())
	}

	receipt = ReceiptConfirm{
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
		Id:       uuid.NewV4().String(),
		Discount: 24, // amount of discount received from DiscountReceipt method.
	}

	if _, err := smartbonus.ConfirmReceipt(receipt); err != nil {
		t.Error(err.Error())
	}
}
