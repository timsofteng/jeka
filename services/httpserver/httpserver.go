package httpserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"telegraminput/lib/logger"
	"time"
)

type HTTPServer struct {
	server *http.Server
}

func New(
	logger logger.Logger,
	host string, port string,
	services StrictServerInterface,
) (*HTTPServer, error) {
	ctx := context.Background()
	mux := http.NewServeMux()

	handlers, err := WrapToOapiHandler(logger, mux, services)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap handlers to oapi handler: %w", err)
	}

	const (
		readTimeout  = time.Second * 4
		writeTimeout = time.Second * 4
	)

	//nolint:exhaustruct
	server := &http.Server{
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
		Addr:         host + ":" + port,
		Handler:      Cors(handlers),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return &HTTPServer{server: server}, nil
}

func (h *HTTPServer) Start() {
	h.server.ListenAndServe()
}

func (h *HTTPServer) Stop(ctx context.Context) {
	h.server.Shutdown(ctx)
}
