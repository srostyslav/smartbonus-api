package smartbonus

import (
	"errors"
)

type MeasureType string

const (
	PIECE    MeasureType = "piece"
	KILOGRAM MeasureType = "kilogram"
	GRAM     MeasureType = "gram"
)

type TextLang struct {
	UK string `json:"uk"`
	RU string `json:"ru"`
	EN string `json:"en"`
}

type Price struct {
	Value    float64     `json:"value"`
	StoreID  string      `json:"store_id,omitempty"`
	OldPrice float64     `json:"old_price,omitempty"`
	Measure  MeasureType `json:"measure"`
	Quantity float64     `json:"quantity,omitempty"`
}

type Characteristic struct {
	ID   int64     `json:"id"`
	Num  *float64  `json:"num"`
	Text *TextLang `json:"text"`
}

type ProductItem struct {
	ID              string           `json:"id"`
	Titles          TextLang         `json:"titles"`
	Descriptions    TextLang         `json:"descriptions"`
	Images          []string         `json:"images"`
	Categories      []string         `json:"categories"`
	Priority        int64            `json:"priority"`
	Barcode         string           `json:"barcode"`
	IsHidden        bool             `json:"is_hidden"`
	CanBuy          bool             `json:"can_buy"`
	Prices          []Price          `json:"prices"`
	Characteristics []Characteristic `json:"characteristics"`
	Brand           string           `json:"brand"`
}

type Product struct {
	ProductItem
	Modifiers []ProductItem `json:"modifiers"`
}

// Sync your catalog to smartbonus app
// warning: ensure that count of elements has to be less than or equal 500 in a request
func syncNomenclatures(customerId string, nomes []Product) error {
	if len(nomes) > 500 {
		return errors.New("length of nomenclatures must be less or equal than 500 elements")
	} else if len(nomes) == 0 {
		return errors.New("no element found")
	}

	var result string
	body := map[string]interface{}{"customer_id": customerId, "elements": nomes}

	if err := sendPostRequest(rootPath+"v3/sync/nomenclature", body, &result); err != nil {
		return err
	} else if result != "Sync success" {
		return errors.New(result)
	}

	return nil
}
