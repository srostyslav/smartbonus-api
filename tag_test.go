package smartbonus

import (
	"testing"
)

func TestTag(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "")
	
	tags := []Tag{
		Tag{Name: "Size", IsGroup: true, Id: "1"},
		Tag{Name: "M", GroupId: "1", Id: "2"},
		Tag{Name: "S", GroupId: "1", Id: "3"},
		Tag{Name: "Red", GroupId: "4", Id: "5"},
		Tag{Name: "Color", IsGroup: true, Id: "4"},
		Tag{Name: "Blue", GroupId: "4", Id: "6"},
		Tag{Name: "Yellow", GroupId: "4", Id: "7"},
	}
	
	if err := smartbonus.SyncTags(tags); err != nil {
		t.Error(err.Error())
	}
}
