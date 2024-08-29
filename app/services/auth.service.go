package services

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/phuslu/log"
)

type authService struct {
	userRepo contract.IUserRepository
}

func NewAuthService(userRepo contract.IUserRepository) contract.IAuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Register(c echo.Context, in *models.Register) error {
	// check if email alreay in use
	userByEmail, err := a.userRepo.GetByCustomAndSelectedFields(map[string]interface{}{
		"email": in.Email,
	}, "id")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	} else if userByEmail != nil {
		return helpers.Response(c, http.StatusBadRequest, "email already in use")
	}

	salt, hash, err := helpers.GenerateHashAndSalt(in.Password)
	if err != nil {
		log.Error().Err(errors.New("FAILED TO GENERATE HASH AND SALT PASSWORD :  " + err.Error())).Msg("")
		return helpers.ResponseUnprocessableEntity(c)
	}

	user := &models.User{
		ID:           ulid.Make().String(),
		Name:         in.Name,
		Email:        in.Email,
		PasswordHash: hash,
		PasswordSalt: salt,
	}

	if err := a.userRepo.Create(user); err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	return helpers.Response(c, http.StatusCreated, "register success")
}

func (a *authService) Login(c echo.Context, in *models.Login) error {
	// get user by email
	user, err := a.userRepo.GetByCustomAndSelectedFields(map[string]interface{}{
		"email": in.Email,
	}, "id, password_hash, password_salt")
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	} else if user == nil {
		log.Error().Err(fmt.Errorf("USER WITH EMAIL %s DOES NOT EXIST", in.Email)).Msg("")
		return helpers.Response(c, http.StatusBadRequest, "incorrect credential")
	}

	if err := helpers.ComparePassword(user.PasswordHash, user.PasswordSalt, in.Password); err != nil {
		if !errors.Is(err, errors.New("incorrect password")) {
			return helpers.ResponseUnprocessableEntity(c)
		}
		return helpers.Response(c, http.StatusBadRequest, "incorrect credential")
	}

	tokenExpiration := time.Now().Add(helpers.LoginExpiration)
	token, err := helpers.GenerateToken(user.ID, tokenExpiration)
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c)
	}

	tokenResponse := &models.TokenResponse{
		Token: token,
	}

	return helpers.Response(c, http.StatusOK, "login success", tokenResponse)
}
