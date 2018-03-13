package amazon

import (

	"github.com/dominicphillips/amazing"

	"github.com/funkeyfreak/vending-machine-api/server/shop"
)

type Amazon struct {
	AmazonAccessName string `json:"amazon_access_name"`
	AmazonAccessKey  string `json:"amazon_access_key"`
	AmazonSecretKey	 string `json:"amazon_secret_key"`

	client *amazing.Amazing
}



func (a *Amazon) newClient() error {
	if a.client != nil {
		return nil
	}
	client, err := amazing.NewAmazing("US", a.AmazonAccessName, a.AmazonAccessKey, a.AmazonSecretKey)
	if err != nil {
		return err
	}
	a.client = client
	return nil
}

func (a *Amazon) SearchProducts(query shop.InventoryQuery) ([]shop.InventoryObj, error){
	a.newClient()
}

func (a *Amazon) GetCategories() ([] shop.CategoriesObj, error) {
	a.newClient()
}