package invoices

import (
	"context"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Service interface {
	Get(ctx context.Context, id int) (models.Invoices, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (serv *service) Get(ctx context.Context, id int) (models.Invoices, error) {
	invoice, err := serv.repository.Get(ctx, id)
	if err != nil {
		return models.Invoices{}, err
	}
	return invoice, nil
}
