package presentation

import (
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/gofiber/fiber/v2"
	"github.com/Nukie90/my-fluffy/app/internal/business"
)

type PaymentHandler struct {
	PaymentUsecase *business.PaymentUsecase
}

func NewPaymentHandler(pu *business.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{PaymentUsecase: pu}
}

func (ph *PaymentHandler) CreateUserPayment(c *fiber.Ctx) error {
	var payment model.CreatePayment
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := ph.PaymentUsecase.CreateUserPayment(&payment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Payment created successfully"})
}

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

