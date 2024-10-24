package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/timsofteng/jeka/lib/logger"
)

type HTTPServer struct {
	server *http.Server
}

func New(
	ctx context.Context,
	logger logger.Logger,
	host string, port string,
	services StrictServerInterface,
) (*HTTPServer, error) {
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

func (h *HTTPServer) Start() error {
	err := h.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to listen and serve http: %w", err)
	}

	return nil
}

func (h *HTTPServer) Stop(ctx context.Context) error {
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)

	defer cancel()

	err := h.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}

	return nil
}
