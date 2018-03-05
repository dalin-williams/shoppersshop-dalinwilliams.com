package server

import(
	"fmt"
	"io"
	"sync"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"github.com/rogpeppe/fastuuid"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	empty "github.com/golang/protobuf/ptypes/empty"

	vending "github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending"
)

func (s *vendingServer) InventorySearchForItems(ctx context.Context, msg *vending.InventorySearchForItemsRequest)(*vending.InventorySearchForItemsResponse, error){
	return &vending.InventorySearchForItemsResponse{}, nil
}

func (s *vendingServer) InventorySearchAllItemCategories(ctx context.Context, msg *empty.Empty)(*vending.InventorySearchAllItemCategoriesResponse, error){
	return &vending.InventorySearchAllItemCategoriesResponse{}, nil
}
