package server

import (
	"flag"
	"net"

	"google.golang.org/grpc"
	"github.com/golang/glog"
	"golang.org/x/net/context"

	vending "github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending"
)

// Starts the vending GRPC server on port 9090
func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()


	//TODO: Refactor into config
	l, err := net.Listen("tcp", "9090")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	vending.RegisterSHOPPERSSHOP_VendingMachineServiceServer(s, newVendingServer())

	return s.Serve(l)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}


