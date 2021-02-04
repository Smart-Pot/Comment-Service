package server

import (
	"commentservice/transport"
	"fmt"
	"net/http"
)

func startHTTPServer(options ServerOptions, isSecondary bool) {
	var port int
	if isSecondary {
		port = options.SecondaryPort
	} else {
		port = options.Port
	}

	endpoint, logger := prepareServer()
	handler := transport.MakeHTTPHandlers(endpoint, logger)

	http.Handle("/comment", handler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		// Log Error
		fmt.Println(err)
	}

}
