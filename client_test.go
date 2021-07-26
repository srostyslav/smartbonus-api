package smartbonus

import (
	"testing"
)

func TestClient(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "")

	if _, err := smartbonus.GetClient(testUserId); err != nil {
		t.Error(err.Error())
	}
}
