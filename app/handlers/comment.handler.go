package handlers

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	commentService contract.ICommentService
}

func NewCommentHandler(commentService contract.ICommentService) contract.ICommentHandler {
	return &commentHandler{
		commentService: commentService,
	}
}

func (cm *commentHandler) Create(c echo.Context) error {
	payload := new(models.CommentWriteRequest)

	if err := c.Bind(payload); err != nil {
		return helpers.ResponseInvalidPayload(c)
	}

	if err := c.Validate(payload); err != nil {
		errMessage := helpers.GenerateValidationErrorMessage(err)
		return helpers.Response(c, http.StatusBadRequest, errMessage)
	}

	payload.AuthorId = c.Get("claims").(jwt.MapClaims)["Id"].(string)
	payload.PostID = c.Param("postId")

	return cm.commentService.Create(c, payload)
}

func (cm *commentHandler) GetAllByPost(c echo.Context) error {
	cursor, err := helpers.GenerateCursorPaginationByEcho(c)
	if err != nil {
		if strings.Contains(err.Error(), "unavailable") {
			return helpers.Response(c, http.StatusOK, "success get all comments", []models.Post{}, helpers.CursorPagination{})
		}
		return helpers.Response(c, http.StatusBadRequest, err.Error())
	}

	return cm.commentService.GetAllByPost(c, c.Param("postId"), c.QueryParam("ParentId"), cursor)
}
