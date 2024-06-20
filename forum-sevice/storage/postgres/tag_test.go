package managers_test

import (
	"fmt"
	pb "forum-service/forum-protos/genprotos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTag(t *testing.T) {
	fmt.Println("Testing create tag...")
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	tag := &pb.TagCReqOrCRes{
		Tag:    "test-tag",
		PostId: postId,
	}

	createdTag, err := tagManager.Create(tx, tag)
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	assert.Equal(t, tag.Tag, createdTag.Tag)
	assert.Equal(t, tag.PostId, createdTag.PostId)
	fmt.Println("OK. Tag created successfully")
}

// func TestDeleteTag(t *testing.T) {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		t.Fatalf("Failed to begin transaction: %v", err)
// 	}
// 	defer tx.Rollback()

// 	_, err = tx.Exec(`
// 		INSERT INTO tags (tag, post_id) VALUES ($1, $2)`,
// 		"test-tag", postId)
// 	if err != nil {
// 		t.Fatalf("Failed to insert test tag: %v", err)
// 	}

// 	deleteReq := &pb.TagGReqOrDReq{
// 		PostId: postId,
// 	}

// 	_, err = tagManager.Delete(tx, deleteReq)
// 	if err != nil {
// 		t.Fatalf("Failed to delete tag: %v", err)
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		t.Fatalf("Failed to commit transaction: %v", err)
// 	}
// }
