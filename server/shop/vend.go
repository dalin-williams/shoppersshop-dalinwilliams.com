package shop

import (
	"github.com/satori/go.uuid"
)

type Status int
const (
	Placed Status = iota
	Approved
	Delivered
	Cancelled
)

type VendObj struct {
	Id 		uuid.UUID				`json:"id"`
	Total   float32					`json:"total"`
	Items   []struct {
		Item 	 	InventoryObj	`json:"item"`
		Quantity	int				`json:"quantity"`
	} `json:"items"`
	Status	Status					`json:"status"`
}


type Vend interface {
	// Create a vend object from an existing  session
	CreateVend()(vendResults []VendObj, err error)

	// Result from the webhook which returns the vend data
	FetchVend()(vendObj VendObj, err error)
}
