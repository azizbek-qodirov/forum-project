package storage

import (
	"database/sql"
	"fmt"

	"forum-service/config"
	"forum-service/config/logger"
	managers "forum-service/storage/postgres"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db        *sql.DB
	Logger    *logger.Logger
	PostS     PostI
	CategoryS CategoryI
	TagS      TagI
	CommentS  CommentI
}

func NewPostgresStorage(config config.Config, logger *logger.Logger) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	p_repo := managers.NewPostManager(db)
	c_repo := managers.NewCategoryManager(db)
	t_repo := managers.NewTagManager(db)
	cm_repo := managers.NewCommentManager(db)

	return &Storage{
		Db:        db,
		PostS:     p_repo,
		CategoryS: c_repo,
		TagS:      t_repo,
		CommentS:  cm_repo,
		Logger:    logger,
	}, nil
}
