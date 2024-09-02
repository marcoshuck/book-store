package workflows

import (
	"github.com/marcoshuck/book-store/domain"
	"github.com/marcoshuck/book-store/workers"
	"go.temporal.io/sdk/workflow"
	"log/slog"
	"time"
)

func PlaceOrderWorkflow(ctx workflow.Context, request *PlaceOrderRequest) error {
	opts := workflow.ActivityOptions{
		TaskQueue:           "create-order",
		StartToCloseTimeout: 10 * time.Minute,
	}
	err := workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, opts), "CreateOrder", request.Order).Get(ctx, nil)
	if err != nil {
		return err
	}

	opts = workflow.ActivityOptions{
		TaskQueue:           "capture-payment",
		StartToCloseTimeout: 10 * time.Minute,
	}
	err = workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, opts), "CapturePayment", request.PaymentID).Get(ctx, nil)
	if err != nil {
		return err
	}

	opts = workflow.ActivityOptions{
		TaskQueue:           "notify-order",
		StartToCloseTimeout: 10 * time.Minute,
	}
	err = workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, opts), "NotifyOrderPlaced", request.Email).Get(ctx, nil)
	return nil
}

type PlaceOrderRequest struct {
	Order     *domain.Order
	PaymentID string
	Email     string
}

func RunPlacerOrderWorkflow() error {
	logger := slog.Default()
	return workers.RunWorkflowWorker("place-order", PlaceOrderWorkflow, logger)
}
