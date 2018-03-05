package server

import(
	"golang.org/x/net/context"

	vending "github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending"
	empty "github.com/golang/protobuf/ptypes/empty"
)

func (s *vendingServer) VendGetVendById(ctx context.Context, msg *vending.VendGetVendByIdRequest)(*vending.Vend, error){
	return &vending.Vend{}, nil
}
func (s *vendingServer) VendGetSessionVend(ctx context.Context, msg *empty.Empty)(*vending.Vend, error){
	return &vending.Vend{}, nil
}
