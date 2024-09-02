package notifier

import (
	"context"
)

type NotifyRequest struct {
	Destination string
	Subject     string
	Body        string
}

type Notifier interface {
	Notify(ctx context.Context, request *NotifyRequest) error
}
