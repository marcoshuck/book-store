package main

import (
	"github.com/marcoshuck/book-store/orders"
	"github.com/marcoshuck/book-store/payments"
	"github.com/marcoshuck/book-store/workflows"
	"os"
)

func main() {
	// Activity workers
	go func() {
		if err := orders.RunOrderCreatorWorker(); err != nil {
			return
		}
	}()

	go func() {
		if err := payments.RunPaymentWorker(); err != nil {
			return
		}
	}()

	go func() {
		if err := orders.RunNotifierWorker(); err != nil {
			return
		}
	}()

	if err := workflows.RunPlacerOrderWorkflow(); err != nil {
		os.Exit(1)
		return
	}
}
