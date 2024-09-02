package notifier

import (
	"context"
	"fmt"
	"log/slog"
	"net/smtp"
)

type emailNotifier struct {
	Address string
	Auth    smtp.Auth
	From    string
}

func (email *emailNotifier) Notify(ctx context.Context, request *NotifyRequest) error {
	slog.Info(fmt.Sprintf("Sending email to %s from %s", request.Destination, email.From))
	return nil
}

func NewEmailNotifier() Notifier {
	return &emailNotifier{
		Address: "",
		Auth:    smtp.PlainAuth("", "marcoshuck", "1234", "localhost"),
		From:    "noreply@huck.com.ar",
	}
}
