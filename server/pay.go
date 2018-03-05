package server

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"time"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	empty "github.com/golang/protobuf/ptypes/empty"

	vending "github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending"
)

func (s *vendingServer) PayViewPaymentInfo(ctx context.Context, in *empty.Empty) (*vending.PayViewPaymentInfoResponse, error){
	return &vending.PayViewPaymentInfoResponse{}, nil
}

// Representing the endpoint PUT /pay/purchase/{id}
func (s *vendingServer) PayAddToCart(ctx context.Context, msg *vending.PayAddToCartRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

// Representing the endpoint GET /pay/purchase
func (s *vendingServer) PayViewCart(ctx context.Context, msg *empty.Empty)(*vending.PayViewCartResponse, error){
	return &vending.PayViewCartResponse{}, nil
}

// Representing the endpoint POST /pay/purchase
func (s *vendingServer) PayPurchaseCart(ctx context.Context, msg *vending.PayPurchaseCartRequest) (*vending.PayPurchaseCartResponse, error){
	return &vending.PayPurchaseCartResponse{}, nil
}

// Representing the endpoint DELETE /pay/purchase/{orderId}
func (s *vendingServer) PayDeleteOrder(ctx context.Context, msg *vending.PayDeleteOrderRequest)(*empty.Empty, error){
	return &empty.Empty{}, nil
}

// Representing the endpoint PUT /pay/purchase/{orderId}
func (s *vendingServer) PayUpdateOrder(ctx context.Context, msg *vending.PayUpdateOrderRequest)(*vending.PayUpdateOrderResponse, error){
	return &vending.PayUpdateOrderResponse{}, nil
}

// Representing the endpoint POST /pay/purchase/{orderId}
func (s *vendingServer) PayReOrderItems(ctx context.Context, msg *vending.PayReOrderItemsRequest)(*empty.Empty, error){
	return &empty.Empty{}, nil
}