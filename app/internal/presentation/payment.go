package presentation

import (
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/business"
	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	PaymentUsecase *business.PaymentUsecase
}

func NewPaymentHandler(pu *business.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{PaymentUsecase: pu}
}

// CreateUserPayment godoc
//
//	@Summary		Create a new payment
//	@Description	Create a new payment
//	@Tags			payments
//	@Accept			json
//	@Param			payment	body	model.CreatePayment	true	"Payment information"
//	@Produce		json
//	@Success		200	{string}	string	"Payment created successfully"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/payments [post]
func (ph *PaymentHandler) CreateUserPayment(c *fiber.Ctx) error {
	payment := model.CreatePayment{}
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := ph.PaymentUsecase.CreateUserPayment(&payment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Payment created successfully"})
}

// GetPaymentsFromSpecificUser godoc
//
//	@Summary		Get payments from specific user
//	@Description	Get payments from specific user
//	@Tags			payments
//	@Accept			json
//	@Param			userID	path	string	true	"User ID"
//	@Produce		json
//	@Success		200	{array}		model.Payment	"Payments from specific user"
//	@Failure		400	{string}	string			"Bad request"
//	@Router			/payments/user/{userID} [get]
func (ph *PaymentHandler) GetPaymentsFromSpecificUser(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userID is required"})
	}

	payments, err := ph.PaymentUsecase.GetPaymentsFromSpecificUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(payments)
}
