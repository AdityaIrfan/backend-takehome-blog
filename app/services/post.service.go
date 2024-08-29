package services

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type postService struct {
	postRepo    contract.IPostRepository
	commentRepo contract.ICommentRepository
}

func NewPostService(
	postRepo contract.IPostRepository,
	commentRepo contract.ICommentRepository) contract.IPostService {
	return &postService{
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

func (p *postService) Create(c echo.Context, in *models.PostCreateRequest) error {
	post := &models.Post{
		ID:       ulid.Make().String(),
		Title:    in.Title,
		Content:  in.Content,
		AuthorID: c.Get("claims").(jwt.MapClaims)["Id"].(string),
	}

	if err := p.postRepo.Create(post); err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	res := post.ToResponse()
	res.AuthorID = post.Author.ID
	res.AuthorName = post.Author.Name

	return helpers.Response(c, http.StatusCreated, "success create post", res)
}

func (p *postService) Update(c echo.Context, in *models.PostUpdateRequest) error {
	post, err := p.postRepo.GetByCustomAndSelectedFields(map[string]interface{}{
		"id":        in.ID,
		"author_id": in.AuthorID,
	}, "*")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	} else if post == nil {
		return helpers.Response(c, http.StatusNotFound, "post not found")
	}

	var anyUpdatedFields bool
	if in.Title != "" && post.Title != in.Title {
		post.Title = in.Title
		anyUpdatedFields = true
	}

	if in.Content != "" && post.Content != in.Content {
		post.Content = in.Content
		anyUpdatedFields = true
	}

	if anyUpdatedFields {
		if err := p.postRepo.Update(post); err != nil {
			return helpers.ResponseUnprocessableEntity(c)
		}
	}

	return helpers.Response(c, http.StatusOK, "success update post", post.ToResponse())
}

func (p *postService) GetAllMine(c echo.Context, authorId string, cursor *helpers.Cursor) error {
	posts, pagination, err := p.postRepo.GetAllByCursorAndSelectedFieldsPaginate(map[string]interface{}{
		"author_id": authorId,
	}, cursor, "id, title, created_at")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	var listPostResponse = []*models.PostResponse{}
	for _, post := range posts {
		postResponse := post.ToResponseList()

		if totalComments, err := p.commentRepo.GetTotalByCustom(map[string]interface{}{
			"post_id": post.ID,
		}); err == nil {
			postResponse.TotalComments = uint32(totalComments)
		}

		listPostResponse = append(listPostResponse, postResponse)
	}

	return helpers.Response(c, http.StatusOK, "success get list posts", listPostResponse, pagination)
}

func (p *postService) GetAll(c echo.Context, cursor *helpers.Cursor) error {
	posts, pagination, err := p.postRepo.GetAllByCursorAndSelectedFieldsPaginate(nil, cursor, "id, title, created_at")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	var listPostResponse = []*models.PostResponse{}
	for _, post := range posts {
		postResponse := post.ToResponseList()

		if totalComments, err := p.commentRepo.GetTotalByCustom(map[string]interface{}{
			"post_id": post.ID,
		}); err == nil {
			postResponse.TotalComments = uint32(totalComments)
		}

		listPostResponse = append(listPostResponse, postResponse)
	}

	return helpers.Response(c, http.StatusOK, "success get list posts", listPostResponse, pagination)
}

func (p *postService) GetDetail(c echo.Context, postId string) error {
	post, err := p.postRepo.GetByCustomAndSelectedFields(map[string]interface{}{
		"id": postId,
	}, "*", "Author")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	} else if post == nil {
		return helpers.Response(c, http.StatusNotFound, "post not found")
	}

	totalComments, err := p.commentRepo.GetTotalByCustom(map[string]interface{}{
		"post_id": post.ID,
	})
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	postResponse := post.ToResponse()
	postResponse.AuthorID = post.Author.ID
	postResponse.AuthorName = post.Author.Name
	postResponse.TotalComments = uint32(totalComments)

	return helpers.Response(c, http.StatusOK, "success get detail post", postResponse)
}

func (p *postService) Delete(c echo.Context, postId string) error {
	post, err := p.postRepo.GetByCustomAndSelectedFields(map[string]interface{}{
		"id":        postId,
		"author_id": c.Get("claims").(jwt.MapClaims)["Id"].(string),
	}, "id")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	} else if post == nil {
		return helpers.Response(c, http.StatusNotFound, "post not found")
	}

	if err := p.postRepo.Delete(post); err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	return helpers.Response(c, http.StatusOK, "success delete post")
}
