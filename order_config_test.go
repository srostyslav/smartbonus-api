package smartbonus

import (
	"testing"
)

func TestOrderConfig(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "")

	orderUrl, statusUrl := "https://domain:port/api/order", "https://domain:port/api/status"
	if err := smartbonus.ConfigOrder(orderUrl, statusUrl, "really strong token of your store", true, false); err != nil {
		t.Error(err.Error())
	}
}
