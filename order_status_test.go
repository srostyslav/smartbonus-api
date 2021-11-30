package smartbonus

import (
	"testing"
)

func TestOrderStatus(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "", testCustomerId)

	statusBody := StatusBody{
		OrderId: "fce887b6-b307-cc0f-309d-933db16e406b",
		Status:  PROCESSING,
	}
	if err := smartbonus.ChangeOrderStatus(statusBody); err != nil {
		t.Error(err.Error())
	}
}
