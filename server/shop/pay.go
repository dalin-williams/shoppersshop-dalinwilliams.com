package shop

import "github.com/satori/go.uuid"

type PayCurrency struct {
	Name		string		`json:"name"`		
	Country 	string		`json:"country"`
	Symbol		string		`json:"symbol"`
	format		string		`json:"format"`
}

type PayObj struct {
	Currency 	PayCurrency		`json:"currency"`
	Payment		float32			`json:"payment"`
}

type Pay interface {
	// Fetches all accepted currencies
	// RETURNS: a list of valid currencies
	FetchAcceptedCurrencies()(payCurrencies []PayCurrency, err error)

	// Adds an item to the session cart
	AddToCart(productId uuid.UUID)(err error)

	// Fetches all items in the current cart
	// RETURNS: a list of items in cart
	FetchCurrentCart()(inventoryInCart []InventoryObj, err error)

	// Deletes an item from the cart
	RemoveFromCart(inventoryId uuid.UUID)(err error)

	// Creates a purchase transaction by "buying" all items in the current cart
	// RETURNS: a callback uri
	PurchaseCart(payment PayObj)(callback string, err error)

	// Populates the existing cart with a previous order's items
	ReOrderOrder(orderId uuid.UUID)(err error)

	// Updates an existing order IF it is not completed
	// RETURNS: a callback uri
	ModifyOrder(orderId uuid.UUID, modifiedVend VendObj)(callback string, err error)
}