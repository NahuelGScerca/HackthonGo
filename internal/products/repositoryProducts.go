package products

import (
	"context"
	"database/sql"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Get(ctx context.Context, id int) (models.Products, error)
	ImportData(ctx context.Context, prodToAdd []models.Products) ([]models.Products, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (models.Products, error) {
	query := "SELECT * FROM products WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := models.Products{}
	err := row.Scan(&b.ID, &b.Description, &b.Price)
	if err != nil {
		return models.Products{}, err
	}

	return b, nil
}

func (r *repository) ImportData(ctx context.Context, prodToAdd []models.Products) ([]models.Products, error) {
	// query := "INSERT INTO products(description,price) VALUES (?,?,?)"
	// stmt, err := r.db.Prepare(query)
	// if err != nil {
	// 	return 0, err
	// }

	// res, err := stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName)
	// if err != nil {
	// 	return 0, err
	// }

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }

	return nil, nil
}
