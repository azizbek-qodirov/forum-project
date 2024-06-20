package managers_test

import (
	"fmt"
	pb "forum-service/forum-protos/genprotos"
	managers "forum-service/storage/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	fmt.Println("Testing create post...")
	postId = uuid.NewString()
	categoryId = uuid.NewString()

	createCategory(t, categoryId)

	post := &pb.PostCReqOrCResOrGResOrUResp{
		PostId:     postId,
		UserId:     "123e4567-e89b-12d3-a456-426614174000", // Replace with an actual user ID
		Title:      "Test Post",
		Body:       "Test Post Body",
		CategoryId: categoryId,
		Tags:       "tag1, tag2",
	}

	tags := []string{"tag1", "tag2"}
	p, err := postManager.Create(post, tags)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, post.PostId, p.PostId)
	assert.Equal(t, post.UserId, p.UserId)
	assert.Equal(t, post.Title, p.Title)
	assert.Equal(t, post.Body, p.Body)
	assert.Equal(t, post.CategoryId, p.CategoryId)
	assert.Equal(t, post.Tags, p.Tags)
	fmt.Println("OK. Post created succesfully.")
}

func TestUpdatePost(t *testing.T) {
	fmt.Println("Testing update post...")
	post := &pb.PostUReq{
		PostId:     postId,
		Title:      "Updated Post",
		Body:       "Updated Post Body",
		CategoryId: categoryId,
		Tags:       "tag1, tag3",
	}

	tags := []string{"tag1", "tag3"}
	p, err := postManager.Update(post, tags)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, post.PostId, p.PostId)
	assert.Equal(t, post.Title, p.Title)
	assert.Equal(t, post.Body, p.Body)
	assert.Equal(t, post.CategoryId, p.CategoryId)
	assert.Equal(t, post.Tags, p.Tags)
	fmt.Println("OK. Post updated succesfully.")
}

func TestGetPostByID(t *testing.T) {
	fmt.Println("Testing get post by ID...")
	req := &pb.PostGReqOrDReq{
		PostId: postId,
	}

	p, err := postManager.GetByID(req)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, req.PostId, p.PostId)
	fmt.Println("OK. Post retrieved succesfully.")
}

func TestDeletePost(t *testing.T) {
	fmt.Println("Testing delete post...")
	req := &pb.PostGReqOrDReq{
		PostId: postId,
	}

	_, err := postManager.Delete(req)
	assert.NoError(t, err)
	fmt.Println("OK. Post deleted succesfully.")
}

func TestGetAllPosts(t *testing.T) {
	fmt.Println("Testing get all posts...")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	postManager := managers.NewPostManager(db, nil, nil)

	rows := sqlmock.NewRows([]string{"post_id", "user_id", "title", "body", "category_id", "tags"}).
		AddRow("1", "user1", "Title 1", "Body 1", "cat1", "tag1").
		AddRow("2", "user2", "Title 2", "Body 2", "cat2", "tag2")

	mock.ExpectQuery("SELECT post_id, user_id, title, body, category_id, tags FROM posts WHERE deleted_at = 0").
		WillReturnRows(rows)

	req := &pb.PostGAReq{
		Filter: &pb.PostFilter{},
		Pagination: &pb.Pagination{
			Limit:  10,
			Offset: 0,
		},
	}

	posts, err := postManager.GetAll(req)

	assert.NoError(t, err)
	assert.NotNil(t, posts)
	assert.True(t, posts.Count > 0)
	assert.Equal(t, 2, int(posts.Count))
	assert.Equal(t, "1", posts.Posts[0].PostId)
	assert.Equal(t, "user1", posts.Posts[0].UserId)
	assert.Equal(t, "Title 1", posts.Posts[0].Title)
	assert.Equal(t, "Body 1", posts.Posts[0].Body)
	assert.Equal(t, "cat1", posts.Posts[0].CategoryId)
	assert.Equal(t, "tag1", posts.Posts[0].Tags)
	fmt.Println("OK. All posts retrieved succesfully.")
}

func createCategory(t *testing.T, categoryId string) {
	fmt.Println("Creating category...")
	query := "INSERT INTO categories (category_id, name) VALUES ($1, $2)"
	_, err := db.Exec(query, categoryId, "Test Category")
	if err != nil {
		t.Fatalf("could not create category: %v", err)
	}
	fmt.Println("OK. Category created succesfully.")
}
