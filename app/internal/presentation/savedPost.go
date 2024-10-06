// Package presentation provides the HTTP handlers for the application.
package presentation

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/business"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
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

	cookie := c.Cookies("session")
	userID, err := ulid.Parse(cookie)
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

// GetAllSavedPostsByUser godoc
//
//	@Summary		Get all saved posts by user
//	@Description	Get all saved posts by user
//	@Tags			saved_posts
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		model.Post	"Saved posts with details"
//	@Failure		400		{string}	string		"Bad request"
//	@Router			/saved_posts [get]
func (ph *SavedPostHandler) GetAllSavedPostsByUser(c *fiber.Ctx) error {
	cookie := c.Cookies("session")
	userID := cookie

	posts, err := ph.SavedPostUsecase.GetAllSavedPostsByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}

// UnsavePost handles the unsaving of a saved post.
// @Summary      Unsave a post
// @Description  Remove a saved post
// @Tags         saved_posts
// @Accept       json
// @Param        body body model.UnsavePostRequest true "Unsave post request body"
// @Produce      json
// @Success      200  {object}  model.SuccessResponse "Saved post removed successfully"
// @Failure      400  {object}  model.ErrorResponse "Bad request"
// @Failure      404  {object}  model.ErrorResponse "Saved post not found"
// @Failure      500  {object}  model.ErrorResponse "Internal server error"
// @Router       /saved_posts [delete]
func (sh *SavedPostHandler) UnsavePost(c *fiber.Ctx) error {
	var request model.UnsavePostRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Error: err.Error()})
	}

	cookie := c.Cookies("session")
	userID := cookie

	// Validate postID
	if request.PostID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Error: "invalid post_id format"})
	}

	// Call the use case to unsave the post
	err := sh.SavedPostUsecase.Unsave(userID, request.PostID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{Error: "saved post not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(model.SuccessResponse{Message: "Saved post removed successfully"})
}
