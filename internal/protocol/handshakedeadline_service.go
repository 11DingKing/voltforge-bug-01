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
	return action(context.Background())
}
