package smartbonus

// Client instance is smartbonus app user
type Client struct {
	Phone   string  `json:"phone"`   // phone number of client (unique)
	Balance float64 `json:"balance"` // amount of bonuses in smartbonus account
	Name    string  `json:"name"`    // client name (optional)
}

// If client exists in smartbonus app return its instance
// userId - phone or scanned key from smartbonus app
func getClient(storeId, userId string) (*Client, error) {
	var client Client
	params := map[string]string{"store": storeId, "user_id": userId}

	return &client, sendGetRequest(rootPath+"user/phone", params, &client)
}
