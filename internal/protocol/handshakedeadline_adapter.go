package protocol

import "context"

func RunHandshakeDeadline(ctx context.Context) error {
	svc := HandshakeDeadlineService{}
	return svc.Execute(ctx, func(callCtx context.Context) error {
		return callCtx.Err()
	})
}
