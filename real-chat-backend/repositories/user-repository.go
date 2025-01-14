package repositories

import (
	"context"

	"github.com/google/uuid"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/oscarvo29/real-chat-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(user *models.User) error {
	uuidString := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (uuid, name, password) VALUES ($1, $2, $3)`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuidString, user.Name, hashedPassword)
	if err != nil {
		return err
	}

	user.Uuid = uuidString
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

func GetUserFromName(name string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM users WHERE name = $1`
	var user models.User
	row := DB.QueryRowContext(ctx, query, name)

	err := row.Scan(&user.Uuid, &user.Name, &user.Password)
	if err != nil {
		return &user, err
	}

	return &user, nil
}
