package controller

import (
	"context"
	"errors"
	"fmt"
)

type db interface {
	CreateUser(ctx context.Context, u User) error
	ReadUser(ctx context.Context, idUser uint) (User, error)
	UpdateUser(ctx context.Context, u User) error
	CreateOrder(ctx context.Context, o Order) error
}

type Controller struct {
	db db
}

func New(
	db db,
) *Controller {
	return &Controller{
		db: db,
	}
}

func (s *Controller) GetBalance(ctx context.Context, req RequestBalance) (u User, err error) {
	if req.UserID <= 0 {
		err = fmt.Errorf("id must be more than 0")
		return
	}

	u, err = s.db.ReadUser(ctx, req.UserID)
	return
}

func (s *Controller) AddBalance(ctx context.Context, req RequestBalance) (u User, err error) {
	if req.UserID <= 0 {
		err = fmt.Errorf("id must be more than 0")
		return
	}

	u, err = s.db.ReadUser(ctx, req.UserID)
	if err != nil {
		u = User{
			ID:      req.UserID,
			Balance: req.AddSum,
		}
		err = s.CreateUser(ctx, u)
		return
	}

	u.Balance += req.AddSum
	err = s.db.UpdateUser(ctx, u)
	return
}

func (s *Controller) CreateUser(ctx context.Context, u User) error {
	return s.db.CreateUser(ctx, u)
}

func (s *Controller) CreateOrder(ctx context.Context, req Order) error {
	if req.UserID <= 0 {
		return fmt.Errorf("id must be more than 0")
	}

	u, err := s.db.ReadUser(ctx, req.UserID)
	if err != nil {
		return err
	}
	if u.Balance < req.Amount {
		return errors.New("out of money")
	}

	u.Balance -= req.Amount
	u.Reserve += req.Amount

	err = s.db.UpdateUser(ctx, u)
	if err != nil {
		return err
	}

	return s.db.CreateOrder(ctx, req)
}

func (s *Controller) FinishOrder(ctx context.Context, req Order) error {
	if req.UserID <= 0 {
		return fmt.Errorf("id must be more than 0")
	}

	u, err := s.db.ReadUser(ctx, req.UserID)
	if err != nil {
		return err
	}

	u.Reserve -= req.Amount

	return s.db.UpdateUser(ctx, u)
}
