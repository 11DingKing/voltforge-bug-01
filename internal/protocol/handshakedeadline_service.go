package protocol

import (
	"context"
	"fmt"
)

type HandshakeDeadlineService struct{}

func (s HandshakeDeadlineService) Execute(ctx context.Context, action func(context.Context) error) error {
	if ctx == nil {
		return fmt.Errorf("handshakedeadline: nil context")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("handshakedeadline before action: %w", err)
	}
	if err := action(ctx); err != nil {
		return fmt.Errorf("handshakedeadline action: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("handshakedeadline after action: %w", err)
	}
	return nil
}
