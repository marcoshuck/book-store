package orders

import (
	"context"
	"github.com/marcoshuck/book-store/domain"
	"github.com/marcoshuck/book-store/notifier"
	"github.com/marcoshuck/book-store/workers"
	"log/slog"
)

type OrderNotifier interface {
	NotifyOrderPlaced(ctx context.Context, request *domain.Customer) error
}

type orderNotifier struct {
	notifier notifier.Notifier
}

func (svc *orderNotifier) NotifyOrderPlaced(ctx context.Context, email string) error {
	return svc.notifier.Notify(ctx, &notifier.NotifyRequest{
		Destination: email,
		Subject:     "Your order has been placed",
		Body:        "Congratulations, your order has been correctly placed. Will send you an email soon.",
	})
}

func RunNotifierWorker() error {
	logger := slog.Default()
	svc := &orderNotifier{
		notifier: notifier.NewEmailNotifier(),
	}
	return workers.RunActivityWorker("notify-order", svc, logger)
}
