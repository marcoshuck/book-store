package workers

import (
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log/slog"
)

func RunWorker(name string, activity any, logger *slog.Logger) error {
	logger.Info("Starting worker", slog.String("worker_name", name))
	c, err := client.Dial(client.Options{})
	if err != nil {
		msg := "unable to create client"
		logger.Error(msg, slog.String("worker_name", name), slog.Any("error", err))
		return fmt.Errorf("%s: %w", msg, err)
	}
	defer c.Close()

	w := worker.New(c, name, worker.Options{})
	w.RegisterActivity(activity)

	logger.Info("Running worker...", slog.String("worker_name", name))
	if err := w.Run(worker.InterruptCh()); err != nil {
		msg := "unable to start worker"
		logger.Error(msg, slog.String("worker_name", name), slog.Any("error", err))
		return fmt.Errorf("%s: %w", msg, err)
	}

	logger.Info("Worker has been stopped", slog.String("worker_name", name))
	return nil
}
