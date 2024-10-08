package presentation

import (
	"fmt"

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
//	@Accept			multipart/form-data
//	@Param			amount		formData	string	true	"Payment amount"
//	@Param			receiver_id	formData	string	true	"Receiver ID"
//	@Produce		json
//	@Success		200	{string}	string	"Payment created successfully"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/payments [post]
func (ph *PaymentHandler) CreateUserPayment(c *fiber.Ctx) error {
	var payment model.CreatePayment

	// Parse the JSON body directly into the payment struct
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Ensure the amount is set and convert it to float
	if payment.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}

	// Get the session cookie
	cookie := c.Cookies("session")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Session cookie not found"})
	}

	payment.UserID = cookie

	// Call the payment use case to create the payment
	err := ph.PaymentUsecase.CreateUserPayment(&payment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Payment created successfully"})
}

// GetPaymentsFromSpecificUser godoc
//
//	@Summary		Get all payments from a specific user
//	@Description	Get all payments from a specific user
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Payments retrieved successfully"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/payments [get]
func (ph *PaymentHandler) GetPaymentsFromSpecificUser(c *fiber.Ctx) error {
	cookie := c.Cookies("session")
	userID := cookie
	fmt.Println("User ID: " + userID)
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userID is required"})
	}

	payments, err := ph.PaymentUsecase.GetPaymentsFromSpecificUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(payments)
}
