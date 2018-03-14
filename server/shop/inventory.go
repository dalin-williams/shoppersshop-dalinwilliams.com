package shop

import (
	"github.com/satori/go.uuid"
)

type InventoryQuery struct {
	Categories []CategoriesObj 	`json:"categories"`
	Name		string			`json:"name"`
	Price		float64			`json:"price"`
}

type InventoryDetails struct {
	Name 		string 			`json:"name"`
	Categories 	[]CategoriesObj	`json:"categories"`
}

type InventoryObj struct {
	 Id 		uuid.UUID 			`json:"id"`
	 ASIN 		string				`json:"asin"`
	 Name 		string				`json:"name"`
	 Cost		PayObj				`json:"cost"`
	 Url		string				`json:"url"`
	 Resources 	InventoryResource	`json:"resources"`
}

type InventoryResource struct {
	Images	[]string 		`json:"images"`
}

	type InventoryPagination struct {
	PaginationUri	string	`json:"pagination_uri"`
}

type Inventory interface {
	// Gets an item by Inventory Query object
	FindItemByName(query InventoryQuery)(inventoryItems []InventoryObj, err error)

	// Gets all categories and stores
	FetchAllCategoriesAndStores() (inventoryCategories []InventoryDetails, err error)

	// Fetched all inventory items *MUST ENFORCE PAGINATION
	//TODO: accept multiple URIs
	FetchAllInventoryItems(pagination InventoryPagination) (inventoryItems []InventoryObj, err error)
}

