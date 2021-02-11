package smartbonus

// OrderProduct instance - element of products in Order:
type OrderProduct struct {
	Id 			string 			`json:"id"`				// your nomenclature identifier
	Price 			float64			`json:"amount"`				// price of product
	Quantity 		float64			`json:"quantity"`   			// quantity of product   
}

// OrderStatus instance - element of statuses in Order:
type OrderStatus struct {
	Date 			int64 			`json:"date_unix"`			// date of status creation
	Status 			uint 			`json:"status"`				// one of OrderStatuses
}

// Order instance - send new order that created in smartbonus to your api after webhook is configured:
type Order struct {
	Store											// your StoreId token that configured
	Id 			string 			`json:"remote_id"`			// unique identifier of order in smartbonus
	Code 			string 			`json:"code"`				// number of order for client
	UserId 			string 			`json:"user_id"`			// client identifier that has to be used to sync receipt
	Phone 			string 			`json:"phone"`				// phone number of client
	UserName 		string 			`json:"user_name"`			// name of client
	Amount 			float64			`json:"amount"`				// amount for pay
	Currency 		string 			`json:"currency"`			// ISO 4217: UAH, USD, EUR
	Date 			int64 			`json:"date_unix"`			// date of order
	IsPaid 			bool 			`json:"is_paid"`			// order paid online by client
	ProductsAmount		float64			`json:"products_amount"`		// amount of products
	DeliveryCost		float64			`json:"delivery_cost"`			// cost of delivery
	Discount 		float64			`json:"discount"`			// amount of discount
	Products 		[]OrderProduct		`json:"products"`
	Statuses 		[]OrderStatus		`json:"statuses"`
	Comment 		string 			`json:"comment"`			// client comment
	DeliveryType 		string 			`json:"delivery"`			// type of delivery
	DeliveryAddress		string 			`json:"deliveryAddress"`
	DeliveryTime		string 			`json:"deliveryTime"`
}
