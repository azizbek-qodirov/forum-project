package service

import (
	"context"
	pb "forum-service/forum-protos/genprotos"
	st "forum-service/storage"
)

type TagService struct {
	storage st.Storage
	pb.UnimplementedTagServiceServer
}

func NewTagService(storage *st.Storage) *TagService {
	return &TagService{storage: *storage}
}

func (s *TagService) Create(ctx context.Context, tag *pb.TagCReqOrCRes) (*pb.TagCReqOrCRes, error) {
	resp, err := s.storage.TagS.Create(tag)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TagService) GetByID(ctx context.Context, idReq *pb.TagGReqOrDReq) (*pb.TagGAResOrPopularRes, error) {
	resp, err := s.storage.TagS.GetByID(idReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TagService) GetAll(ctx context.Context, allTags *pb.TagGAReq) (*pb.TagGAResOrPopularRes, error) {
	tags, err := s.storage.TagS.GetAll(allTags)

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *TagService) Delete(ctx context.Context, idReq *pb.TagGReqOrDReq) (*pb.Void, error) {
	_, err := s.storage.TagS.Delete(idReq)

	return nil, err
}
