package service

import (
	"context"
	pb "forum-service/forum-protos/genprotos"
	st "forum-service/storage"
)

type CommentService struct {
	storage st.Storage
	pb.UnimplementedCommentServiceServer
}

func NewCommentService(storage *st.Storage) *CommentService {
	return &CommentService{storage: *storage}
}

func (s *CommentService) Create(ctx context.Context, comment *pb.CommentCReqOrCResOrGResOrURes) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	resp, err := s.storage.CommentS.Create(comment)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CommentService) GetByID(ctx context.Context, idReq *pb.CommentGReqOrDReq) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	resp, err := s.storage.CommentS.GetByID(idReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CommentService) GetAll(ctx context.Context, allComments *pb.CommentGAReq) (*pb.CommentGARes, error) {
	comments, err := s.storage.CommentS.GetAll(allComments)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentService) Update(ctx context.Context, reservation *pb.CommentUReq) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	resp, err := s.storage.CommentS.Update(reservation)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CommentService) Delete(ctx context.Context, idReq *pb.CommentGReqOrDReq) (*pb.Void, error) {
	_, err := s.storage.CommentS.Delete(idReq)

	return nil, err
}
