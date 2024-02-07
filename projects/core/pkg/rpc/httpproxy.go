package rpc

import (
	"context"
	"core/pkg/helper/config"
	"core/pkg/httpserver"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type HTTPProxy struct {
	server http.Server
}

func (proxy *HTTPProxy) Start(RPCServer *grpc.Server) {
	// proxy http server
	grpcWebServer := grpcweb.WrapServer(
		RPCServer,
		grpcweb.WithOriginFunc(ValidateCorsByOrigin),
	)
	wrapperHandler := httpSplitRequest(
		grpcWebServer,
		httpserver.GetHandler().ServeHTTP,
	)

	// @TODO split servers in order to add http rest handlers
	// ref https://github.com/julienschmidt/httprouter
	// ref https://github.com/johanbrandhorst/grpcweb-example/blob/master/main.go
	proxy.server = http.Server{
		Handler: wrapperHandler,
		Addr:    fmt.Sprintf(":%v", config.Env.Webserver.Port),
	}

	go func() {
		log.Printf("Proxy server listening at %v", proxy.server.Addr)
		proxy.server.ListenAndServe()
	}()
}

func (proxy *HTTPProxy) Stop() {
	// Turn off grpc server and proxy server before finish program
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() { cancel() }()
	if err := proxy.server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP Server proxy graceful shutdown failed: %v", err)
	}
}

// Checks if the origin is allow to make the request
// to the HTTP proxy
// Validates CORS from the whitelist URLs allowed
func ValidateCorsByOrigin(originURL string) bool {
	addresses := config.Env.Cors.AllowAddressesOneLine
	allowed := addresses == config.WILDCARD_CHAR

	if !allowed {
		allowed, _ = regexp.MatchString(originURL, addresses)
	}

	return allowed
}

// redirect requests from gRPC to http webserver
// this fallback http server could have custom routes as regular REST API server
func httpSplitRequest(
	grpcServer *grpcweb.WrappedGrpcServer,
	fallback http.HandlerFunc,
) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		enableProxyCors(&writer)

		if grpcServer.IsAcceptableGrpcCorsRequest(req) || grpcServer.IsGrpcWebRequest(req) {
			grpcServer.ServeHTTP(writer, req)
		} else {
			fallback(writer, req)
		}
	})
}

// split request handler should have CORS enabled
func enableProxyCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", config.WILDCARD_CHAR)
}
