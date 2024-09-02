package payments

import (
	"context"
	"github.com/marcoshuck/book-store/workers"
	"log/slog"
)

type PaymentGateway interface {
	CapturePayment(ctx context.Context, id string) error
}

func RunPaymentWorker() error {
	logger := slog.Default()
	svc := NewStripePaymentGateway(logger)
	return workers.RunActivityWorker("capture-payment", svc, logger)
}
