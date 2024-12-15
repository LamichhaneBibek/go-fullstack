package sqlite

import (
	"database/sql"
	"time"

	"github.com/LamichhaneBibek/familytree/internal/models"
)

type UserModel struct{
	DB *sql.DB
}

func (m *UserModel) Insert(user *models.User) error {
	stmt := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, user.ID,user.Name, user.Email, user.Password, time.Now(), time.Now())
	return err
}

func (m *UserModel) All() ([]models.User, error) {
	stmt := `SELECT * FROM users`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created, &user.Updated)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}