package managers

import (
	"database/sql"
	"fmt"
	pb "forum-service/forum-protos/genprotos"
)

type CategoryManager struct {
	Conn *sql.DB
}

func NewCategoryManager(conn *sql.DB) *CategoryManager {
	return &CategoryManager{Conn: conn}
}

func (m *CategoryManager) Create(category *pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error) {
	query := "INSERT INTO categories (category_id, name) VALUES ($1, $2) RETURNING category_id, name"
	cat := &pb.CategoryCReqOrCResOrGResOrUReqOrURes{}
	err := m.Conn.QueryRow(query, category.CategoryId, category.Name).Scan(&cat.CategoryId, &cat.Name)
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (m *CategoryManager) Update(category *pb.CategoryCReqOrCResOrGResOrUReqOrURes) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error) {
	query := "UPDATE categories SET name = $1, updated_at = NOW() WHERE category_id = $2 RETURNING category_id, name"
	cat := &pb.CategoryCReqOrCResOrGResOrUReqOrURes{}
	err := m.Conn.QueryRow(query, category.Name, category.CategoryId).Scan(&cat.CategoryId, &cat.Name)
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (m *CategoryManager) GetByID(req *pb.CategoryGReqOrDReq) (*pb.CategoryCReqOrCResOrGResOrUReqOrURes, error) {
	query := "SELECT category_id, name FROM categories WHERE category_id = $1 AND deleted_at = 0"
	cat := &pb.CategoryCReqOrCResOrGResOrUReqOrURes{}
	err := m.Conn.QueryRow(query, req.CategoryId).Scan(&cat.CategoryId, &cat.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category not found")
		}
		return nil, err
	}
	return cat, nil
}

func (m *CategoryManager) Delete(req *pb.CategoryGReqOrDReq) (*pb.Void, error) {
	query := "UPDATE categories SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE category_id = $1"
	_, err := m.Conn.Exec(query, req.CategoryId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *CategoryManager) GetAll(req *pb.CategoryGAReq) (*pb.CategoryGARes, error) {
	query := "SELECT category_id, name FROM categories WHERE deleted_at = 0"
	var args []interface{}
	var paramInex = 1
	if req.Filter.CategoryId != "" {
		query += fmt.Sprintf(" AND category_id = $%d", paramInex)
		args = append(args, req.Filter.CategoryId)
		paramInex++

	}
	if req.Pagination.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", paramInex)
		args = append(args, req.Pagination.Limit)
		paramInex++
	}
	if req.Pagination.Offset != 0 {
		query += fmt.Sprintf(" OFFSET $%d", paramInex)
		args = append(args, req.Pagination.Offset)
		paramInex++
	}
	rows, err := m.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := &pb.CategoryGARes{}
	for rows.Next() {
		cat := &pb.CategoryCReqOrCResOrGResOrUReqOrURes{}
		if err := rows.Scan(&cat.CategoryId, &cat.Name); err != nil {
			return nil, err
		}
		categories.Categories = append(categories.Categories, cat)
		categories.Count++
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
