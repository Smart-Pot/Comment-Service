package server

import (
	"commentservice/pb"
	"commentservice/transport"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func startGRPCServer(options ServerOptions, isSecondary bool) {
	var port int
	if isSecondary {
		port = options.SecondaryPort
	} else {
		port = options.Port
	}
	endpoint, logger := prepareServer()
	grpcServer := transport.NewGRPCServer(endpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterCommentServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)

}
