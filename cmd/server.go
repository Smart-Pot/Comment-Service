package cmd

import (
	"commentservice/config"
	"commentservice/endpoints"
	"commentservice/service"
	"commentservice/transport"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func startServer() error {
	// Defining Logger
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewService(logger)
	endpoint := endpoints.MakeEndpoints(service)
	handler := transport.MakeHTTPHandlers(endpoint, logger)

	// Set handler and listen given port
	http.Handle("/comment", handler)
	fmt.Println("Server Start Listening to " + config.C.Server.Address)
	return http.ListenAndServe(config.C.Server.Address, handler)
}
