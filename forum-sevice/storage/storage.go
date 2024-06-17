package storage

import (
	pb "forum-service/forum-protos/genprotos"
)

type StorageI interface {
	Post() PostI
	Comment() CommentI
	Category() CategoryI
	Tag() TagI
}

type PostI interface {
	Create(*pb.PostCReqOrCResOrGResOrUResp) (*pb.PostCReqOrCResOrGResOrUResp, error)
	GetByID(*pb.PostGReqOrDReq) (*pb.PostCReqOrCResOrGResOrUResp, error)
	GetAll(*pb.PostGAReq) (*pb.PostGARes, error)
	Update(*pb.PostUReq) (*pb.PostCReqOrCResOrGResOrUResp, error)
	Delete(*pb.PostGReqOrDReq) (*pb.Void, error)
}

type CommentI interface {
	Create(*pb.CommentCReqOrCResOrGResOrURes) (*pb.CommentCReqOrCResOrGResOrURes, error)
	GetByID(*pb.CommentGReqOrDReq) (*pb.CommentCReqOrCResOrGResOrURes, error)
	GetAll(*pb.CommentGAReq) (*pb.CommentGARes, error)
	Update(*pb.CommentUReq) (*pb.CommentCReqOrCResOrGResOrURes, error)
	Delete(*pb.CommentGReqOrDReq) (*pb.Void, error)
}

type CategoryI interface {
	Create(*pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error)
	GetByID(*pb.CategoryGReqOrDReq) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error)
	GetAll(*pb.CategoryGAReq) (*pb.CategoryGARes, error)
	Update(*pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error)
	Delete(*pb.CategoryGReqOrDReq) (*pb.Void, error)
}

type TagI interface {
	Create(*pb.TagCReqOrCRes) (*pb.TagCReqOrCRes, error)
	GetByID(*pb.TagGReqOrDReq) ([]*pb.TagCReqOrCRes, error)
	GetAll(*pb.TagGAReq) (*pb.TagGAResOrPopularRes, error)
	Delete(*pb.TagGReqOrDReq) (*pb.Void, error)
	Popular(*pb.Pagination) (*pb.TagGAResOrPopularRes, error)
}
