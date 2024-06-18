package service

import (
	"context"
	pb "forum-service/forum-protos/genprotos"
	st "forum-service/storage"

	"github.com/google/uuid"
)

type PostService struct {
	storage st.Storage
	pb.UnimplementedPostServiceServer
}

func NewPostService(storage *st.Storage) *PostService {
	return &PostService{storage: *storage}
}

func (s *PostService) Create(ctx context.Context, post *pb.PostCReqOrCResOrGResOrUResp) (*pb.PostCReqOrCResOrGResOrUResp, error) {
	post.PostId = uuid.NewString()
	resp, err := s.storage.PostS.Create(post)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) GetByID(ctx context.Context, idReq *pb.PostGReqOrDReq) (*pb.PostCReqOrCResOrGResOrUResp, error) {
	resp, err := s.storage.PostS.GetByID(idReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) GetAll(ctx context.Context, allPosts *pb.PostGAReq) (*pb.PostGARes, error) {
	posts, err := s.storage.PostS.GetAll(allPosts)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) Update(ctx context.Context, post *pb.PostUReq) (*pb.PostCReqOrCResOrGResOrUResp, error) {
	resp, err := s.storage.PostS.Update(post)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) Delete(ctx context.Context, idReq *pb.PostGReqOrDReq) (*pb.Void, error) {
	_, err := s.storage.PostS.Delete(idReq)

	return nil, err
}
