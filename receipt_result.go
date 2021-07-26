package smartbonus

// Smartbonus modules that accrued/withdrawn bonuses or added discount
type ExecutedModule struct {
	Id         string  `json:"id"`              // identifier of smartbonus module
	Type       string  `json:"type"`            // type of object: subscriber_module|discount_module|coupon_module
	Accrued    float64 `json:"accrued_bonus"`   // amount of accrued bonuses
	Immediate  float64 `json:"immediate_bonus"` // amount of immediate discount
	Withdrawn  float64 `json:"withdrawn_bonus"` // amount of withdrawn bonuses
	ModuleType string  `json:"module_type"`     // type of smartbonus module
	Name       string  `json:"name"`            // title of module
}

// List of executed modules
type AnalyticObject struct {
	ExecutedModules []ExecutedModule `json:"executed_modules"`
}

// Item of receipt response
type ReceiptItems struct {
	Id        string  `json:"id"`         // your product identifier
	Accrued   float64 `json:"accrued"`    // amount of accrued bonuses
	Withdrawn float64 `json:"withdrawn"`  // amount of withdrawn bonuses
	Immediate float64 `json:"immediate"`  // amount of immediate discount
	Quantity  float64 `json:"amount"`     // quantity of product
	Price     float64 `json:"unit_price"` // price of product
}

// Receipt response of discount and confirm methods
type ReceiptResult struct {
	Discount        float64        `json:"discount"`         // amount of discount: withdrawn + immediate
	Info            string         `json:"info"`             // description of modules
	Accrued         float64        `json:"user_add_bonus"`   // amount of accrued bonuses
	Withdrawn       float64        `json:"withdrawn"`        // amount of withdrawn bonuses
	Immediate       float64        `json:"immediate"`        // amount of immediate discount
	UserName        string         `json:"user_name"`        // client name
	Items           []ReceiptItems `json:"nomenclatures"`    // list of products
	AnalyticObjects AnalyticObject `json:"analytics_object"` // details of smartbonus modules
}
