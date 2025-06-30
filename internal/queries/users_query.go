package queries

import (
	"github.com/jmoiron/sqlx"
	"github.com/physicist2018/lidar-web-app/internal/models"
)

type UsersQueries struct {
	*sqlx.DB
}

func (q *UsersQueries) GetUsers() ([]models.Users, error) {
	users := []models.Users{}
	query := `SELECT * FROM users`

	err := q.Get(&users, query)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (q *UsersQueries) GetUser(id int) (models.Users, error) {
	user := models.Users{}

	query := `SELECT * FROM users WHERE id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}
