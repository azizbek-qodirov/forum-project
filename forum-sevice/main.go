package main

import (
	"log"
	"net"

	cf "forum-service/config"
	"forum-service/storage"

	pb "forum-service/forum-protos/genprotos"
	service "forum-service/service"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()
	db, err := storage.NewPostgresStorage(config)
	em.CheckErr(err)
	defer db.Db.Close()

	listener, err := net.Listen("tcp", config.FORUM_SERVICE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, service.NewPostService(db))
	pb.RegisterCategoryServiceServer(s, service.NewCategoryService(db))
	pb.RegisterCommentServiceServer(s, service.NewCommentService(db))
	pb.RegisterTagServiceServer(s, service.NewTagService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
