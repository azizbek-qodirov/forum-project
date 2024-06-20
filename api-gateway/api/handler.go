package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/docs"
	"api-gateway/api/handlers"
	"api-gateway/api/middleware"
	"api-gateway/config/logger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connF *grpc.ClientConn, logger logger.Logger) *gin.Engine {
	h := handlers.NewHandler(connF, logger)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/", middleware.JWTMiddleware())

	// Category routes
	category := protected.Group("/category")
	category.POST("/", h.CategoryCreate)
	category.GET("/:id", h.CategoryGet)
	category.PUT("/:id", h.CategoryUpdate)
	category.DELETE("/:id", h.CategoryDelete)
	protected.GET("/categories", h.CategoryGetAll)

	// Post routes
	post := protected.Group("/post")
	post.POST("/", h.PostCreate)
	post.GET("/:id", h.PostGet)
	post.PUT("/:id", h.PostUpdate)
	post.DELETE("/:id", h.PostDelete)
	protected.GET("/posts", h.PostGetAll)

	// Comment routes
	comment := protected.Group("/comment")
	comment.POST("/", h.CommentCreate)
	comment.GET("/:id", h.CommentGet)
	comment.PUT("/:id", h.CommentUpdate)
	comment.DELETE("/:id", h.CommentDelete)
	protected.GET("/comments", h.CommentGetAll)

	// Tag routes
	protected.GET("/popular-tags", h.PopularTagsGet)

	return router
}
