package service

import (
	"context"
	r "reservation-service/genproto/reservation"
	st "reservation-service/storage/postgres"
)

type PostService struct {
	storage st.Storage
	r.UnimplementedPostServiceServer
}

func NewPostService(storage *st.Storage) *PostService {
	return &PostService{
		storage: *storage,
	}
}

func (s *PostService) Create(ctx context.Context, post *r.PostReq) (*r.Post, error) {
	resp, err := s.storage.PostS.Create(post)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) Get(ctx context.Context, idReq *r.GetByIdReq) (*r.PostRes, error) {
	resp, err := s.storage.PostS.Get(idReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) GetAll(ctx context.Context, allPosts *r.GetAllPostReq) (*r.GetAllPostRes, error) {
	items, err := s.storage.PostS.GetAll(allPosts)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *PostService) Update(ctx context.Context, post *r.PostUpdate) (*r.Post, error) {
	resp, err := s.storage.PostS.Update(post)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) Delete(ctx context.Context, idReq *r.GetByIdReq) (*r.Void, error) {
	_, err := s.storage.PostS.Delete(idReq)

	return nil, err
}
