package repositories

import (
	"context"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/oscarvo29/real-chat-backend/models"
)

func SaveUser(user *models.User) error {
	query := `INSERT INTO users (name, password) VALUES ($1, $2)`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Password)
	if err != nil {
		return err
	}

	user.Id = 3
	return nil
}

func GetAllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var users []*models.User
	query := `SELECT * FROM users`
	rows, err := DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Name, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func ValidateLogin(name, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM users WHERE name = $1 AND password = $2`
	var user models.User
	row := DB.QueryRowContext(ctx, query, name, password)

	err := row.Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
