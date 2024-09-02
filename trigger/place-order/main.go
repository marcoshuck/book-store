package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/marcoshuck/book-store/domain"
	"github.com/marcoshuck/book-store/workflows"
	"go.temporal.io/sdk/client"
	"log/slog"
	"os"
)

func main() {
	logger := slog.Default()

	logger.Info("Starting workflow: PlaceOrder.")
	c, err := client.Dial(client.Options{})
	if err != nil {
		logger.Error("Unable to create client.", slog.Any("error", err))
		os.Exit(1)
	}
	defer c.Close()

	opts := client.StartWorkflowOptions{
		TaskQueue: "place-order",
	}
	wr, err := c.ExecuteWorkflow(context.Background(), opts, workflows.PlaceOrderWorkflow, &workflows.PlaceOrderRequest{
		Email:     "todo@huck.com.ar",
		PaymentID: "pi_3MrPBM2eZvKYlo2C1TEMacFD",
		Order: &domain.Order{
			CustomerID: uuid.New(),
			OrderItems: []domain.OrderItem{
				{
					BookID:   uuid.New(),
					Quantity: 10,
				},
			},
			ShippingAddress: domain.Address{
				Street:     "Fake Street",
				Number:     123,
				City:       "San Francisco",
				State:      "CA",
				PostalCode: "90210",
				Country:    "USA",
			},
			BillingAddress: domain.Address{
				Street:     "Fake Street",
				Number:     123,
				City:       "San Francisco",
				State:      "CA",
				PostalCode: "90210",
				Country:    "USA",
			},
		},
	})
	if err != nil {
		logger.Error("Unable to execute workflow.", slog.Any("error", err))
		os.Exit(2)
	}

	err = wr.Get(context.Background(), nil)
	if err != nil {
		logger.Error("Unable to execute workflow.", slog.Any("error", err))
		os.Exit(3)
	}

	logger.Info("Workflow has run successfully.", slog.Any("workflow", wr.GetID()))
}
