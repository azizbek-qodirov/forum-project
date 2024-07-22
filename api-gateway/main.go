package main

import (
	"api-gateway/api"
	cf "api-gateway/config"
	"api-gateway/config/logger"
	"fmt"
	"path/filepath"

	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	cll := fmt.Sprintf("forum-service%s", ":50051")
	//fmt.Println(cll, "2343244444")
	ForumConn, err := grpc.NewClient(cll, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println(ForumConn, 1111111111111)
	em.CheckErr(err, 27)
	defer ForumConn.Close()

	r := api.NewRouter(ForumConn, *logger)

	fmt.Printf("Server started on port %s\n", config.HTTPPort)
	logger.INFO.Println("Server started on port: " + config.HTTPPort)
	if r.Run(config.HTTPPort); err != nil {
		logger.ERROR.Panicln("Handling stopped due to error " + err.Error())
	}
}
