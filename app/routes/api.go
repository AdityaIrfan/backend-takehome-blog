package routes

import (
	"net/http"

	handler "backend-takehome-blog/handlers"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/middleware"
	repository "backend-takehome-blog/repositories"
	"backend-takehome-blog/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewApi() *api {
	return &api{}
}

type api struct{}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func (api) Init(db *gorm.DB) *echo.Echo {
	route := echo.New()
	route.Validator = &CustomValidator{Validator: validator.New(validator.WithRequiredStructEnabled())}
	// route.Use(middleware.RateLimiterMiddleware)

	// Repositories
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	commentRepository := repository.NewCommentRepository(db)

	// Services
	authService := services.NewAuthService(userRepository)
	postService := services.NewPostService(postRepository, commentRepository)
	commentService := services.NewCommentService(commentRepository, postRepository)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	postHandler := handler.NewPostHandler(postService)
	commentHandler := handler.NewCommentHandler(commentService)

	// Middleware
	m := middleware.NewMiddleware(userRepository)
	auth := m.Auth

	route.GET("/health", func(c echo.Context) error {
		return helpers.Response(c, http.StatusOK, "health")
	})

	route.POST("/register", authHandler.Register)
	route.POST("/login", authHandler.Login)

	route.POST("/posts", postHandler.Create, auth)
	route.GET("posts/:id", postHandler.GetDetail, auth)
	route.GET("/posts", postHandler.GetAll, auth)
	route.GET("my/posts", postHandler.GetAllMine, auth)
	route.PUT("/posts/:id", postHandler.Update, auth)
	route.DELETE("/posts/:id", postHandler.Delete, auth)

	route.POST("/posts/:postId/comments", commentHandler.Create, auth)
	route.GET("/posts/:postId/comments", commentHandler.GetAllByPost, auth)

	return route
}
