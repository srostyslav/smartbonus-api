package smartbonus

import (
    "errors"
)

// Nomenclature instance has to be sync to smartbonus after it created, changed or deleted.
// If you cannot trigger nomenclature events, send it by some interval: once a day for example.
type Nomenclature struct {
    Id              	string      	`json:"id"`                             // unique identifier of product in your db
    Name 		string 		`json:"name"`				// title of product
    Description		string 		`json:"description,omitempty"`		// description of product (optional)
    Image 		string 		`json:"photo_url,omitempty"`		// image of product, if you have more than one image join them by comma (optional)
    IsDeleted		bool 		`json:"is_deleted,omitempty"`		// send true if you deleted that product (optional)
    CategoryId		string		`json:"category,omitempty"`		// unique identifier of category in your db (optional)
    Barcode		string 		`json:"barcode,omitempty"`		// barcode of your product (optional)
    Price		float64		`json:"price,omitempty"`		// price of your product (optional)
    IsCategory		bool		`json:"is_category,omitempty"`		// send true if current instance is category or false if it's product (optional)
    Tags		[]string	`json:"tags,omitempty"`			// list of tag identifiers (optional)
    CanBuy		bool		`json:"can_buy,omitempty"`		// send true if this product can be buyed in smartbonus app (optional)
    IsHidden 		*bool		`json:"is_hidden,omitempty"`		// send false if you want to show this product in smartbonus app catalog for clients (optional)
}

// Sync your catalog to smartbonus app
// warning: ensure that count of elements has to be less than or equal 500 in a request
func syncNomenclatures(storeId string, nomes []Nomenclature) error {
    if len(nomes) > 500 {
        return errors.New("Length of nomenclatures must be less or equal than 500 elements")
    } else if len(nomes) == 0 {
        return errors.New("No element found")
    }

    var result string
    body := map[string]interface{}{"store": storeId, "elements": nomes}

    if err := sendPostRequest(rootPath + "sync/nomenclature", body, &result); err != nil {
        return err
    } else if result != "Sync success" {
        return errors.New(result)
    }

    return nil
}
