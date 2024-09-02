package payments

import (
	"context"
	"log/slog"
)

type stripePaymentGateway struct {
	logger *slog.Logger
}

func (stripe *stripePaymentGateway) CapturePayment(ctx context.Context, id string) error {
	stripe.logger.InfoContext(ctx, "Capturing stripe payment", slog.String("payment_intent_id", id))
	return nil
}

func NewStripePaymentGateway(logger *slog.Logger) PaymentGateway {
	return &stripePaymentGateway{
		logger: logger,
	}
}
