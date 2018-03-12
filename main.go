package main

import (
	"encoding/json"
	"flag"
	vending "github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"strings"

	"os"
	"path/filepath"
	rt "runtime"
)

var (
	_, b, _, _ = rt.Caller(0)
	basepath   = filepath.Dir(b)

	apiPath = "/"

	ctxShutdown context.CancelFunc
	//ctxRoot context.Context

	vendEndPoint = flag.String("vending_endpoint", "localhost:9090", "endpoint of VendingMachineService")

	swaggerDir  = flag.String("swagger_dir", basepath+"/vendor/github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending", "path to the directory which contains swagger definitions")
	swaggerFile = string(*swaggerDir + "/vending.swagger.json")
	///github.com/dalin-williams/shoppersshop-protoc-dalinwilliams-com/vending/vending.swagger.json
)

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := vending.RegisterSHOPPERSSHOP_VendingMachineServiceHandlerFromEndpoint(ctx, mux, *vendEndPoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix( /*r.URL.Path*/ swaggerFile, ".swagger.json") {
		glog.Errorf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	glog.Infof("Serving %s" /*r.URL.Path*/)
	//p := strings.TrimPrefix( /*r.URL.Path*/ swaggerFile, "/swagger/")
	//p = path.Join(*swaggerDir, p)
	p := swaggerFile
	http.ServeFile(w, r, p)
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

// Run starts a HTTP server and blocks forever if successful.
func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.Handle(apiPath, gw)

	return http.ListenAndServe(address, allowCORS(mux))
}

func main() {
	var configPath string
	root := &cobra.Command{
		Use: "vendor",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if configPath == "" {
				configPath = "./config/config.json"
			}
			b, err := ioutil.ReadFile(configPath)
			if err != nil {
				return errors.Wrapf(err, `failed to read config file at "%s"`, configPath)
			}

			if err := json.Unmarshal(b, &config); err != nil {
				return errors.Wrapf(err, `failed to parse config file at "%s"`, configPath)
			}

			glog.Infof(`using config found at "%s"`, configPath)
			return nil
		},
	}
	root.AddCommand(cmdMigrate)
	root.AddCommand(cmdServer)
	root.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})
	root.PersistentFlags().StringVarP(&configPath, "config", "c", "./config/config.json", "path to config file")
	root.ParseFlags(os.Args)

	if err := root.Execute(); err != nil {
		shutdown(1, "failed to parse command line: %v", err)
	}
}
