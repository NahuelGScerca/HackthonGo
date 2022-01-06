package customers

import (
	"context"
	"database/sql"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Get(ctx context.Context, id int) (models.Customers, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (models.Customers, error) {
	query := "SELECT * FROM products WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := models.Customers{}
	err := row.Scan(&b.ID, &b.LastName, &b.FirstName, &b.Condition)
	if err != nil {
		return models.Customers{}, err
	}

	return b, nil
}
