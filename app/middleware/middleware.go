package middleware

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
)

func NewMiddleware(userRepo contract.IUserRepository) *middleware {
	return &middleware{
		userRepo: userRepo,
	}
}

type middleware struct {
	userRepo contract.IUserRepository
}

func (m *middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := m.checkToken(c)
		if err != nil {
			return middlewareErrorResponse(c, err)
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("claims", claims)

		return next(c)
	}
}

func middlewareErrorResponse(c echo.Context, err error) error {
	switch err.Error() {
	case "missing authorization token":
		return helpers.Response(c, http.StatusBadRequest, "missing authorization header")
	case "invalid token":
		return helpers.Response(c, http.StatusUnauthorized, "invalid token")
	case "forbidden access":
		return helpers.Response(c, http.StatusUnauthorized, "forbidden access")
	case "expired token":
		return helpers.Response(c, http.StatusUnauthorized, "expired token")
	default:
		return helpers.Response(c, http.StatusUnauthorized, "forbidden access")
	}
}

func (m *middleware) checkToken(c echo.Context) (*jwt.Token, error) {
	if c.Request().Header.Get("Authorization") == "" {
		return nil, errors.New("missing authorization token")
	}

	tokens := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(tokens) != 2 {
		return nil, errors.New("invalid token")
	} else if tokens[0] != "Bearer" {
		return nil, errors.New("invalid token")
	}

	tokenString := tokens[1]

	token, err := helpers.VerifyToken(tokenString)
	if err != nil {
		log.Error().Err(errors.New("ERROR VERIFY TOKEN : " + err.Error())).Msg("")
		return nil, errors.New("invalid token")
	}

	userId, ok := token.Claims.(jwt.MapClaims)["Id"].(string)
	if !ok || userId == "" {
		log.Error().Err(errors.New("THIS TOKEN DOES NOT HAVE [Id] IN CLAIMS : ")).Msg("")
		return nil, errors.New("forbidden access")
	}

	expUnix, ok := token.Claims.(jwt.MapClaims)["Exp"].(float64)
	if !ok || userId == "" {
		log.Error().Err(errors.New("THIS TOKEN DOES NOT HAVE [Exp] IN CLAIMS : ")).Msg("")
		return nil, errors.New("forbidden access")
	}
	expirationTime := time.Unix(int64(expUnix), 0)
	if time.Now().After(expirationTime) {
		return nil, errors.New("expired token")
	}

	user, err := m.userRepo.GetByCustomAndSelectedFields(map[string]interface{}{
		"id": userId,
	}, "id")
	if err != nil {
		return nil, helpers.ResponseUnprocessableEntity(c)
	} else if user == nil {
		return nil, errors.New("forbidden access")
	}

	return token, nil
}
