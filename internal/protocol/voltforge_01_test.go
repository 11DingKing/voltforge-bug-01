package protocol

import (
	"context"
	"errors"
	"testing"
)

func TestVoltForge01(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := RunHandshakeDeadline(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
}
