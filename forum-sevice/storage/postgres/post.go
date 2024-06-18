package managers

import (
	"database/sql"
	"fmt"
	pb "forum-service/forum-protos/genprotos"
	"strings"
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
	var args []interface{}
	paramIndex := 1
	if req.Filter.UserId != "" {
		query += fmt.Sprintf(" AND user_id = $%d", paramIndex)
		args = append(args, req.Filter.UserId)
		paramIndex++
	}
	if req.Filter.CategoryId != "" {
		query += fmt.Sprintf(" AND category_id = $%d", paramIndex)
		args = append(args, req.Filter.CategoryId)
		paramIndex++
	}
	if req.Filter.Tags != "" {
		query += fmt.Sprintf(" AND tags ILIKE $%d", paramIndex)
		args = append(args, "%"+strings.ToLower(req.Filter.Tags)+"%")
		paramIndex++
	}
	if req.Filter.Body != "" {
		query += fmt.Sprintf(" AND body = $%d", paramIndex)
		args = append(args, req.Filter.Body)
		paramIndex++
	}
	if req.Filter.Title != "" {
		query += fmt.Sprintf(" AND title = $%d", paramIndex)
		args = append(args, req.Filter.Title)
		paramIndex++
	}
	if req.Pagination.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", paramIndex)
		args = append(args, req.Pagination.Limit)
		paramIndex++
	}
	if req.Pagination.Offset != 0 {
		query += fmt.Sprintf(" OFFSET $%d", paramIndex)
		args = append(args, req.Pagination.Offset)
		paramIndex++
	}
	rows, err := m.Conn.Query(query, args...)
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
