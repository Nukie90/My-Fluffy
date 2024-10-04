package presentation

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/business"
	"github.com/gofiber/fiber/v2"
	"io"
)

type PostHandler struct {
	PostUsecase *business.PostUsecase
}

func NewPostHandler(pu *business.PostUsecase) *PostHandler {
	return &PostHandler{PostUsecase: pu}
}

// CreatePost godoc
//
//	@Summary		Create a new post
//	@Description	Create a new post
//	@Tags			posts
//	@Accept			multipart/form-data
//	@Param			title	formData	string	true	"Post title"
//	@Param			content	formData	string	true	"Post content"
//	@Param			file	formData	file	true	"Post picture"
//	@Produce		json
//	@Success		200	{string}	string	"Post created successfully"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/posts [post]
func (ph *PostHandler) CreatePost(c *fiber.Ctx) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	defer fileContent.Close()

	pictureData, err := io.ReadAll(fileContent)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var post model.CreatePost

	post.Title = title
	post.Content = content
	post.Picture = pictureData
	cookie := c.Cookies("session")
	post.OwnerID = cookie

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = ph.PostUsecase.Create(&post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Post created successfully"})
}

// GetPostsFromSpecificUser godoc
//
//	@Summary		Get posts from specific user
//	@Description	Get posts from specific user
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	[]model.Post	"Posts from specific user"
//	@Failure		400	{string}	string			"Bad request"
//	@Router			/posts/user [get]
func (ph *PostHandler) GetPostsFromSpecificUser(c *fiber.Ctx) error {
	cookie := c.Cookies("session")
	userID := cookie

	fmt.Println(userID)
	posts, err := ph.PostUsecase.GetPostsFromSpecificUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}
