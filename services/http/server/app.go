package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func ListenAndServe(servers ...App) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go handleInterrupt(ctx, servers...)

	errC := make(chan error, len(servers))
	for _, s := range servers {
		s := s
		go func() {
			errC <- s.ListenAndServe()
		}()
	}
	// Stop on first exited server.
	err := <-errC
	if err == http.ErrServerClosed {
		err = nil // Do not report server closed.
	}
	return err
}

func handleInterrupt(ctx context.Context, servers ...App) {
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		return // OK, exit gracefully.
	case sig := <-sigC:
		fmt.Printf("Captured %v signal\n", sig)
	}
	for _, s := range servers {
		_ = s.Shutdown(ctx)
	}
}
