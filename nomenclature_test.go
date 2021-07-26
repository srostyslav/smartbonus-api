package smartbonus

import (
	"testing"
)

func TestNomenlcature(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "")

	isHidden := false // enable in smartbonus app
	nomenclatures := []Nomenclature{
		{ // Category
			Id:         "1",
			Name:       "Shirts",
			IsCategory: true,
		},
		{ // Product
			Id:          "2",
			Name:        "Yellow shirt",
			Description: "Best quality",
			Image:       "https://yoursite.com/products/yellow-shift.png",
			CategoryId:  "1", // Category reference
			Price:       699.99,
			Tags:        []string{"3", "7"}, // List of tags
		},
		{ // Product
			Id:         "3",
			Name:       "Blue shirt",
			Image:      "https://yoursite.com/products/blue-shift-back.png,https://yoursite.com/products/blue-shift-front.png",
			CategoryId: "1", // Category reference
			Price:      699.99,
			Tags:       []string{"3", "6"}, // List of tags
		},
		{ // Product
			Id:       "4",
			Name:     "black hat",
			CanBuy:   true,
			IsHidden: &isHidden,
		},
	}

	if err := smartbonus.SyncNomenclatures(nomenclatures); err != nil {
		t.Error(err.Error())
	}
}
