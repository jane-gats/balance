package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"balance/internal/config"
	"balance/internal/controller"
)

type db struct {
	conn *sql.DB
}

func New(c *config.DBConfig) (*db, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &db{
		conn: conn,
	}, nil
}

func (s *db) Close() {
	s.conn.Close()
}

func (s *db) CreateUser(ctx context.Context, u controller.User) error {
	_, err := s.conn.Query("INSERT INTO public.user (id, balance, reserve) VALUES ($1, $2, $3)", u.ID, u.Balance, u.Reserve)
	return err
}

func (s *db) ReadUser(ctx context.Context, idUser uint) (controller.User, error) {
	rows, err := s.conn.Query("SELECT * FROM public.user WHERE id = $1", idUser)
	if err != nil {
		return controller.User{}, fmt.Errorf("select query: %w", err)
	}

	var (
		u controller.User
		i int
	)

	for rows.Next() {
		i++
		err := rows.Scan(&u.ID, &u.Balance, &u.Reserve)
		if err != nil {
			return controller.User{}, fmt.Errorf("scan result: %w", err)
		}
	}

	if i == 0 {
		return controller.User{}, fmt.Errorf("user not found")
	}

	return u, err
}

func (s *db) UpdateUser(ctx context.Context, u controller.User) error {
	_, err := s.conn.Query(
		`
		UPDATE public.user 
		SET 
			balance = $1,
			reserve = $2
		WHERE
			id = $3
	`,
		u.Balance,
		u.Reserve,
		u.ID,
	)
	return err
}
func (s *db) CreateOrder(ctx context.Context, o controller.Order) error {
	_, err := s.conn.Query(`
	INSERT INTO public.order (id, service_id, user_id, amount) 
	VALUES ($1, $2, $3, $4)`,
		o.ID,
		o.ServiceID,
		o.UserID,
		o.Amount,
	)
	return err
}
