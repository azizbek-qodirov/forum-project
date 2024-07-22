package managers_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	pb "forum-service/forum-protos/genprotos"
	managers "forum-service/storage/postgres"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var categoryManager *managers.CategoryManager
var commentManager *managers.CommentManager
var postManager *managers.PostManager
var tagManager *managers.TagManager

var categoryId string
var commentId string
var postId string

func TestMain(m *testing.M) {
	fmt.Println("Connecting to the database...")
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", "localhost", "mrbek", "forum_db", "QodirovCoder", 5432)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not connect to the DB: %v", err)
	}

	categoryManager = managers.NewCategoryManager(db)
	commentManager = managers.NewCommentManager(db)
	tagManager = managers.NewTagManager(db)
	postManager = managers.NewPostManager(db, tagManager, commentManager)

	fmt.Println("Database connected!!!")
	code := m.Run()

	fmt.Println("Database connection closed")
	db.Close()
	os.Exit(code)
}

func TestCreateCategory(t *testing.T) {
	fmt.Println("Testing create category...")
	categoryId = uuid.NewString()
	category := &pb.CategoryCReqOrCResOrGResOrUReqOrURes{
		CategoryId: categoryId,
		Name:       "Test Category",
	}

	cat, err := categoryManager.Create(category)
	assert.NoError(t, err)
	assert.NotNil(t, cat)
	fmt.Println("OK. Category created successfully")
}

func TestUpdateCategory(t *testing.T) {
	fmt.Println("Testing update category...")
	category := &pb.CategoryCReqOrCResOrGResOrUReqOrURes{
		CategoryId: categoryId,
		Name:       "Updated Category",
	}

	cat, err := categoryManager.Update(category)
	assert.NoError(t, err)
	assert.NotNil(t, cat)
	assert.Equal(t, category.CategoryId, cat.CategoryId)
	assert.Equal(t, category.Name, cat.Name)
	fmt.Println("OK. Category updated successfully")
}

func TestGetCategoryByID(t *testing.T) {
	fmt.Println("Testing get category by ID...")
	req := &pb.CategoryGReqOrDReq{
		CategoryId: categoryId,
	}

	cat, err := categoryManager.GetByID(req)
	assert.NoError(t, err)
	assert.NotNil(t, cat)
	assert.Equal(t, req.CategoryId, cat.CategoryId)
	fmt.Println("OK. Category retrieved successfully")
}

func TestDeleteCategory(t *testing.T) {
	fmt.Println("Testing delete category...")
	req := &pb.CategoryGReqOrDReq{
		CategoryId: categoryId,
	}

	_, err := categoryManager.Delete(req)
	assert.NoError(t, err)

	cat, err := categoryManager.GetByID(req)
	assert.Error(t, err)
	assert.Nil(t, cat)
	fmt.Println("OK. Category deleted successfully")
}

func TestGetAllCategories(t *testing.T) {
	fmt.Println("Testing get all categories...")
	req := &pb.CategoryGAReq{
		Filter: &pb.CategoryFilter{},
		Pagination: &pb.Pagination{
			Limit:  10,
			Offset: 0,
		},
	}

	cats, err := categoryManager.GetAll(req)
	assert.NoError(t, err)
	assert.NotNil(t, cats)
	assert.True(t, cats.Count > 0)
	fmt.Println("OK. Categories retrieved successfully")
}
