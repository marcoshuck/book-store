package orders

import (
	"context"
	"github.com/marcoshuck/book-store/domain"
	"github.com/marcoshuck/book-store/workers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
)

type OrderCreator interface {
	CreateOrder(ctx context.Context, request *domain.Order) error
}

type orderCreator struct {
	db *gorm.DB
}

func (svc *orderCreator) CreateOrder(ctx context.Context, request *domain.Order) error {
	if err := svc.db.Model(&domain.Order{}).Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func NewOrderCreator(db *gorm.DB) OrderCreator {
	return &orderCreator{
		db: db,
	}
}

func RunOrderCreatorWorker() error {
	logger := slog.Default()
	db, err := gorm.Open(sqlite.Open("orders.db"))
	if err != nil {
		return err
	}
	svc := NewOrderCreator(db)
	return workers.RunWorker("create-order", svc, logger)
}
