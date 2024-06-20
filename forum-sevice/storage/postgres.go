package storage

import (
	"database/sql"
	"fmt"
	"log"

	"forum-service/config"
	managers "forum-service/storage/postgres"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db        *sql.DB
	PostS     PostI
	CategoryS CategoryI
	TagS      TagI
	CommentS  CommentI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
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

	c_repo := managers.NewCategoryManager(db)
	t_repo := managers.NewTagManager(db)
	cm_repo := managers.NewCommentManager(db)
	p_repo := managers.NewPostManager(db, t_repo, cm_repo)

	log.Println("Successfully connected to the database")
	return &Storage{
		Db:        db,
		PostS:     p_repo,
		CategoryS: c_repo,
		TagS:      t_repo,
		CommentS:  cm_repo,
	}, nil
}
