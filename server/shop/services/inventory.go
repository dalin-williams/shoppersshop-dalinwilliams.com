package services

import(
	"github.com/funkeyfreak/vending-machine-api/server/shop"
)

type inventory struct {
	services map[string]shop.Service
}


func (i *inventory) FindItemByName(query shop.InventoryQuery)(inventoryItems []shop.InventoryObj, err error){

}

func (i *inventory) FetchAllCategoriesAndStores() (inventoryCategories []shop.InventoryDetails, err error) {
	for key, value := range i.services {
		var t shop.InventoryDetails
		t.Name = key



	}
}

func (i *inventory) FetchAllInventoryItems(pagination shop.InventoryPagination) (inventoryItems []shop.InventoryObj, err error) {

}
