package handlers

import (
	"api-gateway/config/logger"
	pb "api-gateway/forum-protos/genprotos"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Comment  pb.CommentServiceClient
	Post     pb.PostServiceClient
	Category pb.CategoryServiceClient
	Tag      pb.TagServiceClient
	Logger   logger.Logger
}

func NewHandler(connF *grpc.ClientConn, l logger.Logger) *HTTPHandler {
	return &HTTPHandler{
		Comment:  pb.NewCommentServiceClient(connF),
		Post:     pb.NewPostServiceClient(connF),
		Category: pb.NewCategoryServiceClient(connF),
		Tag:      pb.NewTagServiceClient(connF),
		Logger:   l,
	}
}
