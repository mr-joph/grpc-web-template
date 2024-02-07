package rpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"core/pkg/helper/config"
	"core/pkg/helper/errorhandler"
	"core/pkg/rpc/server"

	"google.golang.org/grpc/reflection"
)

var RPCServer = server.GRPCServer

func StartServer() {
	proxy := HTTPProxy{}
	port := strconv.Itoa(config.Env.Server.Port)
	// @TODO add a TLS server for more security
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	errorhandler.Check(err)

	// @TODO disable reflection for gRPC in production
	reflection.Register(RPCServer)
	proxy.Start(RPCServer)

	// subscribe and register all services for this gRPC Server
	loadServices()

	// Now let's go and start the server!! 🚀
	go func() {
		log.Printf("server listening at %v", listen.Addr())
		err = RPCServer.Serve(listen)
		errorhandler.Check(err)
	}()

	// wait for CTRL+C in command line
	// to stop process and close connections
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Blociking until receive CTRL+C Signal
	<-ch

	proxy.Stop()
	RPCServer.Stop()
	fmt.Println("\nBye! 👋")
}
