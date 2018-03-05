package server

import (
	"github.com/rogpeppe/fastuuid"

	vending "github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending"
)

var uuidGen = fastuuid.Generator{}

type vendingServer struct{
	//JIC we want to enable streaming
	//v map[string]*vending.SHOPPERSSHOP_VendingMachineServiceServer
	//m sync.Mutex

}

type VendingServer interface {
	vending.SHOPPERSSHOP_VendingMachineServiceServer
}



func newVendingServer() VendingServer {
	return new(vendingServer)
}