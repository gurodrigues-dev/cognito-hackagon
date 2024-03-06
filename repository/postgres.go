package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/config"
	"gin/types"
)

type Postgres struct {
	conn *sql.DB
}

func NewPostgres() (*Postgres, error) {

	conf := config.Get()

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	repo := &Postgres{
		conn: db,
	}

	return repo, nil
}

func (p *Postgres) SaveUser(ctx context.Context, user *types.User) error {

	sqlQuery := `INSERT INTO cognito (username, password) VALUES ($1, $2)`
	_, err := p.conn.Exec(sqlQuery, user.Username, user.Password)

	return err
}
