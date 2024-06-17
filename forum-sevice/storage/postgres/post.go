package managers

import (
	"database/sql"
	"fmt"
	pb "forum-service/forum-protos/genprotos"
)

type PostManager struct {
	Conn *sql.DB
}

func NewPostManager(conn *sql.DB) *PostManager {
	return &PostManager{Conn: conn}
}

func (m *PostManager) Create(post *pb.PostCReqOrCResOrGResOrUResp) (*pb.PostCReqOrCResOrGResOrUResp, error) {
	query := "INSERT INTO posts (post_id, user_id, title, body, category_id, tags) VALUES ($1, $2, $3, $4, $5, $6) RETURNING post_id, user_id, title, body, category_id, tags"
	p := &pb.PostCReqOrCResOrGResOrUResp{}
	err := m.Conn.QueryRow(query, post.PostId, post.UserId, post.Title, post.Body, post.CategoryId, post.Tags).Scan(&p.PostId, &p.UserId, &p.Title, &p.Body, &p.CategoryId, &p.Tags)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *PostManager) Update(post *pb.PostUReq) (*pb.PostCReqOrCResOrGResOrUResp, error) {
	query := "UPDATE posts SET title = $1, body = $2, category_id = $3, tags = $4, updated_at = NOW() WHERE post_id = $5 RETURNING post_id, user_id, title, body, category_id, tags"
	p := &pb.PostCReqOrCResOrGResOrUResp{}
	err := m.Conn.QueryRow(query, post.Title, post.Body, post.CategoryId, post.Tags, post.PostId).Scan(&p.PostId, &p.UserId, &p.Title, &p.Body, &p.CategoryId, &p.Tags)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *PostManager) GetByID(req *pb.PostGReqOrDReq) (*pb.PostCReqOrCResOrGResOrUResp, error) {
	query := "SELECT post_id, user_id, title, body, category_id, tags FROM posts WHERE post_id = $1"
	p := &pb.PostCReqOrCResOrGResOrUResp{}
	err := m.Conn.QueryRow(query, req.PostId).Scan(&p.PostId, &p.UserId, &p.Title, &p.Body, &p.CategoryId, &p.Tags)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post not found")
		}
		return nil, err
	}
	return p, nil
}

func (m *PostManager) Delete(req *pb.PostGReqOrDReq) (*pb.Void, error) {
	query := "UPDATE posts SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE post_id = $1"
	_, err := m.Conn.Exec(query, req.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *PostManager) GetAll(req *pb.PostGAReq) (*pb.PostGARes, error) {
	query := "SELECT post_id, user_id, title, body, category_id, tags FROM posts WHERE deleted_at = 0"
	rows, err := m.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := &pb.PostGARes{}
	for rows.Next() {
		p := &pb.PostCReqOrCResOrGResOrUResp{}
		if err := rows.Scan(&p.PostId, &p.UserId, &p.Title, &p.Body, &p.CategoryId, &p.Tags); err != nil {
			return nil, err
		}
		posts.Posts = append(posts.Posts, p)
		posts.Count++
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
