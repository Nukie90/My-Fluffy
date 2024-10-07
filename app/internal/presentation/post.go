package presentation

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/business"
	"github.com/gofiber/fiber/v2"
	"io"
	"strconv"
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
//	@Param			reward	formData	string	true	"Post reward"
//	@Produce		json
//	@Success		200	{string}	string	"Post created successfully"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/posts [post]
func (ph *PostHandler) CreatePost(c *fiber.Ctx) error {
	title := c.FormValue("title")
	fmt.Println(title)
	content := c.FormValue("content")
	fmt.Println(content)
	reward := c.FormValue("reward")
	fmt.Println(reward)
	//reward string to float64
	rewardFloat, err := strconv.ParseFloat(reward, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

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
	post.Reward = rewardFloat

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

// FoundPet godoc
//
//	@Summary		Found pet
//	@Description	Found pet
//	@Tags			posts
//	@Accept			json
//	@Param			id	body	model.FoundPost	true	"Post ID"
//	@Produce		json
//	@Success		200	{string}	string	"Pet found"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/posts/found [put]
func (ph *PostHandler) FoundPet(c *fiber.Ctx) error {
	var foundPost model.FoundPost
	fmt.Println(foundPost.FoundID)
	if err := c.BodyParser(&foundPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	cookie := c.Cookies("session")
	foundPost.FoundID = cookie
	err := ph.PostUsecase.FoundPet(&foundPost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"InternalError": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Pet found"})
}

// GetPaginatedPosts godoc
//
//	@Summary		Get paginated posts
//	@Description	Get paginated posts for the feed, 10 at a time
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int			true	"Page number"
//	@Success		200		{array}		model.Post	"Paginated posts"
//	@Failure		400		{string}	string		"Bad request"
//	@Router			/posts/feed [get]
func (ph *PostHandler) GetPaginatedPosts(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page number",
		})
	}

	posts, err := ph.PostUsecase.GetPaginatedPosts(page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch posts",
		})
	}

	return c.JSON(posts)
}

// Confirmation godoc
//
//	@Summary		Confirmation
//	@Description	Confirmation
//	@Tags			posts
//	@Accept			json
//	@Param			id	body	model.ConfirmationPost	true	"Post ID"
//	@Produce		json
//	@Success		200	{string}	string	"Confirmation sent"
//	@Failure		400	{string}	string	"Bad request"
//	@Router			/posts/confirmation/ [put]
func (ph *PostHandler) Confirmation(c *fiber.Ctx) error {
	var confirmation model.ConfirmationPost
	fmt.Println(confirmation.ID)
	if err := c.BodyParser(&confirmation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error1": err.Error()})
	}

	err := ph.PostUsecase.Confirmation(confirmation.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error2": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Confirmation sent"})
}
