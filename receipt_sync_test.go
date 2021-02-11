package smartbonus

import (
	"github.com/satori/go.uuid"
	"testing"
)

func TestSyncReceipt(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "")
	
	receipts := []ReceiptConfirm{
		ReceiptConfirm{
			UserId: testUserId,
			Items: []NomenclatureItem{
				NomenclatureItem{
					Id: "2",
					Quantity: 10,
					Price: 89.65,
				},
				NomenclatureItem{
					Id: "3",
					Quantity: 0.245,
					Price: 23.9,
				},
			},
			Id: uuid.NewV4().String(),
			Discount: 24,
		},
		ReceiptConfirm{
			UserId: testUserId,
			Items: []NomenclatureItem{
				NomenclatureItem{
					Id: "2",
					Quantity: 10,
					Price: 89.65,
				},
				NomenclatureItem{
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
