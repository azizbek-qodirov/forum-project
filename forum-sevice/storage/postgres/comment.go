package managers

import (
	"database/sql"
	"fmt"
	pb "forum-service/forum-protos/genprotos"
)

type CommentManager struct {
	Conn *sql.DB
}

func NewCommentManager(conn *sql.DB) *CommentManager {
	return &CommentManager{Conn: conn}
}

func (m *CommentManager) Create(comment *pb.CommentCReqOrCResOrGResOrURes) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	query := "INSERT INTO comments (comment_id, user_id, post_id, body) VALUES ($1, $2, $3, $4) RETURNING comment_id, user_id, post_id, body"
	com := &pb.CommentCReqOrCResOrGResOrURes{}
	err := m.Conn.QueryRow(query, comment.CommentId, comment.UserId, comment.PostId, comment.Body).Scan(&com.CommentId, &com.UserId, &com.PostId, &com.Body)
	if err != nil {
		return nil, err
	}
	return com, nil
}

func (m *CommentManager) Update(comment *pb.CommentUReq) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	query := "UPDATE comments SET body = $1, updated_at = NOW() WHERE comment_id = $2 RETURNING comment_id, user_id, post_id, body"
	com := &pb.CommentCReqOrCResOrGResOrURes{}
	err := m.Conn.QueryRow(query, comment.Body, comment.CommentId).Scan(&com.CommentId, &com.UserId, &com.PostId, &com.Body)
	if err != nil {
		return nil, err
	}
	return com, nil
}

func (m *CommentManager) GetByID(req *pb.CommentGReqOrDReq) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	query := "SELECT comment_id, user_id, post_id, body FROM comments WHERE comment_id = $1"
	com := &pb.CommentCReqOrCResOrGResOrURes{}
	err := m.Conn.QueryRow(query, req.PostId).Scan(&com.CommentId, &com.UserId, &com.PostId, &com.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("comment not found")
		}
		return nil, err
	}
	return com, nil
}

func (m *CommentManager) Delete(req *pb.CommentGReqOrDReq) (*pb.Void, error) {
	query := "UPDATE comments SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE comment_id = $1"
	_, err := m.Conn.Exec(query, req.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *CommentManager) GetAll(req *pb.CommentGAReq) (*pb.CommentGARes, error) {
	query := "SELECT comment_id, user_id, post_id, body FROM comments WHERE deleted_at = 0"
	rows, err := m.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := &pb.CommentGARes{}
	for rows.Next() {
		com := &pb.CommentCReqOrCResOrGResOrURes{}
		if err := rows.Scan(&com.CommentId, &com.UserId, &com.PostId, &com.Body); err != nil {
			return nil, err
		}
		comments.Comments = append(comments.Comments, com)
		comments.Count++
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
