package service

import (
	"context"
	pb "forum-service/forum-protos/genprotos"
	st "forum-service/storage"

	"github.com/google/uuid"
)

type CategoryService struct {
	storage st.Storage
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService(storage *st.Storage) *CategoryService {
	return &CategoryService{storage: *storage}
}

func (s *CategoryService) Create(ctx context.Context, category *pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error) {
	category.CategoryId = uuid.NewString()
	resp, err := s.storage.CategoryS.Create(category)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CategoryService) GetByID(ctx context.Context, idReq *pb.CategoryGReqOrDReq) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error) {
	resp, err := s.storage.CategoryS.GetByID(idReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CategoryService) GetAll(ctx context.Context, allCategories *pb.CategoryGAReq) (*pb.CategoryGARes, error) {
	orders, err := s.storage.CategoryS.GetAll(allCategories)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *CategoryService) Update(ctx context.Context, category *pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error) {
	resp, err := s.storage.CategoryS.Update(category)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CategoryService) Delete(ctx context.Context, idReq *pb.CategoryGReqOrDReq) (*pb.Void, error) {
	_, err := s.storage.CategoryS.Delete(idReq)

	return nil, err
}
