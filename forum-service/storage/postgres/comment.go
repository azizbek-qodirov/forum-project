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
	err := m.Conn.QueryRow(query, req.CommentId).Scan(&com.CommentId, &com.UserId, &com.PostId, &com.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("comment not found")
		}
		return nil, err
	}
	return com, nil
}

func (m *CommentManager) DeleteByPostID(tx *sql.Tx, req *pb.CommentGReqOrDReqByPostID) (*pb.Void, error) {
	query := "UPDATE comments SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE post_id = $1"
	_, err := tx.Exec(query, req.PostId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *CommentManager) Delete(tx *sql.Tx, req *pb.CommentGReqOrDReq) (*pb.Void, error) {
	query := "UPDATE comments SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE comment_id = $1"
	_, err := tx.Exec(query, req.CommentId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *CommentManager) GetAll(req *pb.CommentGAReq) (*pb.CommentGARes, error) {
	query := "SELECT comment_id, user_id, post_id, body FROM comments WHERE deleted_at = 0"
	var args []interface{}
	paramIndex := 1
	if req.Filter.PostId != "" {
		query += fmt.Sprintf(" AND post_id = $%d", paramIndex)
		args = append(args, req.Filter.PostId)
		paramIndex++
	}
	if req.Filter.UserId != "" {
		query += fmt.Sprintf(" AND user_id = $%d", paramIndex)
		args = append(args, req.Filter.UserId)
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
