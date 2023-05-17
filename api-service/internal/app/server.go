package app

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func serverRun(port string, handler http.Handler) error {
	server := Server{httpServer: &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	},
	}

	return server.httpServer.ListenAndServe()
}
