package services

import (

	"github.com/funkeyfreak/vending-machine-api/server/shop"
	"github.com/funkeyfreak/vending-machine-api/server/shop/services/amazon"
)

type driver struct {
	Amazon         amazon.Amazon `json:"amazon"`

	services	   map[string]shop.Service

}

func (d *driver) GenerateServices() error {
	d.services = make(map[string]shop.Service)
	d.services["amazon"] = &d.Amazon

	return nil
}

//Services has nothing to do with the pay platform.... yet
func (d *driver) Pay(s *shop.Shop) error {
	return nil
}

func (d *driver) Session(s *shop.Shop) error {
	return nil
}

//TODO: fix this
func (d *driver) Inventory(s *shop.Shop) error {

	s.Inventory = inventory{d.services}
	return nil
}

func (d *driver) Vend(s *shop.Shop) error {
	return nil
}


func (d *driver) Category(s *shop.Shop) error {
	return nil
}

func init() {
	shop.RegisterDriver("services", new(driver))
}