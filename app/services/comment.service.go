package services

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type commentService struct {
	commentRepo contract.ICommentRepository
	postRepo    contract.IPostRepository
}

func NewCommentService(
	commentRepo contract.ICommentRepository,
	postRepo contract.IPostRepository) contract.ICommentService {
	return &commentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func (cm *commentService) Create(c echo.Context, in *models.CommentWriteRequest) error {
	// check post by id
	post, err := cm.postRepo.GetByCustomAndSelectedFields(map[string]interface{}{
		"id": in.PostID,
	}, "id")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	} else if post == nil {
		return helpers.Response(c, http.StatusNotFound, "post not found")
	}

	// check comment parent if exist
	if in.ParentID != "" {
		comment, err := cm.commentRepo.GetByCustomAndSelectedFields(map[string]interface{}{
			"id": in.ParentID,
		}, "id")
		if err != nil {
			return helpers.ResponseUnprocessableEntity(c)
		} else if comment == nil {
			return helpers.Response(c, http.StatusNotFound, "parent comment not found")
		}
	}

	comment := &models.Comment{
		ID:       ulid.Make().String(),
		PostID:   in.PostID,
		AuthorID: in.AuthorId,
		Content:  in.Content,
	}

	if in.ParentID != "" {
		comment.ParentID = &in.ParentID
	}

	if err := cm.commentRepo.Create(comment); err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	return helpers.Response(c, http.StatusOK, "success create comment", comment.ToResponse())
}

func (cm *commentService) GetAllByPost(c echo.Context, postId, parentId string, cursor *helpers.Cursor) error {
	var parentIdValue interface{}

	if parentId != "" {
		if commentParent, err := cm.commentRepo.GetByCustomAndSelectedFields(map[string]interface{}{
			"id": parentId,
		}, "id"); err != nil {
			return helpers.ResponseUnprocessableEntity(c)
		} else if commentParent == nil {
			return helpers.Response(c, http.StatusOK, "comment parent not found")
		}
		parentIdValue = &parentId
	}

	comments, pagination, err := cm.commentRepo.GetAllByCursorAndSelectedFieldsPaginate(map[string]interface{}{
		"post_id":   postId,
		"parent_id": parentIdValue,
	}, cursor, "*", "Author")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	var listCommentResponse = []*models.CommentResponse{}
	for _, comment := range comments {
		commentResponse := comment.ToResponse()

		if totalComments, err := cm.commentRepo.GetTotalByCustom(map[string]interface{}{
			"parent_id": comment.ID,
		}); err == nil {
			commentResponse.TotalComment = uint32(totalComments)
		}

		listCommentResponse = append(listCommentResponse, commentResponse)
	}

	return helpers.Response(c, http.StatusOK, "success get all comments", listCommentResponse, pagination)
}
