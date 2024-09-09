package service

import (
	"user-service/internal/database"
	"user-service/internal/model"
)

func CreateUser(user *model.User) error {
	query := `INSERT INTO users (username, password, active) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, user.Username, user.Password, user.Active)
	return err
}

func GetUser(id int64) (*model.User, error) {
	query := `SELECT id, username, password, active FROM users WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Active)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *model.User) error {
	query := `UPDATE users SET username = ?, password = ?, active = ? WHERE id = ?`
	_, err := database.DB.Exec(query, user.Username, user.Password, user.Active, user.ID)
	return err
}

func DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func GetAllUsers() ([]*model.User, error) {
	query := `SELECT id, username, password, active FROM users`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Active)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}