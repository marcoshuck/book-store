package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	BookID   uuid.UUID `json:"book_id"`
	OrderID  int       `json:"order_id"`
	Quantity int       `json:"quantity"`
	Price    Price     `json:"price" gorm:"embedded"`
}

type Order struct {
	gorm.Model
	CustomerID uuid.UUID   `json:"customer_id"`
	OrderItems []OrderItem `json:"order_items"`

	ShippingAddress   Address `json:"shipping_address"`
	ShippingAddressID int     `json:"-"`

	BillingAddress   Address `json:"billing_address"`
	BillingAddressID int     `json:"-"`
}
