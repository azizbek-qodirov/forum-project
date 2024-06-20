package managers

import (
	"database/sql"
	"fmt"
	pb "forum-service/forum-protos/genprotos"
)

type TagManager struct {
	Conn *sql.DB
}

func NewTagManager(conn *sql.DB) *TagManager {
	return &TagManager{Conn: conn}
}

func (m *TagManager) Create(tx *sql.Tx, tag *pb.TagCReqOrCRes) (*pb.TagCReqOrCRes, error) {
	query := "INSERT INTO tags (tag, post_id) VALUES ($1, $2) RETURNING tag, post_id"
	t := &pb.TagCReqOrCRes{}
	err := tx.QueryRow(query, tag.Tag, tag.PostId).Scan(&t.Tag, &t.PostId)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (m *TagManager) Delete(tx *sql.Tx, req *pb.TagGReqOrDReq) (*pb.Void, error) {
	query := "DELETE FROM tags WHERE post_id = $1"
	_, err := tx.Exec(query, req.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *TagManager) GetPopular(req *pb.Pagination) (*pb.TagPopularRes, error) {
	query := `
		SELECT tag, COUNT(*) as count
		FROM tags
		GROUP BY tag
		ORDER BY count DESC
	`
	var args []interface{}
	paramIndex := 1
	if req.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", paramIndex)
		args = append(args, req.Limit)
		paramIndex++
	}
	if req.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", paramIndex)
		args = append(args, req.Offset)
		paramIndex++
	}
	rows, err := m.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := &pb.TagPopularRes{}
	for rows.Next() {
		t := &pb.TagPopular{}
		if err := rows.Scan(&t.Tag, &t.Count); err != nil {
			return nil, err
		}
		tags.Tags = append(tags.Tags, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
