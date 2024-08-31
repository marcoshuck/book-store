package orders

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcoshuck/book-store/domain"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestOrderCreator(t *testing.T) {
	suite.Run(t, new(OrderCreatorTestSuite))
}

type OrderCreatorTestSuite struct {
	suite.Suite
	db           *gorm.DB
	orderCreator OrderCreator
	filename     string
}

func (suite *OrderCreatorTestSuite) SetupTest() {
	var err error
	suite.filename = fmt.Sprintf("%s/order_creator.db", suite.T().TempDir())
	suite.Require().NoError(err)

	suite.db, err = gorm.Open(sqlite.Open(suite.filename))
	suite.Require().NoError(err)

	suite.Require().NoError(suite.db.Migrator().AutoMigrate(
		&domain.Address{},
		&domain.OrderItem{},
		&domain.Order{},
	))

	suite.orderCreator = NewOrderCreator(suite.db)
}

func (suite *OrderCreatorTestSuite) TearDownTest() {
	suite.Require().NoError(os.Remove(suite.filename))
}

func (suite *OrderCreatorTestSuite) TestOrderCreator_CreateOrder() {
	var before int64
	suite.Require().NoError(suite.db.Model(&domain.Order{}).Count(&before).Error)

	err := suite.orderCreator.CreateOrder(context.Background(), &domain.Order{
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
	})
	suite.Assert().NoError(err)

	var after int64
	suite.Require().NoError(suite.db.Model(&domain.Order{}).Count(&after).Error)

	suite.Assert().NotEqual(before, after)
	suite.Assert().Equal(before+1, after)
}
