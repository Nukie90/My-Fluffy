package business

import (
	"fmt"

	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
	"github.com/Nukie90/my-fluffy/app/internal/shared"
	"github.com/oklog/ulid/v2"
)

type PostUsecase struct {
	PostRepo       *repository.PostRepo
	AdminNotifier  *shared.UserNotifier
	ClientNotifier *shared.UserNotifier
}

func NewPostUsecase(pr *repository.PostRepo, an *shared.UserNotifier, cn *shared.UserNotifier) *PostUsecase {
	return &PostUsecase{
		PostRepo:       pr,
		AdminNotifier:  an,
		ClientNotifier: cn,
	}
}

func (pu *PostUsecase) Create(post *model.CreatePost) error {
	p := &entity.Post{
		Title:   post.Title,
		Content: post.Content,
		Picture: post.Picture,
		Reward:  post.Reward,
		OwnerID: ulid.MustParse(post.OwnerID),
	}

	err := pu.PostRepo.Create(p)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PostUsecase) GetPostsFromSpecificUser(userID string) ([]model.Post, error) {
	//to ulid
	useridUlid, err := ulid.Parse(userID)
	if err != nil {
		return nil, err
	}
	posts, err := pu.PostRepo.GetPostFromSpecificUser(useridUlid)
	if err != nil {
		return nil, err
	}

	var postsModel []model.Post
	for _, post := range posts {
		postsModel = append(postsModel, model.Post{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			Status:  post.Status,
			Picture: post.Picture,
			Reward:  post.Reward,
			OwnerID: post.OwnerID.String(),
			FoundID: post.FoundID.String(),
		})
	}

	return postsModel, nil
}

func (pu *PostUsecase) FoundPet(foundPost *model.FoundPost) error {
	//to ulid
	fmt.Println(foundPost.FoundID)
	foundidUlid, err := ulid.Parse(foundPost.FoundID)
	if err != nil {
		fmt.Println("unmarshalling error")
		return err
	}

	err = pu.PostRepo.FoundPet(foundPost.ID, foundidUlid)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//get the post
	foundPostEntity, err := pu.PostRepo.GetPostByID(foundPost.ID)
	if err != nil {
		return err
	}

	//notify the owner
	pu.ClientNotifier.NotifyObserver(foundPostEntity.OwnerID.String(), foundPost.FoundID, "foundPet")

	return nil
}

func (pu *PostUsecase) Confirmation(postID uint) error {
	err := pu.PostRepo.Confirmation(postID)
	if err != nil {
		return err
	}

	//get the post
	post, err := pu.PostRepo.GetPostByID(postID)
	if err != nil {
		return err
	}

	//notify the owner
	pu.ClientNotifier.NotifyObserver(post.FoundID.String(), post.OwnerID.String(), "Confirmation")

	return nil
}

func (pu *PostUsecase) GetPaginatedPosts(page int) ([]model.PaginatedPostResponse, error) {
	const postsPerPage = 10
	offset := (page - 1) * postsPerPage

	posts, err := pu.PostRepo.GetPaginatedPosts(postsPerPage, offset)
	if err != nil {
		return nil, err
	}

	var response []model.PaginatedPostResponse
	for _, post := range posts {
		response = append(response, model.PaginatedPostResponse{
			ID:       post.ID,
			Username: post.Username,
			Title:    post.Title,
			Content:  post.Content,
			Status:   post.Status,
			Picture:  post.Picture,
			OwnerID:  post.OwnerID.String(), // Convert ULID to string
			FoundID:  post.FoundID.String(), // Convert ULID to string
			Reward:   post.Reward,
		})
	}

	return response, nil
}
