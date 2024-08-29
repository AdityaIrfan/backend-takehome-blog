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

type postHandler struct {
	postService contract.IPostService
}

func NewPostHandler(postService contract.IPostService) contract.IPostHandler {
	return &postHandler{
		postService: postService,
	}
}

func (p *postHandler) Create(c echo.Context) error {
	payload := new(models.PostCreateRequest)

	if err := c.Bind(payload); err != nil {
		return helpers.ResponseInvalidPayload(c)
	}

	if err := c.Validate(payload); err != nil {
		errMessage := helpers.GenerateValidationErrorMessage(err)
		return helpers.Response(c, http.StatusBadRequest, errMessage)
	}

	return p.postService.Create(c, payload)
}

func (p *postHandler) Update(c echo.Context) error {
	payload := new(models.PostUpdateRequest)

	if err := c.Bind(payload); err != nil {
		return helpers.ResponseInvalidPayload(c)
	}

	if err := c.Validate(payload); err != nil {
		errMessage := helpers.GenerateValidationErrorMessage(err)
		return helpers.Response(c, http.StatusBadRequest, errMessage)
	}

	payload.ID = c.Param("id")
	payload.AuthorID = c.Get("claims").(jwt.MapClaims)["Id"].(string)

	return p.postService.Update(c, payload)
}

func (p *postHandler) GetAllMine(c echo.Context) error {
	cursor, err := helpers.GenerateCursorPaginationByEcho(c)
	if err != nil {
		if strings.Contains(err.Error(), "unavailable") {
			return helpers.Response(c, http.StatusOK, "success get all posts", []models.Post{}, helpers.CursorPagination{})
		}
		return helpers.Response(c, http.StatusBadRequest, err.Error())
	}

	return p.postService.GetAllMine(c, c.Get("claims").(jwt.MapClaims)["Id"].(string), cursor)
}

func (p *postHandler) GetAll(c echo.Context) error {
	cursor, err := helpers.GenerateCursorPaginationByEcho(c)
	if err != nil {
		if strings.Contains(err.Error(), "unavailable") {
			return helpers.Response(c, http.StatusOK, "success get all posts", []models.Post{}, helpers.CursorPagination{})
		}
		return helpers.Response(c, http.StatusBadRequest, err.Error())
	}

	return p.postService.GetAll(c, cursor)
}

func (p *postHandler) GetDetail(c echo.Context) error {
	return p.postService.GetDetail(c, c.Param("id"))
}

func (p *postHandler) Delete(c echo.Context) error {
	return p.postService.Delete(c, c.Param("id"))
}
