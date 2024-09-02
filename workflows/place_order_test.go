package workflows

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/marcoshuck/book-store/domain"
	"github.com/marcoshuck/book-store/orders"
	"github.com/marcoshuck/book-store/payments"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
	"log/slog"
	"testing"
)

type PlacerOrderTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env            *testsuite.TestWorkflowEnvironment
	orderCreator   orders.OrderCreator
	paymentGateway payments.PaymentGateway
}

func TestPlacerOrder(t *testing.T) {
	suite.Run(t, new(PlacerOrderTestSuite))
}

func (s *PlacerOrderTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
	s.orderCreator = orders.NewOrderCreator(nil)
	s.paymentGateway = payments.NewStripePaymentGateway(slog.Default())
}

func (s *PlacerOrderTestSuite) AfterTest(_, _ string) {
	s.env.AssertExpectations(s.T())
}

func (s *PlacerOrderTestSuite) TestCreateOrder_Fails() {
	s.env.OnActivity(s.orderCreator.CreateOrder, mock.Anything, mock.Anything).Return(func(ctx context.Context, order *domain.Order) error {
		s.Assert().NotNil(order)
		return errors.New("error")
	})
	s.env.ExecuteWorkflow(PlaceOrderWorkflow, &PlaceOrderRequest{
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

	s.Require().True(s.env.IsWorkflowCompleted())
	err := s.env.GetWorkflowError()
	s.Assert().Error(err)
}

func (s *PlacerOrderTestSuite) TestCapturePayment_Fails() {
	s.env.OnActivity(s.orderCreator.CreateOrder, mock.Anything, mock.Anything).Return(func(ctx context.Context, order *domain.Order) error {
		s.Assert().NotNil(order)
		return nil
	})
	s.env.OnActivity(s.paymentGateway.CapturePayment, mock.Anything, mock.Anything).Return(func(ctx context.Context, id string) error {
		s.Assert().NotEmpty(id)
		return errors.New("failed to capture payment")
	})
	s.env.ExecuteWorkflow(PlaceOrderWorkflow, &PlaceOrderRequest{
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

	s.Require().True(s.env.IsWorkflowCompleted())
	err := s.env.GetWorkflowError()
	s.Assert().Error(err)
}

func (s *PlacerOrderTestSuite) TestPlaceOrder_Succeeds() {
	s.env.OnActivity(s.orderCreator.CreateOrder, mock.Anything, mock.Anything).Return(func(ctx context.Context, order *domain.Order) error {
		s.Assert().NotNil(order)
		return nil
	})
	s.env.OnActivity(s.paymentGateway.CapturePayment, mock.Anything, mock.Anything).Return(func(ctx context.Context, id string) error {
		s.Assert().NotEmpty(id)
		return nil
	})
	s.env.ExecuteWorkflow(PlaceOrderWorkflow, &PlaceOrderRequest{
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

	s.Require().True(s.env.IsWorkflowCompleted())
	err := s.env.GetWorkflowError()
	s.Assert().NoError(err)
}
