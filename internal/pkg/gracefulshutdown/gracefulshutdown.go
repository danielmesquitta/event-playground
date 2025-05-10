package gracefulshutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// WithShutdownSignal waits for a CTRL+C signal and then executes the provided callback functions.
// It wraps the provided context and returns a new context that will be canceled when the shutdown signal is received.
func WithShutdownSignal(
	ctx context.Context,
	callbacks ...func(),
) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	// Create a channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to handle the shutdown
	go func() {
		// Wait for the signal
		sig := <-sigChan
		log.Printf("Received signal: %v", sig)

		// Execute the shutdown function
		for _, callback := range callbacks {
			callback()
		}

		// Cancel the context
		cancel()
	}()

	return ctx
}
