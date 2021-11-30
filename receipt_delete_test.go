package smartbonus

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestDeleteReceipt(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "", testCustomerId)

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

	if err := smartbonus.DeleteReceipts([]string{receipt.Id}); err != nil {
		t.Error(err.Error())
	}
}
