package handlers

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authService contract.IAuthService
}

func NewAuthHandler(authService contract.IAuthService) contract.IAuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (a *authHandler) Register(c echo.Context) error {
	payload := new(models.Register)

	if err := c.Bind(payload); err != nil {
		return helpers.ResponseInvalidPayload(c)
	}

	if err := c.Validate(payload); err != nil {
		errMessage := helpers.GenerateValidationErrorMessage(err)
		return helpers.Response(c, http.StatusBadRequest, errMessage)
	}

	return a.authService.Register(c, payload)
}

func (a *authHandler) Login(c echo.Context) error {
	payload := new(models.Login)

	if err := c.Bind(payload); err != nil {
		return helpers.ResponseInvalidPayload(c)
	}

	if err := c.Validate(payload); err != nil {
		errMessage := helpers.GenerateValidationErrorMessage(err)
		return helpers.Response(c, http.StatusBadRequest, errMessage)
	}

	return a.authService.Login(c, payload)
}
