package smartbonus

import (
	"testing"
)

func TestTag(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "", testCustomerId)

	tags := []Tag{
		{Name: "Size", IsGroup: true, Id: "1"},
		{Name: "M", GroupId: "1", Id: "2"},
		{Name: "S", GroupId: "1", Id: "3"},
		{Name: "Red", GroupId: "4", Id: "5"},
		{Name: "Color", IsGroup: true, Id: "4"},
		{Name: "Blue", GroupId: "4", Id: "6"},
		{Name: "Yellow", GroupId: "4", Id: "7"},
	}

	if err := smartbonus.SyncTags(tags); err != nil {
		t.Error(err.Error())
	}
}
