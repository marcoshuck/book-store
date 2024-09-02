package workers

import (
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log/slog"
)

func RunWorkflowWorker(name string, workflow any, logger *slog.Logger) error {
	logger.Info("Starting workflow worker", slog.String("worker_name", name))
	c, err := client.Dial(client.Options{})
	if err != nil {
		msg := "unable to create client"
		logger.Error(msg, slog.String("worker_name", name), slog.Any("error", err))
		return fmt.Errorf("%s: %w", msg, err)
	}
	defer c.Close()

	w := worker.New(c, name, worker.Options{
		Identity: fmt.Sprintf("%s-worker", name),
	})
	w.RegisterWorkflow(workflow)

	logger.Info("Running worker...", slog.String("worker_name", name))
	if err := w.Run(worker.InterruptCh()); err != nil {
		msg := "unable to start worker"
		logger.Error(msg, slog.String("worker_name", name), slog.Any("error", err))
		return fmt.Errorf("%s: %w", msg, err)
	}

	logger.Info("Activity worker has been stopped", slog.String("worker_name", name))
	return nil
}
