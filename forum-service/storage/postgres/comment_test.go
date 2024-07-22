package managers_test

import (
	"fmt"
	"testing"

	pb "forum-service/forum-protos/genprotos"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	fmt.Println("Testing create comment...")
	commentId = uuid.NewString()
	postId = uuid.NewString()

	createPost(t, postId)

	comment := &pb.CommentCReqOrCResOrGResOrURes{
		CommentId: commentId,
		UserId:    "123e4567-e89b-12d3-a456-426614174000",
		PostId:    postId,
		Body:      "Test Comment",
	}

	com, err := commentManager.Create(comment)
	assert.NoError(t, err)
	assert.NotNil(t, com)
	assert.Equal(t, comment.CommentId, com.CommentId)
	assert.Equal(t, comment.UserId, com.UserId)
	assert.Equal(t, comment.PostId, com.PostId)
	assert.Equal(t, comment.Body, com.Body)
	fmt.Println("OK. Comment created successfully.")
}

func TestUpdateComment(t *testing.T) {
	fmt.Println("Testing update comment...")
	comment := &pb.CommentUReq{
		CommentId: commentId,
		Body:      "Updated Comment",
	}

	com, err := commentManager.Update(comment)
	assert.NoError(t, err)
	assert.NotNil(t, com)
	assert.Equal(t, comment.CommentId, com.CommentId)
	assert.Equal(t, comment.Body, com.Body)
	fmt.Println("OK. Comment updated successfully.")
}

func TestGetCommentByID(t *testing.T) {
	fmt.Println("Testing get comment by ID...")
	req := &pb.CommentGReqOrDReq{
		CommentId: commentId,
	}

	com, err := commentManager.GetByID(req)
	assert.NoError(t, err)
	assert.NotNil(t, com)
	assert.Equal(t, req.CommentId, com.CommentId)
	fmt.Println("OK. Comment retrieved successfully.")
}

func TestDeleteComment(t *testing.T) {
	fmt.Println("Testing delete comment...")
	req := &pb.CommentGReqOrDReq{
		CommentId: commentId,
	}
	tx, err := commentManager.Conn.Begin()
	if err != nil {
		t.Fatalf("could not begin transaction: %v", err)
	}

	_, err = commentManager.Delete(tx, req)
	if err != nil {
		tx.Rollback()
		t.Fatalf("could not delete comment: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatalf("could not commit transaction: %v", err)
	}
	assert.NoError(t, err)
	fmt.Println("OK. Comment deleted successfully.")
}

func TestGetAllComments(t *testing.T) {
	fmt.Println("Testing get all comments...")
	req := &pb.CommentGAReq{
		Filter: &pb.CommentFilter{},
		Pagination: &pb.Pagination{
			Limit:  10,
			Offset: 0,
		},
	}

	comments, err := commentManager.GetAll(req)
	assert.NoError(t, err)
	assert.NotNil(t, comments)
	fmt.Println("OK. Comments retrieved successfully.")
}

func createPost(t *testing.T, postId string) {
	fmt.Println("Creating post...")
	query := "INSERT INTO posts (post_id, user_id, title, body) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query, postId, "123e4567-e89b-12d3-a456-426614174000", "Test Post", "Test Post Body")
	if err != nil {
		t.Fatalf("could not create post: %v", err)
	}
	fmt.Println("OK. Post created successfully.")
}
