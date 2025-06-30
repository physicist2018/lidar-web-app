package database

import "github.com/physicist2018/lidar-web-app/internal/queries"

type Queries struct {
	*queries.UsersQueries
}

func OpenDbConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		UsersQueries: &queries.UsersQueries{DB: db},
	}, nil
}
