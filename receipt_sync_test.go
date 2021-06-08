package smartbonus

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestSyncReceipt(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "")
	
	receipts := []ReceiptConfirm{
		{
			UserId: testUserId,
			Items: []NomenclatureItem{
				{
					Id: "2",
					Quantity: 10,
					Price: 89.65,
				},
				{
					Id: "3",
					Quantity: 0.245,
					Price: 23.9,
				},
			},
			Id: uuid.NewV4().String(),
			Discount: 24,
		},
		{
			UserId: testUserId,
			Items: []NomenclatureItem{
				{
					Id: "2",
					Quantity: 10,
					Price: 89.65,
				},
				{
					Id: "3",
					Quantity: 0.245,
					Price: 23.9,
				},
			},
			Id: uuid.NewV4().String(),
			Discount: 24,
		},
	}

	if err := smartbonus.SyncReceipts(receipts); err != nil {
		t.Error(err.Error())
	}
}
