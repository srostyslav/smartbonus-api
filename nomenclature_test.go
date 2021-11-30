package smartbonus

import (
	"testing"
)

func TestNomenlcature(t *testing.T) {
	smartbonus := NewSmartBonus(testStoreId, "", testCustomerId)

	nomenclatures := []Product{
		{
			ProductItem: ProductItem{
				ID:           "2",
				Titles:       TextLang{EN: "Yellow shirt"},
				Descriptions: TextLang{EN: "Best quality"},
				Images:       []string{"https://yoursite.com/products/yellow-shift.png"},
				Categories:   []string{"1", "3"}, // Category reference
				Prices: []Price{
					{
						Value:    699.99,
						OldPrice: 900,
						Measure:  PIECE,
						Quantity: 1,
					},
				},
				IsHidden: true,
				CanBuy:   true,
				Brand:    "Google",
				Characteristics: []Characteristic{
					{
						ID: 5,
						Text: &TextLang{
							UK: "S",
						},
					},
				},
			},
		},
	}

	if err := smartbonus.SyncNomenclatures(nomenclatures); err != nil {
		t.Error(err.Error())
	}
}
