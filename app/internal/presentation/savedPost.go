// Package presentation provides the HTTP handlers for the application.
package presentation

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/business"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

// SavedPostHandler is the handler for saved posts.
type SavedPostHandler struct {
	SavedPostUsecase *business.SavedPostUsecase
}

// NewSavedPostHandler creates a new SavedPostHandler.
func NewSavedPostHandler(spu *business.SavedPostUsecase) *SavedPostHandler {
	return &SavedPostHandler{SavedPostUsecase: spu}
}

// CreateSavedPost handles the creation of a saved post.
// @Summary      Create a new saved post
// @Description  Create a new saved post
// @Tags         saved_posts
// @Accept       json
// @Param        saved_post body model.SavedPost true "Saved post information"
// @Produce      json
// @Success      200  {object}  model.SuccessResponse "Saved post created successfully"
// @Failure      400  {object}  model.ErrorResponse "Bad request"
// @Failure      500  {object}  model.ErrorResponse "Internal server error"
// @Router       /savedPost/saved_posts [post]
func (sh *SavedPostHandler) CreateSavedPost(c *fiber.Ctx) error {
	var savedPost model.SavedPost
	if err := c.BodyParser(&savedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Error: err.Error()})
	}

	// Ensure PostID is valid
	if savedPost.PostID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Error: "invalid post_id format"})
	}

	// Convert UserID from string to ulid.ULID
	userID, err := ulid.Parse(savedPost.UserID) // Assuming UserID in request is a string
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Error: "invalid user_id format"})
	}

	savedPostEntity := entity.SavedPost{
		UserID: userID,
		PostID: savedPost.PostID, // Directly use PostID as it's a uint
	}

	// Call use case to create the saved post
	err = sh.SavedPostUsecase.Create(&savedPostEntity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(model.SuccessResponse{Message: "Saved post created successfully"})
}

// GetAllSavedPostsByUser retrieves all saved posts for a specific user.
// @Summary      Retrieve all saved posts for a user
// @Description  Get all saved posts for a specific user by user ID
// @Tags         saved_posts
// @Accept       json
// @Param        user_id path string true "User ID"
// @Produce      json
// @Success      200  {array}  model.SavedPost "List of saved posts"
// @Failure      404  {object}  model.ErrorResponse "Not found"
// @Failure      500  {object}  model.ErrorResponse "Internal server error"
// @Router       /savedPost/saved_posts/{user_id} [get]
func (sh *SavedPostHandler) GetAllSavedPostsByUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	// Call usecase to get saved posts
	savedPosts, err := sh.SavedPostUsecase.GetAllByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(savedPosts)
}
