package server

import (
	"commentservice/endpoints"
	"commentservice/service"
	"errors"
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

func StartServer(options ServerOptions) error {

	switch options.Mode {
	case "http":
		startHTTPServer(options, false)
		break
	case "grpc":
		startGRPCServer(options, false)
		break
	case "both":
		startHTTPServer(options, false)
		startGRPCServer(options, true)
	default:
		return errors.New(fmt.Sprintf("Invalid server mode: %s", options.Mode))
	}

	return nil
}

func prepareServer() (endpoints.Endpoints, log.Logger) {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewService(logger)
	endpoint := endpoints.MakeEndpoints(service)

	return endpoint, logger
}
