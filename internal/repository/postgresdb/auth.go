package postgresdb

import (
	"Internship_backend_avito/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, password) values($1, $2) RETURNING id", "users")
	row := r.db.QueryRow(query, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", "users")
	err := r.db.Get(&user, query, username, password)

	return user, err
}
