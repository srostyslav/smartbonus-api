package smartbonus

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestRefundReceipt(t *testing.T) {
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
		Id: uuid.NewV4().String(),
	}

	if _, err := smartbonus.ConfirmReceipt(receipt); err != nil {
		t.Error(err.Error())
		return
	}

	refund := ReceiptRefund{
		Id:        uuid.NewV4().String(),
		ReceiptId: receipt.Id,
		Items: []RefundItem{
			{
				Id:       "2",
				Quantity: 8,
			},
		},
	}

	if _, err := smartbonus.RefundReceipt(refund); err != nil {
		t.Error(err.Error())
	}
}
