package products

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Get(ctx context.Context, id int) (models.Products, error)
	Store(ctx context.Context, entidad models.Products) error
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

func (r *repository) Store(ctx context.Context, entidad models.Products) error {
	// fmt.Println("ENTIDAD:  ", entidad)
	// return nil
	query := "INSERT INTO products(id,description,price) VALUES (?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&entidad.ID, &entidad.Description, &entidad.Price)
	if err != nil {
		return err
	}

	if num, err := res.RowsAffected(); num > 0 && err == nil {
		return nil
	}

	return fmt.Errorf("Error insert")
}
