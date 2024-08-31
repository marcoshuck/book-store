package workflows

import (
	"github.com/marcoshuck/book-store/domain"
	"go.temporal.io/sdk/workflow"
	"time"
)

func PlaceOrderWorkflow(ctx workflow.Context, request *PlaceOrderRequest) error {
	opts := workflow.ActivityOptions{
		TaskQueue:           "place-order",
		StartToCloseTimeout: 10 * time.Minute,
	}
	err := workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, opts), "CreateOrder", request.Order).Get(ctx, nil)
	if err != nil {
		return err
	}
	err = workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, opts), "CapturePayment", request.PaymentID).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

type PlaceOrderRequest struct {
	Order     *domain.Order
	PaymentID string
}
