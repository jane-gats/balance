package db

import (
	"balance/internal/config"
	//"avito/internal/pkg/models"
	"database/sql"
	//"fmt"
	//"time"
)

type Pool struct {
	pool *sql.DB
}
func CreatePool(c *config.DBConfig) (*Pool, error) {
	dsn := "postgres://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(c.Connections)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Pool{pool: db}, nil
}

func (p *Pool) Close() error {
	return p.pool.Close()
}