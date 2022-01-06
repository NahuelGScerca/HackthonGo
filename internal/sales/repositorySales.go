package sales

import (
	"context"
	"database/sql"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Get(ctx context.Context, id int) (models.Sales, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (models.Sales, error) {
	query := "SELECT * FROM sales WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := models.Sales{}
	err := row.Scan(&b.ID, &b.LastName, &b.FirstName, &b.Condition)
	if err != nil {
		return models.Sales{}, err
	}

	return b, nil
}
