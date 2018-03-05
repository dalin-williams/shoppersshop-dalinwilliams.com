
package server

import(
	"golang.org/x/net/context"
	empty "github.com/golang/protobuf/ptypes/empty"

	vending "github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending"
)


func (s *vendingServer) SessionCreateUser(ctx context.Context, msg *vending.SessionCreateUserRequest) (*vending.SessionCreateUserResponse, error){
	return &vending.SessionCreateUserResponse{}, nil
}

func (s *vendingServer) SessionCreateSession(ctx context.Context, msg *empty.Empty)(*vending.SessionCreateSessionResponse, error){
	return &vending.SessionCreateSessionResponse{}, nil
}

func (s *vendingServer) SessionUpdateUser(ctx context.Context, msg *vending.SessionUpdateUserRequest) (*empty.Empty, error){
	return &empty.Empty{}, nil
}

func (s *vendingServer) SessionGetUserByUserId(ctx context.Context, msg *vending.SessionGetUserByUserIdRequest) (*vending.User, error){
	return &vending.User{}, nil
}

func (s *vendingServer) SessionDeleteUser(ctx context.Context, msg *vending.SessionDeleteUserRequest) (*empty.Empty, error){
	return &empty.Empty{}, nil
}

func (s *vendingServer) SessionLoginUser(ctx context.Context, msg *vending.SessionLoginUserRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *vendingServer) SessionLogoutUser(ctx context.Context, msg *empty.Empty)(*empty.Empty, error){
	return &empty.Empty{}, nil
}

func (s *vendingServer) SessionGetCurrentSession(ctx context.Context, msg *empty.Empty) (*vending.SessionGetCurrentSessionResponse, error) {
	return &vending.SessionGetCurrentSessionResponse{}, nil
}