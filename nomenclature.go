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

// Nomenclature instance has to be sync to smartbonus after it created, changed or deleted.
// If you cannot trigger nomenclature events, send it by some interval: once a day for example.
type Nomenclature struct {
	Id          string   `json:"id"`                    // unique identifier of product in your db
	Name        string   `json:"name"`                  // title of product
	Description string   `json:"description,omitempty"` // description of product (optional)
	Image       string   `json:"photo_url,omitempty"`   // image of product, if you have more than one image join them by comma (optional)
	IsDeleted   bool     `json:"is_deleted,omitempty"`  // send true if you deleted that product (optional)
	CategoryId  string   `json:"category,omitempty"`    // unique identifier of category in your db (optional)
	Barcode     string   `json:"barcode,omitempty"`     // barcode of your product (optional)
	Price       float64  `json:"price,omitempty"`       // price of your product (optional)
	IsCategory  bool     `json:"is_category,omitempty"` // send true if current instance is category or false if it's product (optional)
	Tags        []string `json:"tags,omitempty"`        // list of tag identifiers (optional)
	CanBuy      bool     `json:"can_buy,omitempty"`     // send true if this product can be buyed in smartbonus app (optional)
	IsHidden    *bool    `json:"is_hidden,omitempty"`   // send false if you want to show this product in smartbonus app catalog for clients (optional)
	IsWeight    bool     `json:"is_weight"`             // send true if nomenclature item is weight product.
	Multiplier  float64  `json:"multiplier"`            // send 0.01 price at 0.1 kg & 0.001 at 1 kg
	Quantity    float64  `json:"quantity"`              // product quantity
	OldPrice    float64  `json:"old_price"`             // previous price that identify discount of product.
	Priority    int64    `json:"priority"`              // priority in apps
	Articul     string   `json:"articul"`               // articul of product
}

// Sync your catalog to smartbonus app
// warning: ensure that count of elements has to be less than or equal 500 in a request
func syncNomenclaturesV2(storeId string, nomes []Nomenclature) error {
	if len(nomes) > 500 {
		return errors.New("length of nomenclatures must be less or equal than 500 elements")
	} else if len(nomes) == 0 {
		return errors.New("no element found")
	}

	var result string
	body := map[string]interface{}{"store": storeId, "elements": nomes}

	if err := sendPostRequest(rootPath+"v2/sync/nomenclature", body, &result); err != nil {
		return err
	} else if result != "Sync success" {
		return errors.New(result)
	}

	return nil
}
