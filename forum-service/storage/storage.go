package storage

import (
	"database/sql"
	pb "forum-service/forum-protos/genprotos"
)

type StorageI interface {
	Post() PostI
	Comment() CommentI
	Category() CategoryI
	Tag() TagI
}

type PostI interface {
	Create(*pb.PostCReqOrCResOrGResOrUResp, []string) (*pb.PostCReqOrCResOrGResOrUResp, error)
	GetByID(*pb.PostGReqOrDReq) (*pb.PostCReqOrCResOrGResOrUResp, error)
	GetAll(*pb.PostGAReq) (*pb.PostGARes, error)
	Update(*pb.PostUReq, []string) (*pb.PostCReqOrCResOrGResOrUResp, error)
	Delete(*pb.PostGReqOrDReq) (*pb.Void, error)
}

type CommentI interface {
	Create(*pb.CommentCReqOrCResOrGResOrURes) (*pb.CommentCReqOrCResOrGResOrURes, error)
	GetByID(*pb.CommentGReqOrDReq) (*pb.CommentCReqOrCResOrGResOrURes, error)
	GetAll(*pb.CommentGAReq) (*pb.CommentGARes, error)
	Update(*pb.CommentUReq) (*pb.CommentCReqOrCResOrGResOrURes, error)
	Delete(*sql.Tx, *pb.CommentGReqOrDReq) (*pb.Void, error)
	DeleteByPostID(*sql.Tx, *pb.CommentGReqOrDReqByPostID) (*pb.Void, error)
}

type CategoryI interface {
	Create(*pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error)
	GetByID(*pb.CategoryGReqOrDReq) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error)
	GetAll(*pb.CategoryGAReq) (*pb.CategoryGARes, error)
	Update(*pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error)
	Delete(*pb.CategoryGReqOrDReq) (*pb.Void, error)
}

type TagI interface {
	Create(*sql.Tx, *pb.TagCReqOrCRes) (*pb.TagCReqOrCRes, error)
	Delete(*sql.Tx, *pb.TagGReqOrDReq) (*pb.Void, error)
	GetPopular(*pb.Pagination) (*pb.TagPopularRes, error)
}
