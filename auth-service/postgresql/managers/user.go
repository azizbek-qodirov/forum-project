package managers

import (
	"auth-service/models"
	"database/sql"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{Conn: db}
}

func (m *UserManager) Register(req models.RegisterReq) error {
	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)"
	_, err := m.Conn.Exec(query, req.ID, req.Username, req.Email, req.Password)
	return err
}

func (m *UserManager) Profile(req models.GetProfileReq) (*models.GetProfileResp, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	row := m.Conn.QueryRow(query, req.Email)
	var user models.GetProfileResp
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserManager) EmailExists(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := m.Conn.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
