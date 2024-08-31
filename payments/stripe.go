package payments

import "context"

type PaymentGateway interface {
	CapturePayment(ctx context.Context, id string) error
}

type stripePaymentGateway struct {
}

func (stripe *stripePaymentGateway) CapturePayment(ctx context.Context, id string) error {
	return nil
}

func NewPaymentGateway() PaymentGateway {
	return &stripePaymentGateway{}
}
