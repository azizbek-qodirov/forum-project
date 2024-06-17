package main

import (
	"log"
	"net"
	"path/filepath"
	"runtime"

	cf "forum-service/config"
	"forum-service/config/logger"
	"forum-service/storage"

	pb "forum-service/forum-protos/genprotos"
	service "forum-service/service"

	"google.golang.org/grpc"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	db, err := storage.NewPostgresStorage(config, logger)
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
