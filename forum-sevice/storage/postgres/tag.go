package managers

import (
	"database/sql"
	pb "forum-service/forum-protos/genprotos"
)

type TagManager struct {
	Conn *sql.DB
}

func NewTagManager(conn *sql.DB) *TagManager {
	return &TagManager{Conn: conn}
}

func (m *TagManager) Create(tag *pb.TagCReqOrCRes) (*pb.TagCReqOrCRes, error) {
	query := "INSERT INTO tags (tag, post_id) VALUES ($1, $2) RETURNING tag, post_id"
	t := &pb.TagCReqOrCRes{}
	err := m.Conn.QueryRow(query, tag.Tag, tag.PostId).Scan(&t.Tag, &t.PostId)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (m *TagManager) GetByID(req *pb.TagGReqOrDReq) (*pb.TagGAResOrPopularRes, error) {
	query := "SELECT tag, post_id FROM tags WHERE post_id = $1"
	rows, err := m.Conn.Query(query, req.PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := &pb.TagGAResOrPopularRes{}
	for rows.Next() {
		t := &pb.TagCReqOrCRes{}
		if err := rows.Scan(&t.Tag, &t.PostId); err != nil {
			return nil, err
		}
		tags.Tags = append(tags.Tags, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (m *TagManager) Delete(req *pb.TagGReqOrDReq) (*pb.Void, error) {
	query := "DELETE FROM tags WHERE post_id = $1"
	_, err := m.Conn.Exec(query, req.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (m *TagManager) GetAll(req *pb.TagGAReq) (*pb.TagGAResOrPopularRes, error) {
	query := "SELECT tag, post_id FROM tags"
	rows, err := m.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := &pb.TagGAResOrPopularRes{}
	for rows.Next() {
		t := &pb.TagCReqOrCRes{}
		if err := rows.Scan(&t.Tag, &t.PostId); err != nil {
			return nil, err
		}
		tags.Tags = append(tags.Tags, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (m *TagManager) Popular(req *pb.Pagination) (*pb.TagGAResOrPopularRes, error) {
	query := `
		SELECT tag, COUNT(*) as count
		FROM tags
		GROUP BY tag
		ORDER BY count DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := m.Conn.Query(query, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := &pb.TagGAResOrPopularRes{}
	for rows.Next() {
		t := &pb.TagCReqOrCRes{}
		if err := rows.Scan(&t.Tag, new(int)); err != nil {
			return nil, err
		}
		tags.Tags = append(tags.Tags, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
