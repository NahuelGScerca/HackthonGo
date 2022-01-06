package invoices

import (
	"context"
	"database/sql"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Get(ctx context.Context, id int) (models.Invoices, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (models.Invoices, error) {
	query := "SELECT * FROM invoices WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := models.Invoices{}
	err := row.Scan(&b.ID, &b.Datetime, &b.IdCustomer, &b.Total)
	if err != nil {
		return models.Invoices{}, err
	}

	return b, nil
}
