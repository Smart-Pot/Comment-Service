package cmd

import (
	"commentservice/server"
	"flag"
)

var serverMode = flag.String("mode", "http", "")
var port = flag.Int("port", 8080, "")
var secondaryPort = flag.Int("secport", 8081, "")

func Execute() error {
	flag.Parse()

	// Define server options
	opt := server.ServerOptions{
		Port:          *port,
		Mode:          *serverMode,
		SecondaryPort: *secondaryPort,
	}

	// Start server
	return server.StartServer(opt)

}

/**


func main() {
	config.ReadConfig()

	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewService(logger)
	endpoint := endpoints.MakeEndpoints(service)
	grpcServer := transport.NewGRPCServer(endpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", config.C.Server.Address)
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



*/
